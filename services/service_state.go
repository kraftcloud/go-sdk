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

// State returns the current state and the configuration of a service group.
// UUIDs can also be names.
func (c *servicesClient) State(ctx context.Context, uuidOrName string) (*ServiceGroup, error) {
	if uuidOrName == "" {
		return nil, errors.New("UUID cannot be empty")
	}

	endpoint := Endpoint + "/" + uuidOrName

	var response client.ServiceResponse[ServiceGroup]
	if err := c.request.DoRequest(ctx, http.MethodGet, endpoint, nil, &response); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	service, err := response.FirstOrErr()
	if service != nil && service.Message != "" {
		err = fmt.Errorf("%w: %s", err, service.Message)
	}

	return service, err
}
