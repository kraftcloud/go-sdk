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
	// FeatureScaleToZero indicates that the instance can be scaled to zero.
	FeatureScaleToZero Feature = "scale-to-zero"

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
