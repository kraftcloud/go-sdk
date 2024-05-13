// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package services

import (
	"context"
	"fmt"
	"net/http"

	kcclient "sdk.kraft.cloud/client"
)

// List implements ServicesService.
func (c *client) List(ctx context.Context) (*kcclient.ServiceResponse[GetResponseItem], error) {
	resp := &kcclient.ServiceResponse[GetResponseItem]{}
	if err := c.request.DoRequest(ctx, http.MethodGet, Endpoint, nil, resp); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	return resp, nil
}
