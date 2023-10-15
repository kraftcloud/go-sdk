// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package instance

import (
	"time"

	kraftcloud "sdk.kraft.cloud"
)

// instancesClient is a basic wrapper around the v1 instance client of
// KraftCloud.
//
// See: https://docs.kraft.cloud/002-rest-api-v1-instances.html
type instancesClient struct {
	defOpts *kraftcloud.Options
	request *kraftcloud.ServiceRequest
}

var _ InstancesService = (*instancesClient)(nil)

// NewInstancesClient instantiates a client which interfaces with KraftCloud's
// instances API.
func NewInstancesClient(opts ...kraftcloud.Option) InstancesService {
	return &instancesClient{
		defOpts: kraftcloud.NewDefaultOptions(opts...),
	}
}

// NewInstancesClientFromOptions instantiates a new instances services client
// based on the provided pre-existing options.
func NewInstancesClientFromOptions(opts *kraftcloud.Options) InstancesService {
	return &instancesClient{
		defOpts: opts,
	}
}

// WithMetro sets the just-in-time metro to use when connecting to the
// KraftCloud API.
func (c *instancesClient) WithMetro(metro string) InstancesService {
	if c.request == nil {
		c.request = kraftcloud.NewServiceRequestFromDefaultOptions(c.defOpts)
	}
	c.request.SetMetro(metro)
	return c
}

// WithHTTPClient overwrites the base HTTP client.
func (c *instancesClient) WithHTTPClient(httpClient kraftcloud.HTTPClient) InstancesService {
	if c.request == nil {
		c.request = kraftcloud.NewServiceRequestFromDefaultOptions(c.defOpts)
	}
	c.request.SetHTTPClient(httpClient)
	return c
}

// WithTimeout sets the timeout when making a request.
func (c *instancesClient) WithTimeout(timeout time.Duration) InstancesService {
	if c.request == nil {
		c.request = kraftcloud.NewServiceRequestFromDefaultOptions(c.defOpts)
	}
	c.request.SetTimeout(timeout)
	return c
}
