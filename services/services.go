// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package client

import (
	kraftcloud "sdk.kraft.cloud"
	"sdk.kraft.cloud/image"
	"sdk.kraft.cloud/instance"
)

// ServicesClient provides access to the KraftCloud API.
type ServicesClient struct {
	instances instance.InstancesService
	images    image.ImagesService
}

// KraftCloudServices are the public endpoint categories for the KraftCloud API.
type KraftCloudServices interface {
	Instances() instance.InstancesService
	Images() image.ImagesService
}

// NewServicesClient is the top-level KraftCloud Services client used to speak
// with the API.
func NewServicesClient(copts ...kraftcloud.Option) (KraftCloudServices, error) {
	return NewServicesClientFromOptions(kraftcloud.NewDefaultOptions(copts...))
}

// NewServicesClientFromOptions is the top-level KraftCloud Services client used
// to speak with the API with pre-defined options.
func NewServicesClientFromOptions(opts *kraftcloud.Options) (KraftCloudServices, error) {
	// TODO(nderjung): Use dependency injection to dynamically instantiate all or
	// user-requested services. For now, instantiate all services.

	client := ServicesClient{
		instances: instance.NewInstancesClientFromOptions(opts),
		images:    image.NewImagesFromOptions(opts),
	}

	return &client, nil
}

// Instances returns InstancesService.
func (client *ServicesClient) Instances() instance.InstancesService {
	return client.instances
}

// Images returns ImagesService.
func (client *ServicesClient) Images() image.ImagesService {
	return client.images
}
