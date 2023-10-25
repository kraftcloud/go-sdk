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

// Deletes the specified service group. Fails if there are still instances
// attached to group. After this call the UUID of the group is no longer
// valid.
//
// This operation cannot be undone.
//
// See: https://docs.kraft.cloud/003-rest-api-v1-services.html#deleting-a-service-group
func (c *servicesClient) Delete(ctx context.Context, uuid string) (*ServiceGroup, error) {
	if uuid == "" {
		return nil, errors.New("UUID cannot be empty")
	}

	endpoint := Endpoint + "/" + uuid

	var response client.ServiceResponse[ServiceGroup]
	if err := c.request.DoRequest(ctx, http.MethodDelete, endpoint, nil, &response); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	return response.FirstOrErr()
}
