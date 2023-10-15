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
	opts    *kraftcloud.Options
	request *kraftcloud.ServiceRequest
}

// NewImagesClient instantiates a new image services client based on the
// provided options.
func NewImagesClient(opts ...kraftcloud.Option) ImagesService {
	return &imagesClient{
		opts: kraftcloud.NewDefaultOptions(opts...),
	}
}

// NewImagesClientFromOptions instantiates a new image services client based on
// the provided pre-existing options.
func NewImagesFromOptions(opts *kraftcloud.Options) ImagesService {
	return &imagesClient{
		opts: opts,
	}
}

// WithMetro sets the just-in-time metro to use when connecting to the
// KraftCloud API.
func (i *imagesClient) WithMetro(metro string) ImagesService {
	if i.request == nil {
		i.request = kraftcloud.NewServiceRequestFromDefaultOptions(i.opts)
	}
	i.request.SetMetro(metro)
	return i
}

// WithHTTPClient overwrites the base HTTP client.
func (i *imagesClient) WithHTTPClient(httpClient kraftcloud.HTTPClient) ImagesService {
	if i.request == nil {
		i.request = kraftcloud.NewServiceRequestFromDefaultOptions(i.opts)
	}
	i.request.SetHTTPClient(httpClient)
	return i
}

// WithTimeout sets the timeout when making a request.
func (i *imagesClient) WithTimeout(timeout time.Duration) ImagesService {
	if i.request == nil {
		i.request = kraftcloud.NewServiceRequestFromDefaultOptions(i.opts)
	}
	i.request.SetTimeout(timeout)
	return i
}
