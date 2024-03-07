// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package volumes

// Endpoint is the public path for the volumes service.
const Endpoint = "/volumes"

// A volume can be in one of the following states.
type State string

const (
	// The volume is scheduled to be created.
	StateUninitialized State = "uninitialized"

	// The volume is currently created and formatted.
	StateInitializing State = "initializing"

	// The volume is healthy and available to be attached to an instance.
	StateAvailable State = "available"

	// The volume is healthy and attached to an instance. It is possible to detach
	// detach it in this state.
	StateIdle State = "idle"

	// The volume is currently mounted in a running unikernel.
	StateMounted State = "mounted"

	// There are maintenance tasks running on the volume.
	StateBusy State = "busy"
)
