// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package image

import (
	"time"

	kraftcloud "sdk.kraft.cloud"
)

// imagesClient wraps the v1 Image client of KraftCloud.
//
// See: https://docs.kraft.cloud/004-rest-api-v1-images.html
type imagesClient struct {
	defOpts *kraftcloud.Options
	request *kraftcloud.ServiceRequest
}

var _ ImagesService = (*imagesClient)(nil)

// NewImagesClient instantiates a new image services client based on the
// provided options.
func NewImagesClient(opts ...kraftcloud.Option) ImagesService {
	return &imagesClient{
		defOpts: kraftcloud.NewDefaultOptions(opts...),
	}
}

// NewImagesClientFromOptions instantiates a new image services client based on
// the provided pre-existing options.
func NewImagesFromOptions(opts *kraftcloud.Options) ImagesService {
	return &imagesClient{
		defOpts: opts,
	}
}

// WithMetro sets the just-in-time metro to use when connecting to the
// KraftCloud API.
func (c *imagesClient) WithMetro(metro string) ImagesService {
	if c.request == nil {
		c.request = kraftcloud.NewServiceRequestFromDefaultOptions(c.defOpts)
	}
	c.request.SetMetro(metro)
	return c
}

// WithHTTPClient overwrites the base HTTP client.
func (c *imagesClient) WithHTTPClient(httpClient kraftcloud.HTTPClient) ImagesService {
	if c.request == nil {
		c.request = kraftcloud.NewServiceRequestFromDefaultOptions(c.defOpts)
	}
	c.request.SetHTTPClient(httpClient)
	return c
}

// WithTimeout sets the timeout when making a request.
func (c *imagesClient) WithTimeout(timeout time.Duration) ImagesService {
	if c.request == nil {
		c.request = kraftcloud.NewServiceRequestFromDefaultOptions(c.defOpts)
	}
	c.request.SetTimeout(timeout)
	return c
}
