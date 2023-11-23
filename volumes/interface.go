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
	Create(ctx context.Context, req VolumeCreateRequest) (*Volume, error)

	// Returns the current state and the configuration of a volume.
	//
	// See: https://docs.kraft.cloud/006-rest-api-v1-volumes.html#state
	Get(ctx context.Context, uuid string) (*Volume, error)

	// Attaches a volume to an instance so that the volume is mounted when the
	// instance starts. The volume needs to be in available state and the instance
	// must in stopped state. Currently, each instance can have only one volume
	// attached at most.
	//
	// See: https://docs.kraft.cloud/006-rest-api-v1-volumes.html#attach
	Attach(ctx context.Context, uuid string, req VolumeAttachRequest) (*Volume, error)

	// Detaches a volume from an instance. The instance from which to detach must
	// in stopped state. If the volume has been created together with an instance,
	// detaching the volume will make it persistent (i.e., it survives the
	// deletion of the instance).
	//
	// See: https://docs.kraft.cloud/006-rest-api-v1-volumes.html#detach
	Detach(ctx context.Context, uuid string) (*Volume, error)

	// Deletes the specified volume. Fails if the volume is still attached to an
	// instance. After this call the UUID of the volumes is no longer valid.
	//
	// See: https://docs.kraft.cloud/006-rest-api-v1-volumes.html#delete
	Delete(ctx context.Context, uuid string) (*Volume, error)

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
