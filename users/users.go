// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package users

const (
	// Endpoint is the public path for the users service.
	Endpoint = "/users"
)

// Quotas associated with an account.
type Quotas struct {
	// UUID of your user
	UUID string `json:"uuid,omitempty"`

	// The name of your user
	Name string `json:"name,omitempty"`

	// Used quota.
	Used struct {
		// Number of instances
		Instances int `json:"instances,omitempty"`

		// Number of instances that are not in the stopped state
		LiveInstances int `json:"live_instances,omitempty"`

		// Amount of memory assigned to instances that are not in the stopped state
		// in megabytes.
		LiveMemoryMb int `json:"live_memory_mb,omitempty"`

		// Number of service groups.
		ServiceGroups int `json:"service_groups,omitempty"`

		// Number of published network services over all existing service groups.
		Services int `json:"services,omitempty"`

		// Number of volumes.
		Volumes int `json:"volumes,omitempty"`

		// Total size of all volumes in megabytes.
		TotalVolumeMb int `json:"total_volume_mb,omitempty"`
	} `json:"used,omitempty"`

	// Same as used but contains the configured quota limits.
	Hard struct {
		// Number of instances
		Instances int `json:"instances,omitempty"`

		// Number of instances that are not in the stopped state
		LiveInstances int `json:"live_instances,omitempty"`

		// Amount of memory assigned to instances that are not in the stopped state
		// in megabytes.
		LiveMemoryMb int `json:"live_memory_mb,omitempty"`

		// Number of service groups.
		ServiceGroups int `json:"service_groups,omitempty"`

		// Number of published network services over all existing service groups.
		Services int `json:"services,omitempty"`

		// Number of volumes.
		Volumes int `json:"volumes,omitempty"`

		// Total size of all volumes in megabytes.
		TotalVolumeMb int `json:"total_volume_mb,omitempty"`
	} `json:"hard,omitempty"`

	// Additional limits.
	Limits struct {
		// Minimum amount of memory assigned to live instances in megabytes
		MinMemoryMb int `json:"min_memory_mb,omitempty"`

		// Maximum amount of memory assigned to live instances in megabytes
		MaxMemoryMb int `json:"max_memory_mb,omitempty"`

		// Minimum size of a volume in megabytes
		MinVolumeMb int `json:"min_volume_mb,omitempty"`

		// Maximum size of a volume in megabytes
		MaxVolumeMb int `json:"max_volume_mb,omitempty"`
	} `json:"limits,omitempty"`

	// Message contains the error message either on `partial_success` or `error`.
	Message string `json:"message,omitempty"`
}
