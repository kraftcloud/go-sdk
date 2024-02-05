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

// DeleteConfigurationByUUID deletes an autoscale configuration given its UUID.
func (c *autoscaleClient) DeleteConfigurationByUUID(ctx context.Context, uuid string) error {
	if uuid == "" {
		return errors.New("UUID cannot be empty")
	}

	endpoint := services.Endpoint + "/" + uuid + AutoscaleEndpoint

	var response client.ServiceResponse[AutoscaleConfiguration]
	if err := c.request.DoRequest(ctx, http.MethodDelete, endpoint, nil, &response); err != nil {
		return fmt.Errorf("performing the request: %w", err)
	}

	service, err := response.FirstOrErr()
	if service != nil && service.Message != "" {
		err = errors.Join(err, fmt.Errorf(service.Message))
	}

	return err
}
