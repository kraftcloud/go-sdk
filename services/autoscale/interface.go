// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2024, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package autoscale

import (
	"context"

	kcclient "sdk.kraft.cloud/client"
)

type AutoscaleService interface {
	kcclient.ServiceClient[AutoscaleService]

	// CreateConfiguration creates a new autoscale configuration.
	CreateConfiguration(ctx context.Context, req CreateRequest) (*kcclient.ServiceResponse[CreateResponseItem], error)

	// GetConfigurationByName returns the current state and the configuration of
	// an autoscale configuration
	GetConfigurationByName(ctx context.Context, name string) (*kcclient.ServiceResponse[GetResponseItem], error)

	// GetConfigurationByUUID returns the current state and the configuration of
	// an autoscale configuration
	GetConfigurationByUUID(ctx context.Context, uuid string) (*kcclient.ServiceResponse[GetResponseItem], error)

	// DeleteConfigurationByUUID deletes an autoscale configuration given its
	// UUID.
	DeleteConfigurationByUUID(ctx context.Context, uuid string) (*kcclient.ServiceResponse[DeleteResponseItem], error)

	// DeleteConfigurationByName deletes an autoscale configuration given its
	// name.
	DeleteConfigurationByName(ctx context.Context, name string) (*kcclient.ServiceResponse[DeleteResponseItem], error)

	// AddPolicy adds a new autoscale policy to an autoscale configuration.
	AddPolicy(ctx context.Context, uuid string, req Policy) (*kcclient.ServiceResponse[AddPolicyResponseItem], error)

	// GetPolicyByName returns the current state and configuration of an
	// autoscale policy.
	GetPolicyByName(ctx context.Context, uuid, name string) (*kcclient.ServiceResponse[GetPolicyResponseItem], error)

	// DeletePolicyByName deletes an autoscale policy given its name.
	DeletePolicyByName(ctx context.Context, uuid, name string) (*kcclient.ServiceResponse[DeletePolicyResponseItem], error)
}
