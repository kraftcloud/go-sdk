// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package volumes

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	kcclient "sdk.kraft.cloud/client"
)

// Clone implements VolumeService.
func (c *client) Clone(ctx context.Context, source string, target string) (*kcclient.ServiceResponse[CloneResponseItem], error) {
	var err error
	var body []byte

	if source == "" {
		return nil, fmt.Errorf("source name must be provided")
	}

	if target == "" {
		return nil, fmt.Errorf("target volume name must be provided")
	}

	body, err = json.Marshal(map[string]any{
		"name":        source,
		"target_name": target,
	})
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %w", err)
	}

	// NOTE(craciunoiuc): Supports multiple pairs, but for now use only one
	// to avoid confusion.
	resp := &kcclient.ServiceResponse[CloneResponseItem]{}
	if err := c.request.DoRequest(ctx, http.MethodPut, Endpoint+"/clone", bytes.NewReader(body), resp); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	return resp, nil
}
