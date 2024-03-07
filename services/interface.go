// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package services

import (
	"context"

	kcclient "sdk.kraft.cloud/client"
)

type ServicesService interface {
	kcclient.ServiceClient[ServicesService]

	// Creates one or more service groups with the given configuration. Note that,
	// the service group properties like published ports can only be defined
	// during creation. They cannot be changed later.
	//
	// Each port in a service group can specify a list of handlers that determine
	// how traffic arriving at the port is handled. See Connection Handlers for a
	// complete overview.
	//
	// You can specify an array of service group descriptions to create multiple
	// groups with different properties with the same call.
	//
	// See: https://docs.kraft.cloud/api/v1/services/#creating-new-service-groups
	Create(ctx context.Context, req CreateRequest) (*CreateResponseItem, error)

	// GetByUUID returns the current state and the configuration of a service group.
	//
	// See: https://docs.kraft.cloud/api/v1/services/#getting-the-state-of-a-service-group
	GetByUUID(ctx context.Context, uuid string) (*GetResponseItem, error)

	// GetByName returns the current state and the configuration of a service group.
	//
	// See: https://docs.kraft.cloud/api/v1/services/#getting-the-state-of-a-service-group
	GetByName(ctx context.Context, name string) (*GetResponseItem, error)

	// DeleteByUUID the specified service group based on its UUID.  Fails if there
	// are still instances attached to group. After this call the UUID of the
	// group is no longer valid.
	//
	// This operation cannot be undone.
	//
	// See: https://docs.kraft.cloud/api/v1/services/#deleting-a-service-group
	DeleteByUUID(ctx context.Context, uuid string) (*DeleteResponseItem, error)

	// DeleteByName the specified service group based on its name.  Fails if there
	// are still instances attached to group. After this call the UUID of the
	// group is no longer valid.
	//
	// This operation cannot be undone.
	//
	// See: https://docs.kraft.cloud/api/v1/services/#deleting-a-service-group
	DeleteByName(ctx context.Context, name string) (*DeleteResponseItem, error)

	// Lists all existing service groups. You can filter by persistence and DNS
	// name. The latter can be used to lookup the UUID of the service group that
	// owns a certain DNS name. The returned groups fulfill all provided filter
	// criteria. No particular value is assumed if a filter is not part of the
	// request.
	//
	// The array of groups in the response can be directly fed into the other
	// endpoints, for example, to delete (empty) groups.
	//
	// See: https://docs.kraft.cloud/api/v1/services/#list-existing-service-groups
	List(ctx context.Context) ([]ListResponseItem, error)
}
