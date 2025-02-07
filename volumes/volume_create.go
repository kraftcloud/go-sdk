// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package volumes

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	kcclient "sdk.kraft.cloud/client"
	"sdk.kraft.cloud/uuid"
)

type CreateRequestTemplate struct {
	Name *string `json:"name,omitempty"`
	UUID *string `json:"uuid,omitempty"`
}

type CreateRequest struct {
	Name     string                 `json:"name"`
	Template *CreateRequestTemplate `json:"template,omitempty"`
	SizeMb   *int                   `json:"size_mb,omitempty"`
}

// Create implements VolumesService.
func (c *client) Create(ctx context.Context, name string, sizeMB int, template string) (*kcclient.ServiceResponse[CreateResponseItem], error) {
	var err error
	var body []byte
	bodyMap := CreateRequest{}

	if template == "" {
		if sizeMB < 1 {
			return nil, fmt.Errorf("size_mb must be greater than 0")
		}
		bodyMap.SizeMb = &sizeMB
	} else {
		if uuid.IsValid(template) {
			bodyMap.Template = &CreateRequestTemplate{UUID: &template}
		} else {
			bodyMap.Template = &CreateRequestTemplate{Name: &template}
		}
	}
	if name != "" {
		bodyMap.Name = name
	}
	body, err = json.Marshal(bodyMap)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %w", err)
	}

	resp := &kcclient.ServiceResponse[CreateResponseItem]{}
	if err := c.request.DoRequest(ctx, http.MethodPost, Endpoint, bytes.NewReader(body), resp); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	return resp, nil
}
