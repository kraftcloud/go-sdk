// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2024, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package certificates

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	ukcclient "sdk.kraft.cloud/client"
)

// Create implements InstancesService.
func (c *client) Create(ctx context.Context, req *CreateRequest) (*ukcclient.ServiceResponse[CreateResponseItem], error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %w", err)
	}

	resp := &ukcclient.ServiceResponse[CreateResponseItem]{}
	if err := c.request.DoRequest(ctx, http.MethodPost, Endpoint, bytes.NewReader(body), resp); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	return resp, nil
}
