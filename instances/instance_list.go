// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package instances

import (
	"context"
	"fmt"
	"net/http"

	"sdk.kraft.cloud/client"
)

// Lists all existing instances.
//
// See: https://docs.kraft.cloud/002-rest-api-v1-instances.html#list
func (c *instancesClient) List(ctx context.Context) ([]Instance, error) {
	endpoint := Endpoint + "/list"

	// Save the metro such that we can force using it again due to the compromise
	// below.
	metro := c.request.Metro()

	var response client.ServiceResponse[Instance]
	if err := c.request.DoRequest(ctx, http.MethodGet, endpoint, nil, &response); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	// TODO(nderjung): For now, the KraftCloud API does not support
	// returning the full details of each instance.  Temporarily request a
	// status for each instance.
	uuids, err := response.AllOrErr()
	if err != nil {
		return nil, err
	}

	var instances []Instance
	for _, uuid := range uuids {
		instance, err := c.WithMetro(metro).State(ctx, uuid.UUID)
		if err != nil {
			return nil, fmt.Errorf("could not get instance status: %w", err)
		}

		instances = append(instances, *instance)
	}

	return instances, nil
}
