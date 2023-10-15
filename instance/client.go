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
	// constructors must ensure that request is non-nil
	request *kraftcloud.ServiceRequest
}

var _ InstancesService = (*instancesClient)(nil)

// NewInstancesClient instantiates a client which interfaces with KraftCloud's
// instances API.
func NewInstancesClient(opts ...kraftcloud.Option) InstancesService {
	return &instancesClient{
		request: kraftcloud.NewServiceRequestFromDefaultOptions(kraftcloud.NewDefaultOptions(opts...)),
	}
}

// NewInstancesClientFromOptions instantiates a new instances services client
// based on the provided pre-existing options.
func NewInstancesClientFromOptions(opts *kraftcloud.Options) InstancesService {
	return &instancesClient{
		request: kraftcloud.NewServiceRequestFromDefaultOptions(opts),
	}
}

// WithMetro sets the just-in-time metro to use when connecting to the
// KraftCloud API.
func (c *instancesClient) WithMetro(m string) InstancesService {
	ccpy := c.clone()
	ccpy.request = c.request.WithMetro(m)
	return ccpy
}

// WithHTTPClient overwrites the base HTTP client.
func (c *instancesClient) WithHTTPClient(hc kraftcloud.HTTPClient) InstancesService {
	ccpy := c.clone()
	ccpy.request = c.request.WithHTTPClient(hc)
	return ccpy
}

// WithTimeout sets the timeout when making a request.
func (c *instancesClient) WithTimeout(to time.Duration) InstancesService {
	ccpy := c.clone()
	ccpy.request = c.request.WithTimeout(to)
	return ccpy
}

// clone returns a shallow copy of c.
func (c *instancesClient) clone() *instancesClient {
	ccpy := *c
	return &ccpy
}
