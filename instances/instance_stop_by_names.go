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

// StopByNames implements InstancesService.
func (c *client) StopByNames(ctx context.Context, drainTimeoutMs int, force bool, names ...string) (*kcclient.ServiceResponse[StopResponseItem], error) {
	if len(names) == 0 {
		return nil, errors.New("requires at least one name")
	}

	reqItems := make([]map[string]any, 0, len(names))
	for _, name := range names {
		reqItem := map[string]any{
			"name":  name,
			"force": force,
		}
		if drainTimeoutMs > 0 {
			reqItem["drain_timeout_ms"] = drainTimeoutMs
		}
		reqItems = append(reqItems, reqItem)
	}

	body, err := json.Marshal(reqItems)
	if err != nil {
		return nil, fmt.Errorf("encoding JSON object: %w", err)
	}

	resp := &kcclient.ServiceResponse[StopResponseItem]{}
	if err := c.request.DoRequest(ctx, http.MethodPut, Endpoint+"/stop", bytes.NewReader(body), resp); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	return resp, nil
}
