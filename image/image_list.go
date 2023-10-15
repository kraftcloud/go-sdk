// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package image

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	kraftcloud "sdk.kraft.cloud"
)

// ImageListResponse holds the list of images description, as returned by the API.
type ImageListResponse struct {
	Status string `json:"status"`
	Data   struct {
		Images []Image `json:"images"`
	} `json:"data"`
}

// List fetches all images from the Kraftcloud API.
// see: https://docs.kraft.cloud/004-rest-api-v1-images.html#list
func (c *imagesClient) List(ctx context.Context, filter map[string]interface{}) ([]Image, error) {
	body, err := json.Marshal(filter)
	if err != nil {
		return nil, fmt.Errorf("marshalling request body: %w", err)
	}

	endpoint := fmt.Sprintf("%s/list", Endpoint)

	if c.request == nil {
		c.request = kraftcloud.NewServiceRequestFromDefaultOptions(c.defOpts)
	}

	defer func() { c.request = nil }()

	var response kraftcloud.ServiceResponse[Image]
	if err := c.request.DoRequest(ctx, http.MethodGet, endpoint, bytes.NewBuffer(body), &response); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	return response.AllOrErr()
}
