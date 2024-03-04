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
	Create(ctx context.Context, name string, sizeMB int) (*CreateResponseItem, error)

	// GetByUUID returns the current state and the configuration of a volume.
	//
	// See: https://docs.kraft.cloud/api/v1/volumes/#getting-the-status-of-a-volume
	GetByUUID(ctx context.Context, uuid string) (*GetResponseItem, error)

	// GetByName returns the current state and the configuration of a volume.
	//
	// See: https://docs.kraft.cloud/api/v1/volumes/#getting-the-status-of-a-volume
	GetByName(ctx context.Context, name string) (*GetResponseItem, error)

	// AttachByUUID a volume to an instance so that the volume is mounted when the
	// instance starts using the volume and instance name.  The volume needs to be
	// in available state and the instance must in stopped state. Currently, each
	// instance can have only one volume attached at most.
	//
	// See: https://docs.kraft.cloud/api/v1/volumes/#attaching-a-volume-to-an-instance
	AttachByUUID(ctx context.Context, volUUID, instanceUUID, at string, readOnly bool) (*AttachResponseItem, error)

	// AttachByName a volume to an instance so that the volume is mounted when the
	// instance starts using the volume and instance name.  The volume needs to be
	// in available state and the instance must in stopped state. Currently, each
	// instance can have only one volume attached at most.
	//
	// See: https://docs.kraft.cloud/api/v1/volumes/#attaching-a-volume-to-an-instance
	AttachByName(ctx context.Context, volName, instanceName, at string, readOnly bool) (*AttachResponseItem, error)

	// DetachByUUID a volume from an instance based on the volume's UUID.  The
	// instance from which to detach must in stopped state. If the volume has been
	// created together with an instance, detaching the volume will make it
	// persistent (i.e., it survives the deletion of the instance).
	//
	// See: https://docs.kraft.cloud/api/v1/volumes/#detaching-a-volume-from-an-instance
	DetachByUUID(ctx context.Context, uuid string) (*DetachResponseItem, error)

	// DetachByNam a volume from an instance based on the volume's name.  The
	// instance from which to detach must in stopped state. If the volume has been
	// created together with an instance, detaching the volume will make it
	// persistent (i.e., it survives the deletion of the instance).
	//
	// See: https://docs.kraft.cloud/api/v1/volumes/#detaching-a-volume-from-an-instance
	DetachByName(ctx context.Context, name string) (*DetachResponseItem, error)

	// DeletebyUUID the specified volume based on its UUID.  Fails if the volume is
	// still attached to an instance. After this call the UUID of the volumes is
	// no longer valid.
	//
	// See: https://docs.kraft.cloud/api/v1/volumes/#deleting-a-volume
	DeleteByUUID(ctx context.Context, uuid string) (*DeleteResponseItem, error)

	// DeleteByName the specified volume based on its name.  Fails if the volume
	// is still attached to an instance. After this call the UUID of the volumes
	// is no longer valid.
	//
	// See: https://docs.kraft.cloud/api/v1/volumes/#deleting-a-volume
	DeleteByName(ctx context.Context, name string) (*DeleteResponseItem, error)

	// Lists all existing volumes. You can filter by persistence and volume
	// state. The returned volumes fulfill all provided filter criteria. No
	// particular value is assumed if a filter is not part of the request.
	//
	// See: https://docs.kraft.cloud/api/v1/volumes/#list-existing-volumes
	List(ctx context.Context) ([]ListResponseItem, error)
}
