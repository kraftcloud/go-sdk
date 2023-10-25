// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package image

import (
	"time"

	"sdk.kraft.cloud/client"
	"sdk.kraft.cloud/client/httpclient"
	"sdk.kraft.cloud/client/options"
)

// imagesClient wraps the v1 Image client of KraftCloud.
//
// See: https://docs.kraft.cloud/004-rest-api-v1-images.html
type imagesClient struct {
	// constructors must ensure that request is non-nil
	request *client.ServiceRequest
}

var _ ImagesService = (*imagesClient)(nil)

// NewImagesClientFromOptions instantiates a new image services client based on
// the provided pre-existing options.
func NewImagesClientFromOptions(opts *options.Options) ImagesService {
	return &imagesClient{
		request: client.NewServiceRequestFromDefaultOptions(opts),
	}
}

// WithMetro sets the just-in-time metro to use when connecting to the
// KraftCloud API.
func (c *imagesClient) WithMetro(m string) ImagesService {
	ccpy := c.clone()
	ccpy.request = c.request.WithMetro(m)
	return ccpy
}

// WithHTTPClient overwrites the base HTTP client.
func (c *imagesClient) WithHTTPClient(hc httpclient.HTTPClient) ImagesService {
	ccpy := c.clone()
	ccpy.request = c.request.WithHTTPClient(hc)
	return ccpy
}

// WithTimeout sets the timeout when making a request.
func (c *imagesClient) WithTimeout(to time.Duration) ImagesService {
	ccpy := c.clone()
	ccpy.request = c.request.WithTimeout(to)
	return ccpy
}

// clone returns a shallow copy of c.
func (c *imagesClient) clone() *imagesClient {
	ccpy := *c
	return &ccpy
}
