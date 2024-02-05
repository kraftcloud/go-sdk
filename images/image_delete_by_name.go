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
	names "github.com/google/go-containerregistry/pkg/name"
	v1 "github.com/google/go-containerregistry/pkg/v1"
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

// Delete an image by its name.
func (c *imagesClient) DeleteByName(ctx context.Context, name string) error {
	data, err := base64.StdEncoding.DecodeString(c.request.GetToken())
	if err != nil {
		return fmt.Errorf("could not decode token: %w", err)
	}

	split := strings.Split(string(data), ":")
	if len(split) != 2 {
		return fmt.Errorf("invalid token format")
	}

	ropts := []remote.Option{
		remote.WithPlatform(v1.Platform{
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

	if strings.HasPrefix(name, split[0]) {
		name = "index.unikraft.io/" + name
	}
	if strings.HasPrefix(name, "unikraft.io") {
		name = "index." + name
	}
	if !strings.HasPrefix(name, "index.unikraft.io") {
		name = "index.unikraft.io/official/" + name
	}

	ref, err := names.ParseReference(name,
		names.WithDefaultRegistry("index.unikraft.io"),
	)
	if err != nil {
		return fmt.Errorf("could not parse name: %w", err)
	}

	desc, err := remote.Get(ref, ropts...)
	if err != nil {
		return fmt.Errorf("could not get image: %w", err)
	}

	fullref, err := names.ParseReference(
		fmt.Sprintf("%s@%s", strings.SplitN(ref.Name(), ":", 2)[0], desc.Digest),
		names.WithDefaultRegistry("index.unikraft.io"),
	)
	if err != nil {
		return fmt.Errorf("could not parse full reference: %w", err)
	}

	if err := remote.Delete(fullref, ropts...); err != nil {
		return fmt.Errorf("could not delete image: %w", err)
	}

	return nil
}
