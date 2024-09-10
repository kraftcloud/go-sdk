// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package instances

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	ukcclient "sdk.kraft.cloud/client"
	"sdk.kraft.cloud/uuid"
)

// Start implements InstancesService.
func (c *client) Start(ctx context.Context, waitTimeoutMs int, ids ...string) (*ukcclient.ServiceResponse[StartResponseItem], error) {
	if len(ids) == 0 {
		return nil, errors.New("requires at least one identifier")
	}

	reqItems := make([]map[string]any, 0, len(ids))
	for _, id := range ids {
		reqItem := make(map[string]any, 2)
		if uuid.IsValid(id) {
			reqItem["uuid"] = id
		} else {
			reqItem["name"] = id
		}
		if waitTimeoutMs > 0 {
			reqItem["wait_timeout_ms"] = waitTimeoutMs
		}
		reqItems = append(reqItems, reqItem)
	}

	body, err := json.Marshal(reqItems)
	if err != nil {
		return nil, fmt.Errorf("encoding JSON object: %w", err)
	}

	resp := &ukcclient.ServiceResponse[StartResponseItem]{}
	if err := c.request.DoRequest(ctx, http.MethodPut, Endpoint+"/start", bytes.NewReader(body), resp); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	return resp, nil
}
