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
	"sdk.kraft.cloud/uuid"
)

// Attach implements VolumesService.
func (c *client) Attach(ctx context.Context, volID, instanceUUID, at string, readOnly bool) (*kcclient.ServiceResponse[AttachResponseItem], error) {
	if volID == "" {
		return nil, errors.New("volume identifier cannot be empty")
	}
	if instanceUUID == "" {
		return nil, errors.New("instance UUID cannot be empty")
	}
	if at == "" {
		return nil, errors.New("at cannot be empty")
	}

	reqItem := make(map[string]any, 4)
	if uuid.IsValid(volID) {
		reqItem["uuid"] = volID
	} else {
		reqItem["name"] = volID
	}
	reqItem["at"] = at
	reqItem["readonly"] = readOnly
	reqItem["attach_to"] = map[string]any{
		"uuid": instanceUUID,
	}

	body, err := json.Marshal([]map[string]any{reqItem})
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %w", err)
	}

	resp := &kcclient.ServiceResponse[AttachResponseItem]{}
	if err := c.request.DoRequest(ctx, http.MethodPut, Endpoint+"/attach", bytes.NewReader(body), resp); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	return resp, nil
}
