// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package instances

import (
	"context"
	"fmt"
	"net/http"

	"sdk.kraft.cloud/client"
)

// Lists all existing instances.
//
// See: https://docs.kraft.cloud/002-rest-api-v1-instances.html#list
func (c *instancesClient) List(ctx context.Context) ([]Instance, error) {
	endpoint := Endpoint + "/list"

	var response client.ServiceResponse[Instance]
	if err := c.request.DoRequest(ctx, http.MethodGet, endpoint, nil, &response); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	return response.AllOrErr()
}
