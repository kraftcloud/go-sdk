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

// Creates one or more volumes with the given configuration. The volumes are
// automatically initialized with an empty file system. After initialization
// completed the volumes are in the available state and can be attached to an
// instance with the attach endpoint. Note that, the size of a volume cannot
// be changed after creation.
//
// See: https://docs.kraft.cloud/api/v1/volumes/#create
func (c *volumesClient) Create(ctx context.Context, name string, sizeMB int) (*Volume, error) {
	body, err := json.Marshal(map[string]interface{}{
		"name":    name,
		"size_mb": sizeMB,
	})
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %w", err)
	}

	var response client.ServiceResponse[Volume]
	if err := c.request.DoRequest(ctx, http.MethodPost, Endpoint, bytes.NewBuffer(body), &response); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	volume, err := response.FirstOrErr()
	if volume != nil && volume.Message != "" {
		err = errors.Join(err, fmt.Errorf(volume.Message))
	}

	return volume, err
}
