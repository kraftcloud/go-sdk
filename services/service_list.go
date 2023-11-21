// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package services

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
func (c *servicesClient) List(ctx context.Context) ([]ServiceGroup, error) {
	endpoint := Endpoint + "/list"

	var response client.ServiceResponse[ServiceGroup]
	if err := c.request.DoRequest(ctx, http.MethodGet, endpoint, nil, &response); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	services, err := response.AllOrErr()

	var errMsgs []string
	for _, service := range services {
		if service.Message != "" {
			errMsgs = append(errMsgs, service.Message)
		}
	}

	return nil, fmt.Errorf("%w: %s", err, strings.Join(errMsgs, ", "))
}
