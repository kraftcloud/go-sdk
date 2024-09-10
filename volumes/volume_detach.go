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

	ukcclient "sdk.kraft.cloud/client"
	"sdk.kraft.cloud/uuid"
)

// Detach implements VolumesService.
func (c *client) Detach(ctx context.Context, id string, from string) (*ukcclient.ServiceResponse[DetachResponseItem], error) {
	if id == "" {
		return nil, errors.New("identifier cannot be empty")
	}

	reqItem := make(map[string]any, 1)
	if uuid.IsValid(id) {
		reqItem["uuid"] = id
	} else {
		reqItem["name"] = id
	}

	if from != "" {
		if uuid.IsValid(from) {
			reqItem["from"] = InstanceAttachment{UUID: from}
		} else {
			reqItem["from"] = InstanceAttachment{Name: from}
		}
	}

	body, err := json.Marshal([]map[string]any{reqItem})
	if err != nil {
		return nil, fmt.Errorf("encoding JSON object: %w", err)
	}

	resp := &ukcclient.ServiceResponse[DetachResponseItem]{}
	if err := c.request.DoRequest(ctx, http.MethodPut, Endpoint+"/detach", bytes.NewReader(body), resp); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	return resp, nil
}
