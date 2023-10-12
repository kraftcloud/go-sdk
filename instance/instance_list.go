// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package instance

import (
	"context"
	"fmt"
	"net/http"

	kraftcloud "sdk.kraft.cloud/v0"
)

// Lists all existing instances.
//
// See: https://docs.kraft.cloud/002-rest-api-v1-instances.html#list
func (i *InstanceClient) List(ctx context.Context) ([]Instance, error) {
	base := i.BaseURL + Endpoint
	endpoint := fmt.Sprintf("%s/list", base)

	var response kraftcloud.ServiceResponse[Instance]
	if err := i.DoRequest(ctx, http.MethodGet, endpoint, nil, &response); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	return response.AllOrErr()
}
