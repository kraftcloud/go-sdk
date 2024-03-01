// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2024, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package services

import (
	"time"

	"sdk.kraft.cloud/client"
	"sdk.kraft.cloud/client/httpclient"
	"sdk.kraft.cloud/client/options"
)

// autoscaleClient wraps the v1 Autoscale client of KraftCloud.
//
// See: https://docs.kraft.cloud/api/v1/autoscale/
type autoscaleClient struct {
	// constructors must ensure that request is non-nil
	request *client.ServiceRequest
}

var _ AutoscaleService = (*autoscaleClient)(nil)

// NewAutoscaleClientFromOptions instantiates a new autoscale services client based on
// the provided pre-existing options.
func NewAutoscaleClientFromOptions(opts *options.Options) AutoscaleService {
	return &autoscaleClient{
		request: client.NewServiceRequestFromDefaultOptions(opts),
	}
}

// WithMetro sets the just-in-time metro to use when connecting to the
// KraftCloud API.
func (c *autoscaleClient) WithMetro(m string) AutoscaleService {
	ccpy := c.clone()
	ccpy.request = c.request.WithMetro(m)
	return ccpy
}

// WithHTTPClient overwrites the base HTTP client.
func (c *autoscaleClient) WithHTTPClient(hc httpclient.HTTPClient) AutoscaleService {
	ccpy := c.clone()
	ccpy.request = c.request.WithHTTPClient(hc)
	return ccpy
}

// WithTimeout sets the timeout when making a request.
func (c *autoscaleClient) WithTimeout(to time.Duration) AutoscaleService {
	ccpy := c.clone()
	ccpy.request = c.request.WithTimeout(to)
	return ccpy
}

// clone returns a shallow copy of c.
func (c *autoscaleClient) clone() *autoscaleClient {
	ccpy := *c
	return &ccpy
}
