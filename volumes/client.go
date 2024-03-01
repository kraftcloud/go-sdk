// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package volumes

import (
	"time"

	"sdk.kraft.cloud/client"
	"sdk.kraft.cloud/client/httpclient"
	"sdk.kraft.cloud/client/options"
)

// volumesClient wraps the v1 Volumes client of KraftCloud.
//
// See: https://docs.kraft.cloud/api/v1/volumes/
type volumesClient struct {
	// constructors must ensure that request is non-nil
	request *client.ServiceRequest
}

var _ VolumesService = (*volumesClient)(nil)

// NewVolumesClientFromOptions instantiates a new users services client based on
// the provided pre-existing options.
func NewVolumesClientFromOptions(opts *options.Options) VolumesService {
	return &volumesClient{
		request: client.NewServiceRequestFromDefaultOptions(opts),
	}
}

// WithMetro sets the just-in-time metro to use when connecting to the
// KraftCloud API.
func (c *volumesClient) WithMetro(m string) VolumesService {
	ccpy := c.clone()
	ccpy.request = c.request.WithMetro(m)
	return ccpy
}

// WithHTTPClient overwrites the base HTTP client.
func (c *volumesClient) WithHTTPClient(hc httpclient.HTTPClient) VolumesService {
	ccpy := c.clone()
	ccpy.request = c.request.WithHTTPClient(hc)
	return ccpy
}

// WithTimeout sets the timeout when making a request.
func (c *volumesClient) WithTimeout(to time.Duration) VolumesService {
	ccpy := c.clone()
	ccpy.request = c.request.WithTimeout(to)
	return ccpy
}

// clone returns a shallow copy of c.
func (c *volumesClient) clone() *volumesClient {
	ccpy := *c
	return &ccpy
}
