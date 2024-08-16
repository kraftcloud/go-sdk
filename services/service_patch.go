// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2024, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	kcclient "sdk.kraft.cloud/client"
)

// Patch implements ServicesService.
func (c *client) Patch(ctx context.Context, req PatchRequest) (*kcclient.ServiceResponse[CreateResponseItem], error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %w", err)
	}

	resp := &kcclient.ServiceResponse[CreateResponseItem]{}
	if err := c.request.DoRequest(ctx, http.MethodPatch, Endpoint, bytes.NewReader(body), resp); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	return resp, nil
}
