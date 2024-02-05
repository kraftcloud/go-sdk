// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2024, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package services

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"sdk.kraft.cloud/client"
	"sdk.kraft.cloud/services"
)

// CreatePolicy creates a new autoscale policy for an autoscale configuration.
func (c *autoscaleClient) CreatePolicy(ctx context.Context, uuid string, typ AutoscalePolicyType, req interface{}) (*AutoscaleConfiguration, error) {
	if uuid == "" {
		return nil, errors.New("UUID cannot be empty")
	}

	asJSON, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON into a map
	var asMap map[string]interface{}
	err = json.Unmarshal(asJSON, &asMap)
	if err != nil {
		return nil, err
	}

	// Dynamically inject the "type"
	asMap["type"] = typ

	// Re-marshal the map into JSON
	body, err := json.Marshal(asMap)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %w", err)
	}

	var response client.ServiceResponse[AutoscaleConfiguration]
	if err := c.request.DoRequest(ctx, http.MethodPost, services.Endpoint+"/"+uuid+"/autoscale/policies", bytes.NewBuffer(body), &response); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	service, err := response.FirstOrErr()
	if service != nil && service.Message != "" {
		err = errors.Join(err, fmt.Errorf(service.Message))
	}

	return service, err
}
