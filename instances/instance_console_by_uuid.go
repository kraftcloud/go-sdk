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

// ConsoleByUUID implements InstancesService.
func (c *client) ConsoleByUUID(ctx context.Context, uuid string, maxLines int, latest bool) (*ConsoleResponseItem, error) {
	if len(uuid) == 0 {
		return nil, fmt.Errorf("UUID cannot be empty")
	}

	body, err := json.Marshal([]map[string]interface{}{{
		"max_lines": maxLines,
	}})
	if err != nil {
		return nil, fmt.Errorf("marshalling request body: %w", err)
	}

	var resp kcclient.ServiceResponse[ConsoleResponseItem]
	if err := c.request.DoRequest(ctx, http.MethodGet, Endpoint+"/"+uuid+"/console", bytes.NewBuffer(body), &resp); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	item, err := resp.FirstOrErr()
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("%s (code=%d)", item.Message, *item.Error))
	}
	return item, nil
}