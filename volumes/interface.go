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

	// GetByUUIDs returns the current state and the configuration of volumes.
	//
	// See: https://docs.kraft.cloud/api/v1/volumes/#getting-the-status-of-a-volume
	GetByUUIDs(ctx context.Context, uuids ...string) (*kcclient.ServiceResponse[GetResponseItem], error)

	// GetByNames returns the current state and the configuration of volumes.
	//
	// See: https://docs.kraft.cloud/api/v1/volumes/#getting-the-status-of-a-volume
	GetByNames(ctx context.Context, names ...string) (*kcclient.ServiceResponse[GetResponseItem], error)

	// AttachByUUID a volume to an instance so that the volume is mounted when the
	// instance starts using the volume and instance name.  The volume needs to be
	// in available state and the instance must in stopped state. Currently, each
	// instance can have only one volume attached at most.
	//
	// See: https://docs.kraft.cloud/api/v1/volumes/#attaching-a-volume-to-an-instance
	AttachByUUID(ctx context.Context, volUUID, instanceUUID, at string, readOnly bool) (*kcclient.ServiceResponse[AttachResponseItem], error)

	// AttachByName a volume to an instance so that the volume is mounted when the
	// instance starts using the volume and instance name.  The volume needs to be
	// in available state and the instance must in stopped state. Currently, each
	// instance can have only one volume attached at most.
	//
	// See: https://docs.kraft.cloud/api/v1/volumes/#attaching-a-volume-to-an-instance
	AttachByName(ctx context.Context, volName, instanceName, at string, readOnly bool) (*kcclient.ServiceResponse[AttachResponseItem], error)

	// DetachByUUID a volume from an instance based on the volume's UUID.  The
	// instance from which to detach must in stopped state. If the volume has been
	// created together with an instance, detaching the volume will make it
	// persistent (i.e., it survives the deletion of the instance).
	//
	// See: https://docs.kraft.cloud/api/v1/volumes/#detaching-a-volume-from-an-instance
	DetachByUUID(ctx context.Context, uuid string) (*kcclient.ServiceResponse[DetachResponseItem], error)

	// DetachByNam a volume from an instance based on the volume's name.  The
	// instance from which to detach must in stopped state. If the volume has been
	// created together with an instance, detaching the volume will make it
	// persistent (i.e., it survives the deletion of the instance).
	//
	// See: https://docs.kraft.cloud/api/v1/volumes/#detaching-a-volume-from-an-instance
	DetachByName(ctx context.Context, name string) (*kcclient.ServiceResponse[DetachResponseItem], error)

	// DeletebyUUID deletes the specified volumes based on their UUIDs.  Fails
	// if any of the specified volumes is still attached to an instance. After
	// this call the UUID of the volumes is no longer valid.
	//
	// See: https://docs.kraft.cloud/api/v1/volumes/#deleting-a-volume
	DeleteByUUIDs(ctx context.Context, uuids ...string) (*kcclient.ServiceResponse[DeleteResponseItem], error)

	// DeletebyNames deletes the specified volumes based on their names.  Fails
	// if any of the specified volumes is still attached to an instance. After
	// this call the UUID of the volumes is no longer valid.
	//
	// See: https://docs.kraft.cloud/api/v1/volumes/#deleting-a-volume
	DeleteByNames(ctx context.Context, names ...string) (*kcclient.ServiceResponse[DeleteResponseItem], error)

	// Lists all existing volumes. You can filter by persistence and volume
	// state. The returned volumes fulfill all provided filter criteria. No
	// particular value is assumed if a filter is not part of the request.
	//
	// See: https://docs.kraft.cloud/api/v1/volumes/#list-existing-volumes
	List(ctx context.Context) (*kcclient.ServiceResponse[ListResponseItem], error)
}
