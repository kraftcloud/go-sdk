// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package users

import (
	"time"

	ukcclient "sdk.kraft.cloud/client"
	"sdk.kraft.cloud/client/httpclient"
	"sdk.kraft.cloud/client/options"
)

// client wraps the v1 Users client of UnikraftCloud.
//
// See: https://docs.kraft.cloud/api/v1/users/
type client struct {
	// constructors must ensure that request is non-nil
	request *ukcclient.ServiceRequest
}

var _ UsersService = (*client)(nil)

// NewUsersClientFromOptions instantiates a new users services client based on
// the provided pre-existing options.
func NewUsersClientFromOptions(opts *options.Options) UsersService {
	return &client{
		request: ukcclient.NewServiceRequestFromDefaultOptions(opts),
	}
}

// WithMetro sets the just-in-time metro to use when connecting to the
// UnikraftCloud API.
func (c *client) WithMetro(m string) UsersService {
	ccpy := c.clone()
	ccpy.request = c.request.WithMetro(m)
	return ccpy
}

// WithHTTPClient overwrites the base HTTP client.
func (c *client) WithHTTPClient(hc httpclient.HTTPClient) UsersService {
	ccpy := c.clone()
	ccpy.request = c.request.WithHTTPClient(hc)
	return ccpy
}

// WithTimeout sets the timeout when making a request.
func (c *client) WithTimeout(to time.Duration) UsersService {
	ccpy := c.clone()
	ccpy.request = c.request.WithTimeout(to)
	return ccpy
}

// clone returns a shallow copy of c.
func (c *client) clone() *client {
	ccpy := *c
	return &ccpy
}
