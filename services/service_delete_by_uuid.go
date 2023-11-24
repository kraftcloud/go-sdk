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

// DeleteByUUID the specified service group given its UUID. Fails if there are
// still instances attached to group. After this call the UUID of the group is
// no longer valid.
//
// This operation cannot be undone.
//
// See:
// https://docs.kraft.cloud/003-rest-api-v1-services.html#deleting-a-service-group
func (c *servicesClient) DeleteByUUID(ctx context.Context, uuid string) error {
	if uuid == "" {
		return errors.New("UUID cannot be empty")
	}

	var response client.ServiceResponse[ServiceGroup]
	if err := c.request.DoRequest(ctx, http.MethodDelete, Endpoint+"/"+uuid, nil, &response); err != nil {
		return fmt.Errorf("performing the request: %w", err)
	}

	service, err := response.FirstOrErr()
	if service != nil && service.Message != "" {
		err = errors.Join(err, fmt.Errorf(service.Message))
	}

	return err
}
