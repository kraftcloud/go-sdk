// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2022, Unikraft GmbH and The KraftKit Authors.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package instances

// Endpoint is the public path for the instances service.
const Endpoint = "/instances"

// Feature is a special feature of an instance.
type Feature string

// FeatureScaleToZero indicates that the instance can be scaled to zero.
const FeatureScaleToZero Feature = "scale-to-zero"

// State is the state of an instance.
type State string

const (
	// StateStopped indicates that the instance is stopped.
	StateStopped State = "stopped"

	// StateStarting indicates that the instance is starting.
	StateStarting State = "starting"

	// StateRunning indicates that the instance is running.
	StateRunning State = "running"

	// StateDraining indicates that the instance is draining.
	StateDraining State = "draining"

	// StateStopping indicates that the instance is stopping.
	StateStopping State = "stopping"

	// StateStandby indicates that the instance is in standby.
	StateStandby State = "standby"
)
