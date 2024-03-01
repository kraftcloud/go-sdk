// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package volumes

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"sdk.kraft.cloud/client"
)

// DeleteByName the specified volume given its UUID. Fails if the volume is
// still attached to an instance. After this call the UUID of the volumes is no
// longer valid.
//
// See: https://docs.kraft.cloud/api/v1/volumes/#delete
func (c *volumesClient) DeleteByName(ctx context.Context, name string) error {
	if name == "" {
		return errors.New("name cannot be empty")
	}

	body, err := json.Marshal([]map[string]interface{}{{"name": name}})
	if err != nil {
		return fmt.Errorf("encoding JSON object: %w", err)
	}

	var response client.ServiceResponse[Volume]
	if err := c.request.DoRequest(ctx, http.MethodDelete, Endpoint, bytes.NewBuffer(body), &response); err != nil {
		return fmt.Errorf("performing the request: %w", err)
	}

	volume, err := response.FirstOrErr()
	if volume != nil && volume.Message != "" {
		err = errors.Join(err, fmt.Errorf(volume.Message))
	}

	return err
}
