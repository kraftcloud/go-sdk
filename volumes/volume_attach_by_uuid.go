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

// AttachByUUID a volume to an instance so that the volume is mounted when the
// instance starts using the volume and instance name.  The volume needs to be in
// available state and the instance must in stopped state. Currently, each
// instance can have only one volume attached at most.
//
// See: https://docs.kraft.cloud/api/v1/volumes/#attach
func (c *volumesClient) AttachByUUID(ctx context.Context, volUUID, instanceUUID, at string, readOnly bool) (*Volume, error) {
	if volUUID == "" {
		return nil, errors.New("volume name cannot be empty")
	}
	if instanceUUID == "" {
		return nil, errors.New("instance name cannot be empty")
	}
	if at == "" {
		return nil, errors.New("destination at cannot be empty")
	}

	body, err := json.Marshal(map[string]interface{}{
		"at":       at,
		"uuid":     volUUID,
		"readonly": readOnly,
		"attach_to": map[string]interface{}{
			"uuid": instanceUUID,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %w", err)
	}

	var response client.ServiceResponse[Volume]
	if err := c.request.DoRequest(ctx, http.MethodPut, Endpoint+"/attach", bytes.NewBuffer(body), &response); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	volume, err := response.FirstOrErr()
	if volume != nil && volume.Message != "" {
		err = errors.Join(err, fmt.Errorf(volume.Message))
	}

	return volume, err
}
