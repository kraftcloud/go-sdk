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

// DeleteByUUID implements ServicesService.
func (c *client) DeleteByUUID(ctx context.Context, uuid string) (*DeleteResponseItem, error) {
	if uuid == "" {
		return nil, errors.New("UUID cannot be empty")
	}

	var resp kcclient.ServiceResponse[DeleteResponseItem]
	if err := c.request.DoRequest(ctx, http.MethodDelete, Endpoint+"/"+uuid, nil, &resp); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	item, err := resp.FirstOrErr()
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("%s (code=%d)", item.Message, *item.Error))
	}
	return item, nil
}
