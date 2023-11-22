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

type VolumeAttachRequest struct {
	// UUID of the instance to attach the volume to.
	AttachTo VolumeAttachedToInstance `json:"attach_to"`

	// Path of the mountpoint.
	At string `json:"at"`

	// Whether the volume should be mounted read-only
	ReadOnly bool `json:"readonly,omitempty"`
}

// Attaches a volume to an instance so that the volume is mounted when the
// instance starts. The volume needs to be in available state and the instance
// must in stopped state. Currently, each instance can have only one volume
// attached at most.
//
// See: https://docs.kraft.cloud/006-rest-api-v1-volumes.html#attach
func (c *volumesClient) Attach(ctx context.Context, uuidOrName string, req VolumeAttachRequest) (*Volume, error) {
	if uuidOrName == "" {
		return nil, errors.New("UUID cannot be empty")
	}

	if req.AttachTo.UUID == "" && req.AttachTo.Name == "" {
		return nil, errors.New("attach_to UUID or name cannot be empty")
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %w", err)
	}

	endpoint := Endpoint + "/" + uuidOrName + "/attach"

	var response client.ServiceResponse[Volume]
	if err := c.request.DoRequest(ctx, http.MethodPut, endpoint, bytes.NewBuffer(body), &response); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	volume, err := response.FirstOrErr()
	if volume != nil && volume.Message != "" {
		err = fmt.Errorf("%w: %s", err, volume.Message)
	}

	return volume, err
}
