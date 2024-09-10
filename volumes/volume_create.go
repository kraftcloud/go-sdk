// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package volumes

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	ukcclient "sdk.kraft.cloud/client"
)

// Create implements VolumesService.
func (c *client) Create(ctx context.Context, name string, sizeMB int) (*ukcclient.ServiceResponse[CreateResponseItem], error) {
	var err error
	var body []byte

	if sizeMB < 1 {
		return nil, fmt.Errorf("size_mb must be greater than 0")
	}

	if name == "" {
		body, err = json.Marshal(map[string]any{
			"size_mb": sizeMB,
		})
	} else {
		body, err = json.Marshal(map[string]any{
			"name":    name,
			"size_mb": sizeMB,
		})
	}
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %w", err)
	}

	resp := &ukcclient.ServiceResponse[CreateResponseItem]{}
	if err := c.request.DoRequest(ctx, http.MethodPost, Endpoint, bytes.NewReader(body), resp); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	return resp, nil
}
