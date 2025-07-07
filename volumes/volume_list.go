// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package volumes

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	kcclient "sdk.kraft.cloud/client"
)

// List implements VolumesService.
func (c *client) List(ctx context.Context, tags []string) (*kcclient.ServiceResponse[GetResponseItem], error) {
	resp := &kcclient.ServiceResponse[GetResponseItem]{}

	endpoint := Endpoint
	if len(tags) > 0 {
		endpoint = fmt.Sprintf("%s?tags=%s", Endpoint, strings.Join(tags, ","))
	}

	if err := c.request.DoRequest(ctx, http.MethodGet, endpoint, nil, resp); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	return resp, nil
}
