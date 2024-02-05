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

// CreateConfigurationByName creates a new autoscale configuration with a service group name.
func (c *autoscaleClient) CreateConfigurationByName(ctx context.Context, name string, req AutoscaleConfiguration) (*services.ServiceGroup, error) {
	if req.Master == nil {
		return nil, errors.New("master cannot be nil")
	}

	if req.Master.UUID == "" && req.Master.Name == "" {
		return nil, errors.New("master name and uuid cannot be both empty")
	}

	if req.MaxSize == 0 {
		req.MaxSize = 10
	}

	if req.WarmupTimeMs == 0 {
		req.WarmupTimeMs = 1000
	}

	if req.CooldownTimeMs == 0 {
		req.CooldownTimeMs = 1000
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %w", err)
	}

	var response client.ServiceResponse[services.ServiceGroup]
	if err := c.request.DoRequest(ctx, http.MethodPost, Endpoint, bytes.NewBuffer(body), &response); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	service, err := response.FirstOrErr()
	if service != nil && service.Message != "" {
		err = errors.Join(err, fmt.Errorf(service.Message))
	}

	return service, err
}
