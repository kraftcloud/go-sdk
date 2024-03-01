// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package instances

import (
	"context"

	"sdk.kraft.cloud/client"
)

type InstancesService interface {
	client.ServiceClient[InstancesService]

	// Creates one or more new instances of the specified Unikraft images. You can
	// describe the properties of the new instances such as their startup
	// arguments and amount of memory. Note that, the instance properties can only
	// be defined during creation. They cannot be changed later.
	//
	// See: https://docs.kraft.cloud/api/v1/instances/#create
	Create(ctx context.Context, req CreateInstanceRequest) (*Instance, error)

	// GetByUUID returns the current state and the configuration of an instance
	// based on the provided UUID.
	//
	// See: https://docs.kraft.cloud/api/v1/instances/#status
	GetByUUID(ctx context.Context, uuid string) (*Instance, error)

	// GetByName returns the current state and the configuration of an instance
	// based on the provided name.
	//
	// See: https://docs.kraft.cloud/api/v1/instances/#status
	GetByName(ctx context.Context, name string) (*Instance, error)

	// Lists all existing instances.
	//
	// See: https://docs.kraft.cloud/api/v1/instances/#list
	List(ctx context.Context) ([]Instance, error)

	// StopByUUID the specified instance based on its UUID, but does not destroy
	// it.  All volatile state (e.g., RAM contents) is lost. Does nothing for an
	// instance that is already stopped. The instance can be started again with
	// the start endpoint.
	//
	// See: https://docs.kraft.cloud/api/v1/instances/#stop
	StopByUUID(ctx context.Context, uuid string, drainTimeoutMS int64) (*Instance, error)

	// Stops the specified instance based on its name, but does not destroy it.
	// All volatile state (e.g., RAM contents) is lost. Does nothing for an
	// instance that is already stopped. The instance can be started again with
	// the start endpoint.
	//
	// See: https://docs.kraft.cloud/api/v1/instances/#stop
	StopByName(ctx context.Context, name string, drainTimeoutMS int64) (*Instance, error)

	// Starts a previously stopped instance based on its UUID. Does nothing for an
	// instance that is already running.
	//
	// See: https://docs.kraft.cloud/api/v1/instances/#start
	StartByUUID(ctx context.Context, uuid string, waitTimeoutMS int) (*Instance, error)

	// Starts a previously stopped instance based on its name. Does nothing for an
	// instance that is already running.
	//
	// See: https://docs.kraft.cloud/api/v1/instances/#start
	StartByName(ctx context.Context, name string, waitTimeoutMS int) (*Instance, error)

	// DeleteByUUID the specified instance based on its UUID. After this call the
	// UUID of the instance is no longer valid. If the instance is currently
	// running it is force stopped.
	//
	// See: https://docs.kraft.cloud/api/v1/instances/#delete
	DeleteByUUID(ctx context.Context, uuid string) error

	// DeleteByName deletes the specified instance based on its name. After this
	// call the UUID of the instance is no longer valid. If the instance is
	// currently running it is force stopped.
	//
	// See: https://docs.kraft.cloud/api/v1/instances/#delete
	DeleteByName(ctx context.Context, name string) error

	// LogsByName returns the console output of the specified instance based on
	// its name.
	//
	// See: https://docs.kraft.cloud/api/v1/instances/#console
	LogsByName(ctx context.Context, name string, maxLines int, latest bool) (string, error)

	// LogsByUUID returns the console output of the specified instance based on its
	// UUID.
	//
	// See: https://docs.kraft.cloud/api/v1/instances/#console
	LogsByUUID(ctx context.Context, uuid string, maxLines int, latest bool) (string, error)
}
