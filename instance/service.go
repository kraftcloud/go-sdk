// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package instance

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
	// See: https://docs.kraft.cloud/002-rest-api-v1-instances.html#create
	Create(ctx context.Context, req CreateInstanceRequest) (*Instance, error)

	// Status returns the current status and the configuration of an instance.
	//
	// See: https://docs.kraft.cloud/002-rest-api-v1-instances.html#status
	Status(ctx context.Context, uuid string) (*Instance, error)

	// Lists all existing instances.
	//
	// See: https://docs.kraft.cloud/002-rest-api-v1-instances.html#list
	List(ctx context.Context) ([]Instance, error)

	// Stops the specified instance, but does not destroy it. All volatile state
	// (e.g., RAM contents) is lost. Does nothing for an instance that is already
	// stopped. The instance can be started again with the start endpoint.
	//
	// See: https://docs.kraft.cloud/002-rest-api-v1-instances.html#stop
	Stop(ctx context.Context, uuid string, drainTimeoutMS int64) (*Instance, error)

	// Starts a previously stopped instance. Does nothing for an instance that is
	// already running.
	//
	// See: https://docs.kraft.cloud/002-rest-api-v1-instances.html#start
	Start(ctx context.Context, uuid string, waitTimeoutMS int) (*Instance, error)

	// Deletes the specified instance. After this call the UUID of the instance is
	// no longer valid. If the instance is currently running it is force stopped.
	//
	// See: https://docs.kraft.cloud/002-rest-api-v1-instances.html#delete
	Delete(ctx context.Context, uuid string) error

	// Logs returns the console output of the specified instance.
	//
	// See: https://docs.kraft.cloud/002-rest-api-v1-instances.html#console
	Logs(ctx context.Context, uuid string, maxLines int, latest bool) (string, error)
}
