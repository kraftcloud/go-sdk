// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package instances

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"sdk.kraft.cloud/client"
)

// Lists all existing instances.
//
// See: https://docs.kraft.cloud/002-rest-api-v1-instances.html#list
func (c *instancesClient) List(ctx context.Context) ([]Instance, error) {
	var response client.ServiceResponse[Instance]
	if err := c.request.DoRequest(ctx, http.MethodGet, Endpoint+"/list", nil, &response); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	instances, err := response.AllOrErr()
	if err != nil {
		return nil, err
	}

	// TODO(nderjung): For now, the KraftCloud API does not support
	// returning the full details of each instance.  Temporarily request a
	// status for each instance.

	req := make([]map[string]interface{}, len(instances))
	for i, instance := range instances {
		req[i] = map[string]interface{}{
			"uuid": instance.UUID,
		}
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("marshalling request body: %w", err)
	}

	if err := c.request.DoRequest(ctx, http.MethodGet, Endpoint, bytes.NewBuffer(body), &response); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	return instances, nil
}
