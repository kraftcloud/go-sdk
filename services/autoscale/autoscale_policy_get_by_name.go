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

	kcclient "sdk.kraft.cloud/client"
	"sdk.kraft.cloud/services"
)

// GetPolicyByName implements AutoscaleService.
func (c *client) GetPolicyByName(ctx context.Context, uuid, name string) (*GetPolicyResponseItem, error) {
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}

	var resp kcclient.ServiceResponse[GetPolicyResponseItem]
	if err := c.request.DoRequest(ctx, http.MethodGet, services.Endpoint+"/"+uuid+"/autoscale/policies/"+name, nil, &resp); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	item, err := resp.FirstOrErr()
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("%s (code=%d)", item.Message, *item.Error))
	}
	return item, nil
}
