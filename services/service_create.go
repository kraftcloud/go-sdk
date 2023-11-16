// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"sdk.kraft.cloud/client"
)

type ServiceCreateRequest struct {
	// Services is a list of descriptiosn of exposed network services.
	Services []Service `json:"services"`

	// Name is the name of the service.
	Name string `json:"name,omitempty"`

	// DNS is the DNS name of the service.
	DNS string `json:"dns,omitempty"`
}

// Creates one or more service groups with the given configuration. Note that,
// the service group properties like published ports can only be defined
// during creation. They cannot be changed later.
//
// Each port in a service group can specify a list of handlers that determine
// how traffic arriving at the port is handled. See Connection Handlers for a
// complete overview.
//
// You can specify an array of service group descriptions to create multiple
// groups with different properties with the same call.
//
// See: https://docs.kraft.cloud/003-rest-api-v1-services.html#creating-new-service-groups
func (c *servicesClient) Create(ctx context.Context, req ServiceCreateRequest) (*ServiceGroup, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %w", err)
	}

	var response client.ServiceResponse[ServiceGroup]
	if err := c.request.DoRequest(ctx, http.MethodPost, Endpoint, bytes.NewBuffer(body), &response); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	return response.FirstOrErr()
}
