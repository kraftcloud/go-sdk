// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package kraftcloud

import (
	"sdk.kraft.cloud/client/options"
	"sdk.kraft.cloud/images"
	"sdk.kraft.cloud/instances"
	"sdk.kraft.cloud/services"
	scale "sdk.kraft.cloud/services/autoscale"
	"sdk.kraft.cloud/users"
	"sdk.kraft.cloud/volumes"
)

// Client provides access to the KraftCloud API.
type Client struct {
	autoscale scale.AutoscaleService
	instances instances.InstancesService
	images    images.ImagesService
	services  services.ServicesService
	volumes   volumes.VolumesService
	users     users.UsersService
}

// KraftCloud are the public endpoint categories for the KraftCloud API.
type KraftCloud interface {
	Autoscale() scale.AutoscaleService
	Instances() instances.InstancesService
	Images() images.ImagesService
	Services() services.ServicesService
	Users() users.UsersService
	Volumes() volumes.VolumesService
}

// NewClient is the top-level KraftCloud Services client used to speak
// with the API.
func NewClient(copts ...Option) KraftCloud {
	return NewClientFromOptions(NewDefaultOptions(copts...))
}

// NewClientFromOptions is the top-level KraftCloud Services client used
// to speak with the API with pre-defined options.
func NewClientFromOptions(opts *options.Options) KraftCloud {
	// TODO(nderjung): Use dependency injection to dynamically instantiate all or
	// user-requested services. For now, instantiate all services.

	client := Client{
		autoscale: scale.NewAutoscaleClientFromOptions(opts),
		instances: instances.NewInstancesClientFromOptions(opts),
		images:    images.NewImagesClientFromOptions(opts),
		services:  services.NewServicesClientFromOptions(opts),
		users:     users.NewUsersClientFromOptions(opts),
		volumes:   volumes.NewVolumesClientFromOptions(opts),
	}

	return &client
}

// NewAutoscaleCLient instantiates a client which interfaces with KraftCloud's
// autoscale API.
func NewAutoscaleClient(opts ...Option) scale.AutoscaleService {
	return scale.NewAutoscaleClientFromOptions(NewDefaultOptions(opts...))
}

// NewImagesClient instantiates a new image services client based on the
// provided options.
func NewImagesClient(opts ...Option) images.ImagesService {
	return images.NewImagesClientFromOptions(NewDefaultOptions(opts...))
}

// NewInstancesClient instantiates a client which interfaces with KraftCloud's
// instances API.
func NewInstancesClient(opts ...Option) instances.InstancesService {
	return instances.NewInstancesClientFromOptions(NewDefaultOptions(opts...))
}

// NewServicesClient instantiates a client which interfaces with KraftCloud's
// volumes API.
func NewServicesClient(opts ...Option) services.ServicesService {
	return services.NewServicesClientFromOptions(NewDefaultOptions(opts...))
}

// NewUsersClient instantiates a client which interfaces with KraftCloud's users
// API.
func NewUsersClient(opts ...Option) users.UsersService {
	return users.NewUsersClientFromOptions(NewDefaultOptions(opts...))
}

// NewVolumesClient instantiates a client which interfaces with KraftCloud's
// volumes API.
func NewVolumesClient(opts ...Option) volumes.VolumesService {
	return volumes.NewVolumesClientFromOptions(NewDefaultOptions(opts...))
}

// Autoscale returns AutoscaleService.
func (client *Client) Autoscale() scale.AutoscaleService {
	return client.autoscale
}

// Instances returns InstancesService.
func (client *Client) Instances() instances.InstancesService {
	return client.instances
}

// Images returns ImagesService.
func (client *Client) Images() images.ImagesService {
	return client.images
}

// Services returns ServicesService.
func (client *Client) Services() services.ServicesService {
	return client.services
}

// Users returns UsersService.
func (client *Client) Users() users.UsersService {
	return client.users
}

// Volumes returns VolumesService.
func (client *Client) Volumes() volumes.VolumesService {
	return client.volumes
}
