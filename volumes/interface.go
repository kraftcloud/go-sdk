// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package volumes

import (
	"context"

	"sdk.kraft.cloud/client"
)

type VolumesService interface {
	client.ServiceClient[VolumesService]

	// Creates one or more volumes with the given configuration. The volumes are
	// automatically initialized with an empty file system. After initialization
	// completed the volumes are in the available state and can be attached to an
	// instance with the attach endpoint. Note that, the size of a volume cannot
	// be changed after creation.
	//
	// See: https://docs.kraft.cloud/006-rest-api-v1-volumes.html#create
	Create(ctx context.Context, name string, sizeMB int) (*Volume, error)

	// GetByUUID returns the current state and the configuration of a volume.
	//
	// See: https://docs.kraft.cloud/006-rest-api-v1-volumes.html#state
	GetByUUID(ctx context.Context, uuid string) (*Volume, error)

	// GetByName returns the current state and the configuration of a volume.
	//
	// See: https://docs.kraft.cloud/006-rest-api-v1-volumes.html#state
	GetByName(ctx context.Context, name string) (*Volume, error)

	// AttachByUUID a volume to an instance so that the volume is mounted when the
	// instance starts using the volume and instance name.  The volume needs to be
	// in available state and the instance must in stopped state. Currently, each
	// instance can have only one volume attached at most.
	//
	// See: https://docs.kraft.cloud/006-rest-api-v1-volumes.html#attach
	AttachByUUID(ctx context.Context, volUUID, instanceUUID, at string, readOnly bool) (*Volume, error)

	// AttachByName a volume to an instance so that the volume is mounted when the
	// instance starts using the volume and instance name.  The volume needs to be
	// in available state and the instance must in stopped state. Currently, each
	// instance can have only one volume attached at most.
	//
	// See: https://docs.kraft.cloud/006-rest-api-v1-volumes.html#attach
	AttachByName(ctx context.Context, volName, instanceName, at string, readOnly bool) (*Volume, error)

	// DetachByUUID a volume from an instance based on the volume's UUID.  The
	// instance from which to detach must in stopped state. If the volume has been
	// created together with an instance, detaching the volume will make it
	// persistent (i.e., it survives the deletion of the instance).
	//
	// See: https://docs.kraft.cloud/006-rest-api-v1-volumes.html#detach
	DetachByUUID(ctx context.Context, uuid string) (*Volume, error)

	// DetachByNam a volume from an instance based on the volume's name.  The
	// instance from which to detach must in stopped state. If the volume has been
	// created together with an instance, detaching the volume will make it
	// persistent (i.e., it survives the deletion of the instance).
	//
	// See: https://docs.kraft.cloud/006-rest-api-v1-volumes.html#detach
	DetachByName(ctx context.Context, name string) (*Volume, error)

	// DeletebyUUID the specified volume based on its UUID.  Fails if the volume is
	// still attached to an instance. After this call the UUID of the volumes is
	// no longer valid.
	//
	// See: https://docs.kraft.cloud/006-rest-api-v1-volumes.html#delete
	DeleteByUUID(ctx context.Context, uuid string) error

	// DeleteByName the specified volume based on its name.  Fails if the volume
	// is still attached to an instance. After this call the UUID of the volumes
	// is no longer valid.
	//
	// See: https://docs.kraft.cloud/006-rest-api-v1-volumes.html#delete
	DeleteByName(ctx context.Context, name string) error

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
	List(ctx context.Context) ([]Volume, error)
}
