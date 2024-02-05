// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2024, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package services

import (
	"context"

	"sdk.kraft.cloud/client"
	"sdk.kraft.cloud/services"
)

type AutoscaleService interface {
	client.ServiceClient[AutoscaleService]

	// CreateConfigurationByUUID creates a new autoscale configuration with the
	// UUID of a service group.
	CreateConfigurationByUUID(ctx context.Context, uuid string, req AutoscaleConfiguration) (*services.ServiceGroup, error)

	// CreateConfigurationByName creates a new autoscale configuration with the
	// Name of a service group.
	CreateConfigurationByName(ctx context.Context, name string, req AutoscaleConfiguration) (*services.ServiceGroup, error)

	// GetConfigurationByName returns the current state and the configuration of
	// an autoscale configuration
	GetConfigurationByName(ctx context.Context, name string) (*AutoscaleConfiguration, error)

	// GetConfigurationByUUID returns the current state and the configuration of
	// an autoscale configuration
	GetConfigurationByUUID(ctx context.Context, uuid string) (*AutoscaleConfiguration, error)

	// DeleteConfigurationByUUID deletes an autoscale configuration given its
	// UUID.
	DeleteConfigurationByUUID(ctx context.Context, uuid string) error

	// DeleteConfigurationByName deletes an autoscale configuration given its
	// name.
	DeleteConfigurationByName(ctx context.Context, name string) error

	// CreatePolicy creates a new autoscale policy for an autoscale configuration.
	CreatePolicy(ctx context.Context, uuid string, typ AutoscalePolicyType, req interface{}) (*AutoscaleConfiguration, error)

	// DeletePolicyByName deletes an autoscale policy given its name.
	DeletePolicyByName(ctx context.Context, uuid, name string) (*AutoscaleConfiguration, error)

	// GetPolicyByName returns the current state and the configuration of an
	// autoscale policy
	GetPolicyByName(ctx context.Context, uuid, name string) (*map[string]interface{}, error)
}
