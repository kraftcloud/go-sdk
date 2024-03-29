// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2024, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package autoscale

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	kcclient "sdk.kraft.cloud/client"
	"sdk.kraft.cloud/services"
)

// AddPolicy implements AutoscaleService.
func (c *client) AddPolicy(ctx context.Context, autoscaleUUID string, req Policy) (*kcclient.ServiceResponse[AddPolicyResponseItem], error) {
	if autoscaleUUID == "" {
		return nil, errors.New("autoscaleUUID cannot be empty")
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %w", err)
	}

	endpoint := services.Endpoint + "/" + autoscaleUUID + AutoscaleEndpoint + AutoscalePolicyEndpoint

	resp := &kcclient.ServiceResponse[AddPolicyResponseItem]{}
	if err := c.request.DoRequest(ctx, http.MethodPost, endpoint, bytes.NewReader(body), resp); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	return resp, nil
}
