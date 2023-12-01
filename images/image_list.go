// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package images

import (
	"context"
	"fmt"
	"net/http"

	"sdk.kraft.cloud/client"
)

// List fetches all images from the Kraftcloud API.
//
// See: https://docs.kraft.cloud/004-rest-api-v1-images.html#list
func (c *imagesClient) List(ctx context.Context) ([]Image, error) {
	var response client.ServiceResponse[Image]
	if err := c.request.DoRequest(ctx, http.MethodGet, Endpoint+"/list", nil, &response); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	return response.AllOrErr()
}
