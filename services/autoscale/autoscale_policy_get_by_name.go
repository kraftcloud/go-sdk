// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2024, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package services

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"sdk.kraft.cloud/client"
	"sdk.kraft.cloud/services"
)

// GetPolicyByName returns the current state and the configuration of an autoscale policy
func (c *autoscaleClient) GetPolicyByName(ctx context.Context, uuid, name string) (*map[string]interface{}, error) {
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}

	var response client.ServiceResponse[map[string]interface{}]
	if err := c.request.DoRequest(ctx, http.MethodGet, services.Endpoint+"/"+uuid+"/autoscale/policies/"+name, nil, &response); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	autoscaleConf, err := response.FirstOrErr()
	if autoscaleConf != nil {
		conf := *autoscaleConf
		if message, ok := conf["message"]; ok {
			err = errors.Join(err, fmt.Errorf(message.(string)))
		}
	}

	return autoscaleConf, err
}
