// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package services

import (
	"context"

	"sdk.kraft.cloud/client"
)

type ServicesService interface {
	client.ServiceClient[ServicesService]

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
	// See: https://docs.kraft.cloud/003-rest-api-v1-services.html#creating-new-service-groups
	Create(ctx context.Context, req ServiceCreateRequest) (*ServiceGroup, error)

	// Returns the current status and the configuration of a service group.
	//
	// See: https://docs.kraft.cloud/003-rest-api-v1-services.html#getting-the-status-of-a-service-group
	Status(ctx context.Context, uuid string) (*ServiceGroup, error)

	// Lists all existing service groups. You can filter by persistence and DNS
	// name. The latter can be used to lookup the UUID of the service group that
	// owns a certain DNS name. The returned groups fulfill all provided filter
	// criteria. No particular value is assumed if a filter is not part of the
	// request.
	//
	// The array of groups in the response can be directly fed into the other
	// endpoints, for example, to delete (empty) groups.
	//
	// See: https://docs.kraft.cloud/003-rest-api-v1-services.html#list-existing-service-groups
	List(ctx context.Context) ([]ServiceGroup, error)

	// Deletes the specified service group. Fails if there are still instances
	// attached to group. After this call the UUID of the group is no longer
	// valid.
	//
	// This operation cannot be undone.
	//
	// See: https://docs.kraft.cloud/003-rest-api-v1-services.html#deleting-a-service-group
	Delete(ctx context.Context, uuid string) (*ServiceGroup, error)
}