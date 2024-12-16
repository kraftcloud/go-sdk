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

// Stop implements InstancesService.
func (c *client) Stop(ctx context.Context, drainTimeoutMs int, force bool, ids ...string) (*ukcclient.ServiceResponse[StopResponseItem], error) {
	if len(ids) == 0 {
		return nil, errors.New("requires at least one identifier")
	}

	reqItems := make([]map[string]any, 0, len(ids))
	for _, id := range ids {
		reqItem := make(map[string]any, 3)
		if uuid.IsValid(id) {
			reqItem["uuid"] = id
		} else {
			reqItem["name"] = id
		}
		reqItem["force"] = force
		if drainTimeoutMs > 0 {
			reqItem["drain_timeout_ms"] = drainTimeoutMs
		}
		reqItems = append(reqItems, reqItem)
	}

	body, err := json.Marshal(reqItems)
	if err != nil {
		return nil, fmt.Errorf("encoding JSON object: %w", err)
	}

	resp := &ukcclient.ServiceResponse[StopResponseItem]{}
	if err := c.request.DoRequest(ctx, http.MethodPut, Endpoint+"/stop", bytes.NewReader(body), resp); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	return resp, nil
}
