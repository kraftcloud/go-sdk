// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package instances

import (
	"context"

	kcclient "sdk.kraft.cloud/client"
)

type InstancesService interface {
	kcclient.ServiceClient[InstancesService]

	// Creates one or more new instances of the specified Unikraft images. You can
	// describe the properties of the new instances such as their startup
	// arguments and amount of memory. Note that, the instance properties can only
	// be defined during creation. They cannot be changed later.
	//
	// See: https://docs.kraft.cloud/api/v1/instances/#creating-a-new-instance
	Create(ctx context.Context, req CreateRequest) (*CreateResponseItem, error)

	// GetByUUIDs returns the current state and the configuration of one or
	// more instance(s) based on the provided UUID(s).
	//
	// See: https://docs.kraft.cloud/api/v1/instances/#getting-the-status-of-an-instance
	GetByUUIDs(ctx context.Context, uuids ...string) ([]GetResponseItem, error)

	// GetByNames returns the current state and the configuration of one or
	// more instances based on the provided name(s).
	//
	// See: https://docs.kraft.cloud/api/v1/instances/#getting-the-status-of-an-instance
	GetByNames(ctx context.Context, names ...string) ([]GetResponseItem, error)

	// DeleteByUUIDs deletes the specified instances based on their UUID.
	// After this call the UUIDs of the instances are no longer valid. If the
	// instances are currently running, they are force stopped.
	//
	// See: https://docs.kraft.cloud/api/v1/instances/#deleting-an-instance
	DeleteByUUIDs(ctx context.Context, uuids ...string) ([]DeleteResponseItem, error)

	// DeleteByNames deletes the specified instances based on their names.
	// After this call the UUIDs of the instances are no longer valid. If the
	// instances are currently running, they are force stopped.
	//
	// See: https://docs.kraft.cloud/api/v1/instances/#deleting-an-instance
	DeleteByNames(ctx context.Context, names ...string) ([]DeleteResponseItem, error)

	// Lists all existing instances.
	//
	// See: https://docs.kraft.cloud/api/v1/instances/#list-existing-instances
	List(ctx context.Context) ([]ListResponseItem, error)

	// Starts a previously stopped instance(s) based on their UUID(s).
	// Does nothing for instances that are already running.
	//
	// See: https://docs.kraft.cloud/api/v1/instances/#starting-an-instance
	StartByUUIDs(ctx context.Context, waitTimeoutMs int, uuids ...string) ([]StartResponseItem, error)

	// Starts a previously stopped instance(s) based on their name(s).
	// Does nothing for instances that are already running.
	//
	// See: https://docs.kraft.cloud/api/v1/instances/#starting-an-instance
	StartByNames(ctx context.Context, waitTimeoutMs int, names ...string) ([]StartResponseItem, error)

	// StopByUUIDs stops the specified instance(s) based on their UUID(s), but
	// does not destroy them.  All volatile state (e.g., RAM contents) is lost.
	// Does nothing for instances that are already stopped. Instances can be
	// started again with the start endpoint.
	//
	// See: https://docs.kraft.cloud/api/v1/instances/#stopping-an-instance
	StopByUUIDs(ctx context.Context, drainTimeoutMs int, force bool, uuid ...string) ([]StopResponseItem, error)

	// StopByNames stops the specified instance(s) based on their name(s), but
	// does not destroy them.  All volatile state (e.g., RAM contents) is lost.
	// Does nothing for instances that are already stopped. Instances can be
	// started again with the start endpoint.
	//
	// See: https://docs.kraft.cloud/api/v1/instances/#stopping-an-instance
	StopByNames(ctx context.Context, drainTimeoutMs int, force bool, names ...string) ([]StopResponseItem, error)

	// ConsoleByName returns the console output of the specified instance based
	// on its name.
	//
	// See: https://docs.kraft.cloud/api/v1/instances/#retrieve-the-console-output
	ConsoleByName(ctx context.Context, name string, maxLines int, latest bool) (*ConsoleResponseItem, error)

	// ConsoleByUUID returns the console output of the specified instance based
	// on its UUID.
	//
	// See: https://docs.kraft.cloud/api/v1/instances/#retrieve-the-console-output
	ConsoleByUUID(ctx context.Context, uuid string, maxLines int, latest bool) (*ConsoleResponseItem, error)

	// WaitByUUIDs waits for the specified instance(s) based on their UUID(s) to
	// reach the desired state.
	// See: https://docs.kraft.cloud/api/v1/instances/#waiting-for-an-instance-to-reach-a-desired-state
	WaitByUUIDs(ctx context.Context, state State, timeoutMs int, uuids ...string) ([]WaitResponseItem, error)

	// WaitByNames waits for the specified instance(s) based on their name(s) to
	// reach the desired state.
	// See: https://docs.kraft.cloud/api/v1/instances/#waiting-for-an-instance-to-reach-a-desired-state
	WaitByNames(ctx context.Context, state State, timeoutMs int, names ...string) ([]WaitResponseItem, error)
}
