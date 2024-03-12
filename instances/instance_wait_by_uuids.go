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

// WaitByUUIDs implements InstancesService.
func (c *client) WaitByUUIDs(ctx context.Context, state State, timeoutMs int, uuids ...string) ([]WaitResponseItem, error) {
	if len(uuids) == 0 {
		return nil, errors.New("requires at least one name")
	}

	reqItems := make([]map[string]any, 0, len(uuids))
	for _, uuid := range uuids {
		reqItems = append(reqItems, map[string]any{
			"uuid":       uuid,
			"state":      state,
			"timeout_ms": timeoutMs,
		})
	}

	body, err := json.Marshal(reqItems)
	if err != nil {
		return nil, fmt.Errorf("encoding JSON object: %w", err)
	}

	var resp kcclient.ServiceResponse[WaitResponseItem]
	if err := c.request.DoRequest(ctx, http.MethodGet, Endpoint+"/wait", bytes.NewReader(body), &resp); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	items, err := resp.AllOrErr()
	if err != nil {
		errs := make([]error, 0, len(items)+1)
		errs = append(errs, err)
		for _, item := range items {
			if item.Error != nil {
				errs = append(errs, fmt.Errorf("%s (code=%d)", item.Message, *item.Error))
			}
		}
		return nil, errors.Join(errs...)
	}
	return items, nil
}
