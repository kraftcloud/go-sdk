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

// DeletePolicyByName deletes an autoscale policy given its name.
func (c *autoscaleClient) DeletePolicyByName(ctx context.Context, autoscaleUUID, policyName string) (*AutoscaleConfiguration, error) {
	if autoscaleUUID == "" || policyName == "" {
		return nil, errors.New("policyName and autoscaleUUID cannot be empty")
	}

	endpoint := services.Endpoint + "/" + autoscaleUUID + AutoscaleEndpoint + AutoscalePolicyEndpoint + "/" + policyName

	var response client.ServiceResponse[AutoscaleConfiguration]
	if err := c.request.DoRequest(ctx, http.MethodDelete, endpoint, nil, &response); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	autoscaleConf, err := response.FirstOrErr()
	if autoscaleConf != nil && autoscaleConf.Message != "" {
		err = errors.Join(err, fmt.Errorf(autoscaleConf.Message))
	}

	return autoscaleConf, err
}
