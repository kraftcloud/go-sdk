// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2022, Unikraft GmbH and The KraftKit Authors.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package instances

import "fmt"

// Endpoint is the public path for the instances service.
const Endpoint = "/instances"

// Feature is a special feature of an instance.
type Feature string

const (
	// FeatureDeleteOnStop indicates that the instance should be deleted when stopped.
	FeatureDeleteOnStop Feature = "delete-on-stop"
)

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

// ScaleToZeroPolicy is the scale to zero policy of an instance.
type ScaleToZeroPolicy string

const (

	// ScaleToZeroPolicyOn indicates that the instance has scale to zero enabled.
	ScaleToZeroPolicyOn ScaleToZeroPolicy = "on"

	// ScaleToZeroPolicyOff indicates that the instance has scale to zero disabled.
	ScaleToZeroPolicyOff ScaleToZeroPolicy = "off"

	// ScaleToZeroPolicyIdle indicates that the instance will scale down even with established but idle TCP connections.
	ScaleToZeroPolicyIdle ScaleToZeroPolicy = "idle"
)

var _ fmt.Stringer = (*ScaleToZeroPolicy)(nil)

// String implements fmt.Stringer
func (policy ScaleToZeroPolicy) String() string {
	return string(policy)
}

// ScaleToZeroPolicies returns all scale to zero policies.
func ScaleToZeroPolicies() []ScaleToZeroPolicy {
	return []ScaleToZeroPolicy{
		ScaleToZeroPolicyOn,
		ScaleToZeroPolicyOff,
		ScaleToZeroPolicyIdle,
	}
}

// RestartPolicy is the restart policy of an instance.
type RestartPolicy string

const (
	// RestartPolicyNever indicates that the instance should never be restarted.
	RestartPolicyNever RestartPolicy = "never"

	// RestartPolicyAlways indicates that the instance should always be restarted.
	RestartPolicyAlways RestartPolicy = "always"

	// RestartPolicyOnFailure indicates that the instance should be restarted on failure.
	RestartPolicyOnFailure RestartPolicy = "on-failure"
)

var _ fmt.Stringer = (*RestartPolicy)(nil)

// String implements fmt.Stringer
func (policy RestartPolicy) String() string {
	return string(policy)
}

// RestartPolicies returns all restart policies.
func RestartPolicies() []RestartPolicy {
	return []RestartPolicy{
		RestartPolicyNever,
		RestartPolicyAlways,
		RestartPolicyOnFailure,
	}
}

// LogDefaultPageSize is the default page size for log requests.
const LogDefaultPageSize = 4096

// LogMaxPageSize is the maximum page size for log requests.
const LogMaxPageSize = LogDefaultPageSize*4 - 1
