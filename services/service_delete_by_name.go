// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
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

// DeleteByName the specified service group given its name. Fails if there are
// still instances attached to group. After this call the UUID of the group is
// no longer valid.
//
// This operation cannot be undone.
//
// See:
// https://docs.kraft.cloud/003-rest-api-v1-services.html#deleting-a-service-group
func (c *servicesClient) DeleteByName(ctx context.Context, name string) error {
	if name == "" {
		return errors.New("name cannot be empty")
	}

	body, err := json.Marshal([]map[string]interface{}{{
		"name": name,
	}})
	if err != nil {
		return fmt.Errorf("marshalling request body: %w", err)
	}

	var response client.ServiceResponse[ServiceGroup]
	if err := c.request.DoRequest(ctx, http.MethodDelete, Endpoint, bytes.NewBuffer(body), &response); err != nil {
		return fmt.Errorf("performing the request: %w", err)
	}

	service, err := response.FirstOrErr()
	if service != nil && service.Message != "" {
		err = errors.Join(err, fmt.Errorf(service.Message))
	}

	return err
}
