// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package users

import (
	"time"

	"sdk.kraft.cloud/client"
	"sdk.kraft.cloud/client/httpclient"
	"sdk.kraft.cloud/client/options"
)

// usersClient wraps the v1 Image client of KraftCloud.
//
// See: https://docs.kraft.cloud/003-rest-api-v1-services.html
type usersClient struct {
	// constructors must ensure that request is non-nil
	request *client.ServiceRequest
}

var _ UsersService = (*usersClient)(nil)

// NewUsersClientFromOptions instantiates a new users services client based on
// the provided pre-existing options.
func NewUsersClientFromOptions(opts *options.Options) UsersService {
	return &usersClient{
		request: client.NewServiceRequestFromDefaultOptions(opts),
	}
}

// WithMetro sets the just-in-time metro to use when connecting to the
// KraftCloud API.
func (c *usersClient) WithMetro(m string) UsersService {
	ccpy := c.clone()
	ccpy.request = c.request.WithMetro(m)
	return ccpy
}

// WithHTTPClient overwrites the base HTTP client.
func (c *usersClient) WithHTTPClient(hc httpclient.HTTPClient) UsersService {
	ccpy := c.clone()
	ccpy.request = c.request.WithHTTPClient(hc)
	return ccpy
}

// WithTimeout sets the timeout when making a request.
func (c *usersClient) WithTimeout(to time.Duration) UsersService {
	ccpy := c.clone()
	ccpy.request = c.request.WithTimeout(to)
	return ccpy
}

// clone returns a shallow copy of c.
func (c *usersClient) clone() *usersClient {
	ccpy := *c
	return &ccpy
}
