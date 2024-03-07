// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2024, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package autoscale

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	kcclient "sdk.kraft.cloud/client"
)

// DeleteConfigurationByName implements AutoscaleService.
func (c *client) DeleteConfigurationByName(ctx context.Context, name string) (*DeleteResponseItem, error) {
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}

	body, err := json.Marshal([]map[string]interface{}{{
		"name": name,
	}})
	if err != nil {
		return nil, fmt.Errorf("marshalling request body: %w", err)
	}

	var resp kcclient.ServiceResponse[DeleteResponseItem]
	if err := c.request.DoRequest(ctx, http.MethodDelete, Endpoint, bytes.NewBuffer(body), &resp); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	item, err := resp.FirstOrErr()
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("%s (code=%d)", item.Message, *item.Error))
	}
	return item, nil
}
