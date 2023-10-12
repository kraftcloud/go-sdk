// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package instance

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"sdk.kraft.cloud/v0/util"
)

// CreateInstanceRequest holds all the data necessary to create an instance via
// the API.
//
// See: https://docs.kraft.cloud/002-rest-api-v1-instances.html#create
type CreateInstanceRequest struct {
	// Name of the Unikraft image to instantiate. Private images will be available
	// under your user's namespace.
	Image string

	// Application arguments.
	Args []string

	// Amount of memory to assign to the instance in megabytes.
	Memory int64

	// Connection handlers. Must be [ "tls" ].
	Handlers []string

	// Public-facing Port.
	Port int64

	// Port that the image listens on.
	InternalPort int64

	// Autostart behavior. If true the instance will start immediately after
	// creation.
	Autostart bool
}

// Creates one or more new instances of the specified Unikraft images. You can
// describe the properties of the new instances such as their startup
// arguments and amount of memory. Note that, the instance properties can only
// be defined during creation. They cannot be changed later.
//
// See: https://docs.kraft.cloud/002-rest-api-v1-instances.html#create
func (i *InstanceClient) Create(ctx context.Context, req CreateInstanceRequest) (*Instance, error) {
	// normalize into the from kraftcloud API expects:
	image, err := util.NormalizeImageName(req.Image)
	if err != nil {
		return nil, fmt.Errorf("normalizing image name: %w", err)
	}

	requestBody := map[string]interface{}{
		"image":     image,
		"args":      req.Args,
		"memory_mb": req.Memory,
		"services": []map[string]interface{}{
			{
				"port":          req.Port,
				"internal_port": req.InternalPort,
				"handlers":      req.Handlers,
			},
		},
		"autostart": req.Autostart,
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("marshalling request body: %w", err)
	}

	endpoint := i.BaseURL + Endpoint

	var response InstanceResponse
	if err := i.DoRequest(ctx, http.MethodPost, endpoint, bytes.NewBuffer(body), &response); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	return firstInstanceOrErr(&response)
}
