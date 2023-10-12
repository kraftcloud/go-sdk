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

	kraftcloud "sdk.kraft.cloud/v0"
	"sdk.kraft.cloud/v0/util"
)

// CreateInstanceServicesRequest contains the description of an exposed network
// service.
//
// See: https://docs.kraft.cloud/002-rest-api-v1-instances.html#create
type CreateInstanceServicesRequest struct {
	// Public-facing Port
	Port int `json:"port,omitempty"`

	// Port that the image listens on.
	InternalPort int `json:"internal_port,omitempty"`

	// Connection handlers. Must be [ "tls" ].
	Handlers []string `json:"handlers,omitempty"`
}

// CreateInstanceRequest holds all the data necessary to create an instance via
// the API.
//
// See: https://docs.kraft.cloud/002-rest-api-v1-instances.html#create
type CreateInstanceRequest struct {
	// Name of the Unikraft image to instantiate. Private images will be available
	// under your user's namespace.
	Image string `json:"image,omitempty"`

	// Application arguments.
	Args []string `json:"args,omitempty"`

	// Amount of memory to assign to the instance in megabytes.
	MemoryMB int64 `json:"memory_mb,omitempty"`

	// Description of exposed network services.
	Services []CreateInstanceServicesRequest `json:"services,omitempty"`

	// Autostart behavior. If true the instance will start immediately after
	// creation.
	Autostart bool `json:"autostart,omitempty"`
}

// Creates one or more new instances of the specified Unikraft images. You can
// describe the properties of the new instances such as their startup
// arguments and amount of memory. Note that, the instance properties can only
// be defined during creation. They cannot be changed later.
//
// See: https://docs.kraft.cloud/002-rest-api-v1-instances.html#create
func (i *instancesClient) Create(ctx context.Context, req CreateInstanceRequest) (*Instance, error) {
	// normalize into the from kraftcloud API expects:
	image, err := util.NormalizeImageName(req.Image)
	if err != nil {
		return nil, fmt.Errorf("normalizing image name: %w", err)
	}

	req.Image = image

	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %w", err)
	}

	if i.request == nil {
		i.request = kraftcloud.NewServiceRequestFromDefaultOptions(i.opts)
	}

	defer func() { i.request = nil }()

	var response kraftcloud.ServiceResponse[Instance]
	if err := i.request.DoRequest(ctx, http.MethodPost, Endpoint, bytes.NewBuffer(body), &response); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	return response.FirstOrErr()
}
