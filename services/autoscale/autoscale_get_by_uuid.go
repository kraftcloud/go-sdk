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

// GetConfigurationByUUID implements AutoscaleService.
func (c *client) GetConfigurationByUUID(ctx context.Context, uuid string) (*kcclient.ServiceResponse[GetResponseItem], error) {
	if uuid == "" {
		return nil, errors.New("UUID cannot be empty")
	}

	endpoint := services.Endpoint + "/" + uuid + AutoscaleEndpoint

	resp := &kcclient.ServiceResponse[GetResponseItem]{}
	if err := c.request.DoRequest(ctx, http.MethodGet, endpoint, nil, resp); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	return resp, nil
}
