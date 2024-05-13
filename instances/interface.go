// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package instances

import (
	"context"
	"time"

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
	Create(ctx context.Context, req CreateRequest) (*kcclient.ServiceResponse[CreateResponseItem], error)

	// Get returns the current state and the configuration of one or more instance(s).
	//
	// See: https://docs.kraft.cloud/api/v1/instances/#getting-the-status-of-an-instance
	Get(ctx context.Context, ids ...string) (*kcclient.ServiceResponse[GetResponseItem], error)

	// Delete deletes the specified instance(s).
	// After this call the UUIDs of the instances are no longer valid. If the
	// instances are currently running, they are force stopped.
	//
	// See: https://docs.kraft.cloud/api/v1/instances/#deleting-an-instance
	Delete(ctx context.Context, ids ...string) (*kcclient.ServiceResponse[DeleteResponseItem], error)

	// Lists all existing instances.
	//
	// See: https://docs.kraft.cloud/api/v1/instances/#list-existing-instances
	List(ctx context.Context) (*kcclient.ServiceResponse[GetResponseItem], error)

	// Start starts previously stopped instance(s).
	// Does nothing for instances that are already running.
	//
	// See: https://docs.kraft.cloud/api/v1/instances/#starting-an-instance
	Start(ctx context.Context, waitTimeoutMs int, ids ...string) (*kcclient.ServiceResponse[StartResponseItem], error)

	// Stop stops the specified instance(s), but does not destroy them.
	// All volatile state (e.g., RAM contents) is lost. Does nothing for
	// instances that are already stopped. Instances can be started again with
	// the start endpoint.
	//
	// See: https://docs.kraft.cloud/api/v1/instances/#stopping-an-instance
	Stop(ctx context.Context, drainTimeoutMs int, force bool, ids ...string) (*kcclient.ServiceResponse[StopResponseItem], error)

	// Log returns the console output of the specified instance.
	//
	// See: https://docs.kraft.cloud/api/v1/instances/#retrieve-the-console-output
	Log(ctx context.Context, id string, offset int, limit int) (*kcclient.ServiceResponse[LogResponseItem], error)

	// TailLogs is a utility method which returns a channel that streams the
	// console output of the specified instance.
	TailLogs(ctx context.Context, id string, follow bool, tail int, delay time.Duration) (chan string, chan error, error)

	// Wait waits for the specified instance(s) to reach the desired state.
	//
	// See: https://docs.kraft.cloud/api/v1/instances/#waiting-for-an-instance-to-reach-a-desired-state
	Wait(ctx context.Context, state State, timeoutMs int, ids ...string) (*kcclient.ServiceResponse[WaitResponseItem], error)
}
