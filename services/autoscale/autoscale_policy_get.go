// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2024, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package autoscale

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	ukcclient "sdk.kraft.cloud/client"
	"sdk.kraft.cloud/services"
)

// GetPolicy implements AutoscaleService.
func (c *client) GetPolicy(ctx context.Context, autoscaleUUID, name string) (*ukcclient.ServiceResponse[GetPolicyResponseItem], error) {
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}

	endpoint := services.Endpoint + "/" + autoscaleUUID + AutoscaleEndpoint + AutoscalePolicyEndpoint + "/" + name

	resp := &ukcclient.ServiceResponse[GetPolicyResponseItem]{}
	if err := c.request.DoRequest(ctx, http.MethodGet, endpoint, nil, resp); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	return resp, nil
}
