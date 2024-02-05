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
)

// GetConfigurationByName returns the current state and the configuration of an autoscale configuration
// given its name.
func (c *autoscaleClient) GetConfigurationByName(ctx context.Context, name string) (*AutoscaleConfiguration, error) {
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}

	body, err := json.Marshal([]map[string]interface{}{{"name": name}})
	if err != nil {
		return nil, fmt.Errorf("encoding JSON object: %w", err)
	}

	var response client.ServiceResponse[AutoscaleConfiguration]
	if err := c.request.DoRequest(ctx, http.MethodGet, Endpoint, bytes.NewBuffer(body), &response); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	autoscaleConf, err := response.FirstOrErr()
	if autoscaleConf != nil && autoscaleConf.Message != "" {
		err = errors.Join(err, fmt.Errorf(autoscaleConf.Message))
	}

	return autoscaleConf, err
}
