// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2024, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package autoscale

import (
	"context"

	ukcclient "sdk.kraft.cloud/client"
)

type AutoscaleService interface {
	ukcclient.ServiceClient[AutoscaleService]

	// CreateConfiguration creates a new autoscale configuration.
	CreateConfiguration(ctx context.Context, req CreateRequest) (*ukcclient.ServiceResponse[CreateResponseItem], error)

	// GetConfigurations returns the current states and configurations of
	// autoscale configurations.
	GetConfigurations(ctx context.Context, ids ...string) (*ukcclient.ServiceResponse[GetResponseItem], error)

	// DeleteConfigurations deletes autoscale configurations.
	DeleteConfigurations(ctx context.Context, ids ...string) (*ukcclient.ServiceResponse[DeleteResponseItem], error)

	// AddPolicy adds a new autoscale policy to an autoscale configuration.
	AddPolicy(ctx context.Context, autoscaleUUID string, req Policy) (*ukcclient.ServiceResponse[AddPolicyResponseItem], error)

	// GetPolicy returns the current state and configuration of an autoscale policy.
	GetPolicy(ctx context.Context, autoscaleUUID, name string) (*ukcclient.ServiceResponse[GetPolicyResponseItem], error)

	// DeletePolicy deletes an autoscale policy.
	DeletePolicy(ctx context.Context, autoscaleUUID, name string) (*ukcclient.ServiceResponse[DeletePolicyResponseItem], error)
}
