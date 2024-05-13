// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package volumes

import (
	"context"

	kcclient "sdk.kraft.cloud/client"
)

type VolumesService interface {
	kcclient.ServiceClient[VolumesService]

	// Creates one or more volumes with the given configuration. The volumes are
	// automatically initialized with an empty file system. After initialization
	// completed the volumes are in the available state and can be attached to an
	// instance with the attach endpoint. Note that, the size of a volume cannot
	// be changed after creation.
	//
	// See: https://docs.kraft.cloud/api/v1/volumes/#creating-volumes
	Create(ctx context.Context, name string, sizeMB int) (*kcclient.ServiceResponse[CreateResponseItem], error)

	// Get returns the current state and the configuration of volumes.
	//
	// See: https://docs.kraft.cloud/api/v1/volumes/#getting-the-status-of-a-volume
	Get(ctx context.Context, ids ...string) (*kcclient.ServiceResponse[GetResponseItem], error)

	// Attach attaches a volume to an instance so that the volume is mounted
	// when the instance starts using the volume and instance name.  The volume
	// needs to be in available state and the instance must in stopped state.
	// Currently, each instance can have only one volume attached at most.
	//
	// See: https://docs.kraft.cloud/api/v1/volumes/#attaching-a-volume-to-an-instance
	Attach(ctx context.Context, volID, instanceUUID, at string, readOnly bool) (*kcclient.ServiceResponse[AttachResponseItem], error)

	// Detach detaches a volume from an instance.
	// The instance from which to detach must in stopped state. If the volume
	// has been created together with an instance, detaching the volume will
	// make it persistent (i.e., it survives the deletion of the instance).
	//
	// See: https://docs.kraft.cloud/api/v1/volumes/#detaching-a-volume-from-an-instance
	Detach(ctx context.Context, id string) (*kcclient.ServiceResponse[DetachResponseItem], error)

	// Delete deletes the specified volume(s).
	// Fails if any of the specified volumes is still attached to an instance.
	// After this call the UUID of the volumes is no longer valid.
	//
	// See: https://docs.kraft.cloud/api/v1/volumes/#deleting-a-volume
	Delete(ctx context.Context, ids ...string) (*kcclient.ServiceResponse[DeleteResponseItem], error)

	// Lists all existing volumes. You can filter by persistence and volume
	// state. The returned volumes fulfill all provided filter criteria. No
	// particular value is assumed if a filter is not part of the request.
	//
	// See: https://docs.kraft.cloud/api/v1/volumes/#list-existing-volumes
	List(ctx context.Context) (*kcclient.ServiceResponse[GetResponseItem], error)
}
