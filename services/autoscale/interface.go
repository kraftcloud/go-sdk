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

	// GetConfigurations returns the current states and configurations of
	// autoscale configurations.
	GetConfigurations(ctx context.Context, ids ...string) (*kcclient.ServiceResponse[GetResponseItem], error)

	// DeleteConfigurations deletes autoscale configurations.
	DeleteConfigurations(ctx context.Context, ids ...string) (*kcclient.ServiceResponse[DeleteResponseItem], error)

	// AddPolicy adds a new autoscale policy to an autoscale configuration.
	AddPolicy(ctx context.Context, autoscaleUUID string, req Policy) (*kcclient.ServiceResponse[AddPolicyResponseItem], error)

	// GetPolicy returns the current state and configuration of an autoscale policy.
	GetPolicy(ctx context.Context, autoscaleUUID, name string) (*kcclient.ServiceResponse[GetPolicyResponseItem], error)

	// DeletePolicy deletes an autoscale policy.
	DeletePolicy(ctx context.Context, autoscaleUUID, name string) (*kcclient.ServiceResponse[DeletePolicyResponseItem], error)
}
