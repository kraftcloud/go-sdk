// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package volumes

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"sdk.kraft.cloud/client"
)

// Lists all existing service groups. You can filter by persistence and DNS
// name. The latter can be used to lookup the UUID of the service group that
// owns a certain DNS name. The returned groups fulfill all provided filter
// criteria. No particular value is assumed if a filter is not part of the
// request.
//
// The array of groups in the response can be directly fed into the other
// endpoints, for example, to delete (empty) groups.
//
// See: https://docs.kraft.cloud/003-rest-api-v1-services.html#list-existing-service-groups
func (c *volumesClient) List(ctx context.Context) ([]Volume, error) {
	endpoint := Endpoint + "/list"

	// Save the metro such that we can force using it again due to the compromise
	// below.
	metro := c.request.Metro()

	var response client.ServiceResponse[Volume]
	if err := c.request.DoRequest(ctx, http.MethodGet, endpoint, nil, &response); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	// TODO(nderjung): For now, the KraftCloud API does not support
	// returning the full details of each volume.  Temporarily request a
	// status for each volume.
	uuids, err := response.AllOrErr()
	if err != nil {
		var errMsgs []string

		for _, volume := range uuids {
			if volume.Message != "" {
				errMsgs = append(errMsgs, volume.Message)
			}
		}
		return nil, fmt.Errorf("%w: %s", err, strings.Join(errMsgs, ", "))
	}

	var volumes []Volume
	for _, uuid := range uuids {
		instance, err := c.WithMetro(metro).Get(ctx, uuid.UUID)
		if err != nil {
			return nil, fmt.Errorf("could not get instance status: %w", err)
		}

		volumes = append(volumes, *instance)
	}

	return volumes, nil
}
