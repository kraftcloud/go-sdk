// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package volumes

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	kcclient "sdk.kraft.cloud/client"
)

// AttachByName implements VolumesService.
func (c *client) AttachByName(ctx context.Context, volName, instanceName, at string, readOnly bool) (*AttachResponseItem, error) {
	if volName == "" {
		return nil, errors.New("volume name cannot be empty")
	}
	if instanceName == "" {
		return nil, errors.New("instance name cannot be empty")
	}
	if at == "" {
		return nil, errors.New("destination at cannot be empty")
	}

	body, err := json.Marshal(map[string]interface{}{
		"at":       at,
		"name":     volName,
		"readonly": readOnly,
		"attach_to": map[string]interface{}{
			"name": instanceName,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %w", err)
	}

	var resp kcclient.ServiceResponse[AttachResponseItem]
	if err := c.request.DoRequest(ctx, http.MethodPut, Endpoint+"/attach", bytes.NewBuffer(body), &resp); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	item, err := resp.FirstOrErr()
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("%s (code=%d)", item.Message, *item.Error))
	}
	return item, nil
}
