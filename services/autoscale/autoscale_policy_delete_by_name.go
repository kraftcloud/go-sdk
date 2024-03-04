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

// DeletePolicyByName implements AutoscaleService.
func (c *client) DeletePolicyByName(ctx context.Context, autoscaleUUID, policyName string) (*DeletePolicyResponseItem, error) {
	if autoscaleUUID == "" || policyName == "" {
		return nil, errors.New("policyName and autoscaleUUID cannot be empty")
	}

	endpoint := services.Endpoint + "/" + autoscaleUUID + AutoscaleEndpoint + AutoscalePolicyEndpoint + "/" + policyName

	var resp kcclient.ServiceResponse[DeletePolicyResponseItem]
	if err := c.request.DoRequest(ctx, http.MethodDelete, endpoint, nil, &resp); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	item, err := resp.FirstOrErr()
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("%s (code=%d)", item.Message, *item.Error))
	}
	return item, nil
}
