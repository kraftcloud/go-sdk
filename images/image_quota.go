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

	"github.com/goharbor/go-client/pkg/harbor"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/project"
)

// Delete implements ImagesService.
func (c *client) Quotas(ctx context.Context) (*QuotasResponseItem, error) {
	data, err := base64.StdEncoding.DecodeString(c.request.GetToken())
	if err != nil {
		return nil, fmt.Errorf("could not decode token: %w", err)
	}

	split := strings.Split(string(data), ":")
	if len(split) != 2 {
		return nil, fmt.Errorf("invalid token format")
	}

	user := split[0]
	pass := split[1]

	split[0] = strings.TrimPrefix(split[0], "robot$")
	split[0] = strings.TrimSuffix(split[0], ".users.kraftcloud")

	harborAPI, err := harbor.NewClientSet(&harbor.ClientSetConfig{
		URL:      "https://harbor.unikraft.io",
		Insecure: false,
		Username: user,
		Password: pass,
	})
	if err != nil {
		return nil, fmt.Errorf("could not create harbor client: %s", err)
	}

	params := project.NewGetProjectSummaryParams()
	params.SetProjectNameOrID(split[0])
	ok, err := harborAPI.V2().Project.GetProjectSummary(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("could not check quotas: %s", err)
	}
	quota := ok.GetPayload()

	if quota.Quota == nil {
		return nil, fmt.Errorf("could not check quotas: no quota found")
	}

	response := &QuotasResponseItem{
		Used: quota.Quota.Used["storage"],
		Hard: quota.Quota.Hard["storage"],
	}

	return response, nil
}
