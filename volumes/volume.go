// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package volumes

const (
	// Endpoint is the public path for the volumes service.
	Endpoint = "/volumes"
)

// A volume can be in one of the following states.
type State string

const (
	// The volume is scheduled to be created.
	StateUninitialized = State("uninitialized")

	// The volume is currently created and formatted.
	StateInitializing = State("initializing")

	// The volume is healthy and available to be attached to an instance.
	StateAvailable = State("available")

	// The volume is healthy and attached to an instance. It is possible to detach
	// detach it in this state.
	StateIdle = State("idle")

	// The volume is currently mounted in a running unikernel.
	StateMounted = State("mounted")

	// There are maintenance tasks running on the volume.
	StateBusy = State("busy")

	// The volume is in an error state that needs inspection by a KraftCloud
	// engineer.
	StateError = State("error")

	// The request was successful.
	StateSuccess = State("success")
)

// VolumeAttachedToInstance represents an instance that a volume is attached to.
type VolumeAttachedToInstance struct {
	// UUID of the instance.
	UUID string `json:"uuid,omitempty"`

	// Name of the instance.
	Name string `json:"name,omitempty"`
}

// Compared to an initrd (initial ramdisk), volumes serve different purposes. An
// initrd is a file system loaded into memory during the boot process of the
// unikernel. Any new files created as well as any modifications made to files
// in the initrd are lost when the instance is stopped. In contrast, a volume is
// a persistent storage device that keeps data across restarts. In addition, it
// can be initialized with a set of files with one instance and then be detached
// and attached to a another one.
type Volume struct {
	// Current state of the volume (see states) or error if the request failed
	Status State `json:"status,omitempty"`

	// UUID of the volume.
	UUID string `json:"uuid,omitempty"`

	// Size of the volume in megabytes.
	SizeMB int `json:"size_mb,omitempty"`

	// List of instances that this volume is attached to.
	AttachedTo []VolumeAttachedToInstance `json:"attached_to,omitempty"`

	// Message contains the error message either on `partial_success` or `error`.
	Message string `json:"message,omitempty"`

	// Indicates if the volume will stay alive when the last instance is deleted
	// that this volume is attached to.
	Persistent bool `json:"persistent,omitempty"`

	// Date and time of creation in ISO8601.
	CreatedAt string `json:"created_at,omitempty"`
}
