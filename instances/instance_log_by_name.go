// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package instances

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	kcclient "sdk.kraft.cloud/client"
)

// LogByName implements InstancesService.
func (c *client) LogByName(ctx context.Context, name string, offset int, limit int) (*kcclient.ServiceResponse[LogResponseItem], error) {
	if len(name) == 0 {
		return nil, fmt.Errorf("name cannot be empty")
	}

	body, err := json.Marshal([]map[string]interface{}{{
		"name":   name,
		"offset": offset,
		"limit":  limit,
	}})
	if err != nil {
		return nil, fmt.Errorf("marshalling request body: %w", err)
	}

	resp := &kcclient.ServiceResponse[LogResponseItem]{}
	if err := c.request.DoRequest(ctx, http.MethodGet, Endpoint+"/log", bytes.NewReader(body), resp); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	return resp, nil
}
