// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2022, Unikraft GmbH and The KraftKit Authors.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package instances

import (
	"sdk.kraft.cloud/services"
)

const (
	// Endpoint is the public path for the instances service.
	Endpoint = "/instances"
)

// InstanceFeature is a special feature of an instance.
type InstanceFeature string

const (
	// FeatureScaleToZero indicates that the instance can be scaled to zero.
	FeatureScaleToZero InstanceFeature = "scale-to-zero"
)

// NetworkInterface holds interface data returned by the Instance API.
type NetworkInterface struct {
	// UUID of the network interface.
	UUID string `json:"uuid,omitempty"`

	// Name of the network interface.
	Name string `json:"name,omitempty"`

	// Private IPv4 of network interface in CIDR notation.
	PrivateIP string `json:"private_ip,omitempty"`

	// MAC address of the network interface.
	MAC string `json:"mac,omitempty"`
}

type InstanceVolume struct {
	// UUID of the volume
	UUID string `json:"uuid,omitempty"`

	// Name of the volume
	Name string `json:"name,omitempty"`

	// Path of the mountpoint
	At string `json:"at,omitempty"`

	// Whether the volume is mounted read-only
	ReadOnly bool `json:"readonly,omitempty"`
}

// Instance holds the description of the KraftCloud compute instance, as
// understood by the API server.
//
// See: https://docs.kraft.cloud/002-rest-api-v1-instances.html#response_2
type Instance struct {
	// UUID of the instance.
	UUID string `json:"uuid,omitempty"`

	// Name of the instance.
	Name string `json:"name,omitempty"`

	// Publicly accessible FQDN name of the instance.
	FQDN string `json:"fqdn,omitempty"`

	// Private IPv4 of the instance in CIDR notation for communication between
	// instances of the same user. This is equivalent to the IPv4 address of the
	// first network interface.
	PrivateIP string `json:"private_ip,omitempty"`

	// Private fully qualified domain name of the instance for communication
	// between instances of the same user.
	PrivateFQDN string `json:"private_fqdn,omitempty"`

	// Current state of the instance or error if the request failed.
	State string `json:"state,omitempty"`

	// Date and time of creation in ISO8601.
	CreatedAt string `json:"created_at,omitempty"`

	// Digest of the image that the instance uses.  Note that the image tag (e.g.,
	// latest) is translated by KraftCloud to the image digest that was assigned
	// the tag at the time of instance creation. The image is pinned to this
	// particular version.
	Image string `json:"image,omitempty"`

	// Amount of memory assigned to the instance in megabytes.
	MemoryMB int `json:"memory_mb,omitempty"`

	// Application arguments.
	Args []string `json:"args,omitempty"`

	// Key/value pairs to be set as environment variables at boot time.
	Env map[string]string `json:"env,omitempty"`

	// The service group that the instance is part of.
	ServiceGroup services.ServiceGroup `json:"service_group,omitempty"`

	// Description of volumes.
	Volumes []InstanceVolume `json:"volumes,omitempty"`

	// Special features of the instance.
	Features []InstanceFeature `json:"features,omitempty"`

	// List of network interfaces attached to the instance.
	NetworkInterfaces []NetworkInterface `json:"network_interfaces,omitempty"`

	// Time it took to start the instance including booting Unikraft in
	// microseconds.
	BootTimeUS int64 `json:"boot_time_us,omitempty"`

	// When an instance has a specific issue an accompanying message is included
	// to help diagnose the state of the instance.
	Message string `json:"message,omitempty"`

	// An error response code dictating the specific error type.
	Error int64 `json:"error,omitempty"`

	// Base 64 encoded console output.
	Output string `json:"output,omitempty"`
}
