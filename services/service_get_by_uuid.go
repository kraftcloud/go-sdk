// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package services

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"sdk.kraft.cloud/client"
)

// GetByUUID returns the current state and the configuration of a service group
// given its UUID.
func (c *servicesClient) GetByUUID(ctx context.Context, uuid string) (*ServiceGroup, error) {
	if uuid == "" {
		return nil, errors.New("UUID cannot be empty")
	}

	endpoint := Endpoint + "/" + uuid

	var response client.ServiceResponse[ServiceGroup]
	if err := c.request.DoRequest(ctx, http.MethodGet, endpoint, nil, &response); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	service, err := response.FirstOrErr()
	if service != nil && service.Message != "" {
		err = errors.Join(err, fmt.Errorf(service.Message))
	}

	return service, err
}
