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
	"sdk.kraft.cloud/uuid"
)

// Log implements InstancesService.
func (c *client) Log(ctx context.Context, id string, offset int, limit int) (*kcclient.ServiceResponse[LogResponseItem], error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("identifier cannot be empty")
	}

	reqItem := make(map[string]any, 3)
	if uuid.IsValid(id) {
		reqItem["uuid"] = id
	} else {
		reqItem["name"] = id
	}
	reqItem["offset"] = offset
	reqItem["limit"] = limit

	body, err := json.Marshal([]map[string]any{reqItem})
	if err != nil {
		return nil, fmt.Errorf("marshalling request body: %w", err)
	}

	resp := &kcclient.ServiceResponse[LogResponseItem]{}
	if err := c.request.DoRequest(ctx, http.MethodGet, Endpoint+"/log", bytes.NewReader(body), resp); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	return resp, nil
}
