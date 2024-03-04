// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package images

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/google/go-containerregistry/pkg/authn"
	gcrname "github.com/google/go-containerregistry/pkg/name"
	gcrv1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/remote"
)

// simpleAuth is used to handle looking up the already populated
// user configuration that is used when speaking with the remote registry.
type simpleAuth struct {
	Auth *authn.AuthConfig
}

// Authorization implements authn.Authenticator.
func (auth *simpleAuth) Authorization() (*authn.AuthConfig, error) {
	return auth.Auth, nil
}

// Delete implements ImagesService.
func (c *client) DeleteByName(ctx context.Context, name string) error {
	data, err := base64.StdEncoding.DecodeString(c.request.GetToken())
	if err != nil {
		return fmt.Errorf("could not decode token: %w", err)
	}

	split := strings.Split(string(data), ":")
	if len(split) != 2 {
		return fmt.Errorf("invalid token format")
	}

	ropts := []remote.Option{
		remote.WithPlatform(gcrv1.Platform{
			OS:           "kraftcloud",
			Architecture: "x86_64",
		}),
		remote.WithAuth(&simpleAuth{
			Auth: &authn.AuthConfig{
				Username: split[0],
				Password: split[1],
			},
		}),
	}

	split[0] = strings.TrimPrefix(split[0], "robot$")
	split[0] = strings.TrimSuffix(split[0], ".users.kraftcloud")

	// If it has the user, add the domain
	if strings.HasPrefix(name, split[0]) {
		name = "index.unikraft.io/" + name
	}

	// If it has the old `unikraft.io` domain, add the `index.` prefix
	if strings.HasPrefix(name, "unikraft.io") {
		name = "index." + name
	}

	// If it has no domain and no user, add them both
	if !strings.Contains(name, "/") {
		name = "index.unikraft.io/" + split[0] + "/" + name
	}

	ref, err := gcrname.ParseReference(name,
		gcrname.WithDefaultRegistry("index.unikraft.io"),
	)
	if err != nil {
		return fmt.Errorf("could not parse name: %w", err)
	}

	desc, err := remote.Get(ref, ropts...)
	if err != nil {
		return fmt.Errorf("could not get image: %w", err)
	}

	name = strings.SplitN(ref.Name(), ":", 2)[0]
	name = strings.TrimSuffix(name, "@sha256")

	fullref, err := gcrname.ParseReference(
		fmt.Sprintf("%s@%s", name, desc.Digest),
		gcrname.WithDefaultRegistry("index.unikraft.io"),
	)
	if err != nil {
		return fmt.Errorf("could not parse full reference: %w", err)
	}

	if err := remote.Delete(fullref, ropts...); err != nil {
		return fmt.Errorf("could not delete image: %w", err)
	}

	return nil
}
