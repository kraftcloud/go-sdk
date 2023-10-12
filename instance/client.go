// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package instance

import (
	"time"

	kraftcloud "sdk.kraft.cloud/v0"
)

// instancesClient is a basic wrapper around the v1 instance client of
// KraftCloud.
//
// See: https://docs.kraft.cloud/002-rest-api-v1-instances.html
type instancesClient struct {
	opts    *kraftcloud.Options
	request *kraftcloud.ServiceRequest
}

// NewInstancesClient instantiates a client which interface's with KraftCloud's
// instances API.
func NewInstancesClient(opts ...kraftcloud.Option) InstancesService {
	return &instancesClient{
		opts: kraftcloud.NewDefaultOptions(opts...),
	}
}

// NewInstancesClientFromOptions instantiates a new instances services client
// based on the provided pre-existing options.
func NewInstancesClientFromOptions(opts *kraftcloud.Options) InstancesService {
	return &instancesClient{
		opts: opts,
	}
}

// WithMetro sets the just-in-time metro to use when connecting to the
// KraftCloud API.
func (i *instancesClient) WithMetro(metro string) InstancesService {
	if i.request == nil {
		i.request = kraftcloud.NewServiceRequest()
	}
	i.request.SetMetro(metro)
	return i
}

// WithHTTPClient overwrites the base HTTP client.
func (i *instancesClient) WithHTTPClient(httpClient kraftcloud.HTTPClient) InstancesService {
	if i.request == nil {
		i.request = kraftcloud.NewServiceRequest()
	}
	i.request.SetHTTPClient(httpClient)
	return i
}

// WithTimeout sets the timeout when making a request.
func (i *instancesClient) WithTimeout(timeout time.Duration) InstancesService {
	if i.request == nil {
		i.request = kraftcloud.NewServiceRequest()
	}
	i.request.SetTimeout(timeout)
	return i
}
