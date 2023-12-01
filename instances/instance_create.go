// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package instances

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"sdk.kraft.cloud/client"
	"sdk.kraft.cloud/services"
)

type CreateInstanceServiceGroupRequest struct {
	// Name of the existing service group.
	Name string `json:"name,omitempty"`

	// UUID of the existing service group.
	UUID string `json:"uuid,omitempty"`

	// The DNS name under which the group is accessible from the internet.  If the
	// DNSName is terminates with a `.` it represents a FQDN, otherwise the
	// provided string will be used as subdomain on the given metro.
	DNSName string `json:"dns_name,omitempty"`

	// Services contains the descriptions of exposed network services.
	Services []services.Service `json:"services,omitempty"`
}

type CreateInstanceVolumeRequest struct {
	// Name of the existing service group.
	Name string `json:"name,omitempty"`

	// UUID of the existing service group.
	UUID string `json:"uuid,omitempty"`

	// Size of the new volume in megabytes.
	SizeMB int `json:"size_mb,omitempty"`

	// Path of the mountpoint. Must be empty. Automatically created if it does not
	// exist.
	At string `json:"at,omitempty"`

	// Whether the volume should be mounted read-only.
	ReadOnly bool `json:"readonly,omitempty"`
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

	// Service group to assign the instance to.
	ServiceGroup CreateInstanceServiceGroupRequest `json:"service_group,omitempty"`

	// Description of volumes
	Volumes []CreateInstanceVolumeRequest `json:"volumes,omitempty"`

	// Autostart behavior. If true the instance will start immediately after
	// creation.
	Autostart bool `json:"autostart,omitempty"`

	// Number of replicas to create with these properties.
	Replicas int `json:"replicas,omitempty"`

	// Key/value pairs to be set as environment variables at boot time.
	// Values must be strings.
	Env map[string]string `json:"env,omitempty"`

	// Name of the created instance. If not set, a random name will be generated.
	Name string `json:"name,omitempty"`
}

// Creates one or more new instances of the specified Unikraft images. You can
// describe the properties of the new instances such as their startup
// arguments and amount of memory. Note that, the instance properties can only
// be defined during creation. They cannot be changed later.
//
// See: https://docs.kraft.cloud/002-rest-api-v1-instances.html#create
func (c *instancesClient) Create(ctx context.Context, req CreateInstanceRequest) (*Instance, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %w", err)
	}

	var response client.ServiceResponse[Instance]
	if err := c.request.DoRequest(ctx, http.MethodPost, Endpoint, bytes.NewBuffer(body), &response); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	instance, err := response.FirstOrErr()
	if instance != nil && instance.Message != "" {
		err = fmt.Errorf("%w: %s", err, instance.Message)
	}
	return instance, err
}
