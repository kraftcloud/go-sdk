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

// DetachByUUID detaches a volume from an instance by its UUID. The instance
// from which to detach must in stopped state. If the volume has been created
// together with an instance, detaching the volume will make it persistent
// (i.e., it survives the deletion of the instance).
//
// See: https://docs.kraft.cloud/api/v1/volumes/#detach
func (c *volumesClient) DetachByUUID(ctx context.Context, uuid string) (*Volume, error) {
	if uuid == "" {
		return nil, errors.New("UUID cannot be empty")
	}

	var response client.ServiceResponse[Volume]
	if err := c.request.DoRequest(ctx, http.MethodPut, Endpoint+"/"+uuid+"/detach", nil, &response); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	volume, err := response.FirstOrErr()
	if volume != nil && volume.Message != "" {
		err = errors.Join(err, fmt.Errorf(volume.Message))
	}

	return volume, err
}
