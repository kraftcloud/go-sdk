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

type VolumeCreateRequest struct {
	// Name of the volume.
	Name string `json:"name,omitempty"`

	// Size of the volume in megabytes.
	SizeMB int `json:"size_mb,omitempty"`
}

// Creates one or more volumes with the given configuration. The volumes are
// automatically initialized with an empty file system. After initialization
// completed the volumes are in the available state and can be attached to an
// instance with the attach endpoint. Note that, the size of a volume cannot
// be changed after creation.
//
// See: https://docs.kraft.cloud/006-rest-api-v1-volumes.html#create
func (c *volumesClient) Create(ctx context.Context, req VolumeCreateRequest) (*Volume, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %w", err)
	}

	var response client.ServiceResponse[Volume]
	if err := c.request.DoRequest(ctx, http.MethodPost, Endpoint, bytes.NewBuffer(body), &response); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	volume, err := response.FirstOrErr()
	if volume != nil && volume.Message != "" {
		err = fmt.Errorf("%w: %s", err, volume.Message)
	}

	return volume, err
}
