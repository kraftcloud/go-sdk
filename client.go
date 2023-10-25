// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package client

import (
	"sdk.kraft.cloud/client/options"
	"sdk.kraft.cloud/image"
	"sdk.kraft.cloud/instance"
)

// Client provides access to the KraftCloud API.
type Client struct {
	instances instance.InstancesService
	images    image.ImagesService
}

// KraftCloud are the public endpoint categories for the KraftCloud API.
type KraftCloud interface {
	Instances() instance.InstancesService
	Images() image.ImagesService
}

// NewClient is the top-level KraftCloud Services client used to speak
// with the API.
func NewClient(copts ...Option) (KraftCloud, error) {
	return NewClientFromOptions(NewDefaultOptions(copts...))
}

// NewClientFromOptions is the top-level KraftCloud Services client used
// to speak with the API with pre-defined options.
func NewClientFromOptions(opts *options.Options) (KraftCloud, error) {
	// TODO(nderjung): Use dependency injection to dynamically instantiate all or
	// user-requested services. For now, instantiate all services.

	client := Client{
		instances: instance.NewInstancesClientFromOptions(opts),
		images:    image.NewImagesClientFromOptions(opts),
	}

	return &client, nil
}

// NewImagesClient instantiates a new image services client based on the
// provided options.
func NewImagesClient(opts ...Option) image.ImagesService {
	return image.NewImagesClientFromOptions(NewDefaultOptions(opts...))
}

// NewInstancesClient instantiates a client which interfaces with KraftCloud's
// instances API.
func NewInstancesClient(opts ...Option) instance.InstancesService {
	return instance.NewInstancesClientFromOptions(NewDefaultOptions(opts...))
}

// Instances returns InstancesService.
func (client *Client) Instances() instance.InstancesService {
	return client.instances
}

// Images returns ImagesService.
func (client *Client) Images() image.ImagesService {
	return client.images
}
