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

// StartByUUIDs implements InstancesService.
func (c *client) StartByUUIDs(ctx context.Context, waitTimeoutMs int, uuids ...string) ([]StartResponseItem, error) {
	if len(uuids) == 0 {
		return nil, errors.New("requires at least one uuid")
	}

	reqItems := make([]map[string]any, 0, len(uuids))
	for _, uuid := range uuids {
		reqItem := map[string]any{
			"uuid": uuid,
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

	var resp kcclient.ServiceResponse[StartResponseItem]
	if err := c.request.DoRequest(ctx, http.MethodPut, Endpoint+"/start", bytes.NewReader(body), &resp); err != nil {
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