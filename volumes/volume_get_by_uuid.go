// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package volumes

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"sdk.kraft.cloud/client"
)

// GetByUUID returns the current state and the configuration of a volume.
//
// See: https://docs.kraft.cloud/006-rest-api-v1-volumes.html#state
func (c *volumesClient) GetByUUID(ctx context.Context, uuid string) (*Volume, error) {
	if uuid == "" {
		return nil, errors.New("UUID cannot be empty")
	}

	var response client.ServiceResponse[Volume]
	if err := c.request.DoRequest(ctx, http.MethodGet, Endpoint+"/"+uuid, nil, &response); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	volume, err := response.FirstOrErr()
	if volume != nil && volume.Message != "" {
		err = errors.Join(err, fmt.Errorf(volume.Message))
	}

	return volume, err
}
