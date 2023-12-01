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

// DeleteByUUID the specified volume given its UUID. Fails if the volume is
// still attached to an instance. After this call the UUID of the volumes is no
// longer valid.
//
// See: https://docs.kraft.cloud/006-rest-api-v1-volumes.html#delete
func (c *volumesClient) DeleteByUUID(ctx context.Context, uuidOrName string) error {
	if uuidOrName == "" {
		return errors.New("UUID or Name cannot be empty")
	}

	var response client.ServiceResponse[Volume]
	if err := c.request.DoRequest(ctx, http.MethodDelete, Endpoint+"/"+uuidOrName, nil, &response); err != nil {
		return fmt.Errorf("performing the request: %w", err)
	}

	volume, err := response.FirstOrErr()
	if volume != nil && volume.Message != "" {
		err = errors.Join(err, fmt.Errorf(volume.Message))
	}

	return err
}
