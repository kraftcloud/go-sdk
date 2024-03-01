// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package users

import (
	"context"
	"fmt"
	"net/http"

	"sdk.kraft.cloud/client"
)

// Lists quota usage and limits of your user account.  Limits are hard limits
// that cannot be exceeded.
//
// See: https://docs.kraft.cloud/api/v1/users/#list
func (c *usersClient) Quotas(ctx context.Context) (*Quotas, error) {
	var response client.ServiceResponse[Quotas]
	if err := c.request.DoRequest(ctx, http.MethodGet, Endpoint+"/quotas", nil, &response); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	user, err := response.FirstOrErr()
	if user != nil && user.Message != "" {
		err = fmt.Errorf("%w: %s", err, user.Message)
	}

	return user, err
}
