// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package services

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	kcclient "sdk.kraft.cloud/client"
)

// List implements ServicesService.
func (c *client) List(ctx context.Context) ([]ListResponseItem, error) {
	var resp kcclient.ServiceResponse[ListResponseItem]
	if err := c.request.DoRequest(ctx, http.MethodGet, Endpoint+"/list", nil, &resp); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	items, err := resp.AllOrErr()
	if err != nil {
		errs := make([]error, 0, len(items)+1)
		errs = append(errs, err)
		for _, item := range items {
			if item.Error != nil {
				errs = append(errs, fmt.Errorf("%s (code=%d)", item.Message, *item.Error))
			}
		}
		return nil, errors.Join(errs...)
	}
	return items, nil
}
