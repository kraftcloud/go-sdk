// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package users

import (
	"context"
	"fmt"
	"net/http"

	ukcclient "sdk.kraft.cloud/client"
)

// Quotas implements UsersService.
func (c *client) Quotas(ctx context.Context) (*ukcclient.ServiceResponse[QuotasResponseItem], error) {
	resp := &ukcclient.ServiceResponse[QuotasResponseItem]{}
	if err := c.request.DoRequest(ctx, http.MethodGet, Endpoint+"/quotas", nil, resp); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	return resp, nil
}
