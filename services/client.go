// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package services

import (
	"time"

	"sdk.kraft.cloud/client"
	"sdk.kraft.cloud/client/httpclient"
	"sdk.kraft.cloud/client/options"
)

// servicesClient wraps the v1 Services client of KraftCloud.
//
// See: https://docs.kraft.cloud/api/v1/services/
type servicesClient struct {
	// constructors must ensure that request is non-nil
	request *client.ServiceRequest
}

var _ ServicesService = (*servicesClient)(nil)

// NewServicesClientFromOptions instantiates a new image services client based on
// the provided pre-existing options.
func NewServicesClientFromOptions(opts *options.Options) ServicesService {
	return &servicesClient{
		request: client.NewServiceRequestFromDefaultOptions(opts),
	}
}

// WithMetro sets the just-in-time metro to use when connecting to the
// KraftCloud API.
func (c *servicesClient) WithMetro(m string) ServicesService {
	ccpy := c.clone()
	ccpy.request = c.request.WithMetro(m)
	return ccpy
}

// WithHTTPClient overwrites the base HTTP client.
func (c *servicesClient) WithHTTPClient(hc httpclient.HTTPClient) ServicesService {
	ccpy := c.clone()
	ccpy.request = c.request.WithHTTPClient(hc)
	return ccpy
}

// WithTimeout sets the timeout when making a request.
func (c *servicesClient) WithTimeout(to time.Duration) ServicesService {
	ccpy := c.clone()
	ccpy.request = c.request.WithTimeout(to)
	return ccpy
}

// clone returns a shallow copy of c.
func (c *servicesClient) clone() *servicesClient {
	ccpy := *c
	return &ccpy
}
