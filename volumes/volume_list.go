// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package volumes

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

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
// See: https://docs.kraft.cloud/api/v1/volumes/#list
func (c *volumesClient) List(ctx context.Context) ([]Volume, error) {
	var response client.ServiceResponse[Volume]
	if err := c.request.DoRequest(ctx, http.MethodGet, Endpoint+"/list", nil, &response); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	volumes, err := response.AllOrErr()
	if err != nil {
		return nil, err
	}

	// TODO(nderjung): For now, the KraftCloud API does not support
	// returning the full details of each volume.  Temporarily request a
	// status for each volume.

	req := make([]map[string]interface{}, len(volumes))
	for i, instance := range volumes {
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

	return volumes, nil
}
