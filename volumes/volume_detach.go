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

// Detaches a volume from an instance. The instance from which to detach must
// in stopped state. If the volume has been created together with an instance,
// detaching the volume will make it persistent (i.e., it survives the
// deletion of the instance).
//
// See: https://docs.kraft.cloud/006-rest-api-v1-volumes.html#detach
func (c *volumesClient) Detach(ctx context.Context, uuid string) (*Volume, error) {
	if uuid == "" {
		return nil, errors.New("UUID cannot be empty")
	}

	endpoint := Endpoint + "/" + uuid + "/detach"

	var response client.ServiceResponse[Volume]
	if err := c.request.DoRequest(ctx, http.MethodPut, endpoint, nil, &response); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	volume, err := response.FirstOrErr()
	if volume != nil && volume.Message != "" {
		err = fmt.Errorf("%w: %s", err, volume.Message)
	}

	return volume, err
}
