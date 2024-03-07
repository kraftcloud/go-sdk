// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package users

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	kcclient "sdk.kraft.cloud/client"
)

// Quotas implements UsersService.
func (c *client) Quotas(ctx context.Context) (*QuotasResponseItem, error) {
	var resp kcclient.ServiceResponse[QuotasResponseItem]
	if err := c.request.DoRequest(ctx, http.MethodGet, Endpoint+"/quotas", nil, &resp); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	item, err := resp.FirstOrErr()
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("%s (code=%d)", item.Message, *item.Error))
	}
	return item, nil
}
