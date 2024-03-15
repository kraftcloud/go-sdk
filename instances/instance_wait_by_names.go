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

	kcclient "sdk.kraft.cloud/client"
)

// WaitByNames implements InstancesService.
func (c *client) WaitByNames(ctx context.Context, state State, timeoutMs int, names ...string) (*kcclient.ServiceResponse[WaitResponseItem], error) {
	if len(names) == 0 {
		return nil, errors.New("requires at least one name")
	}

	reqItems := make([]map[string]any, 0, len(names))
	for _, name := range names {
		reqItems = append(reqItems, map[string]any{
			"name":       name,
			"state":      state,
			"timeout_ms": timeoutMs,
		})
	}

	body, err := json.Marshal(reqItems)
	if err != nil {
		return nil, fmt.Errorf("encoding JSON object: %w", err)
	}

	resp := &kcclient.ServiceResponse[WaitResponseItem]{}
	if err := c.request.DoRequest(ctx, http.MethodGet, Endpoint+"/wait", bytes.NewReader(body), resp); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	return resp, nil
}
