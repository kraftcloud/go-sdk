// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2022, Unikraft GmbH and The KraftKit Authors.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.
package instance

import (
	"errors"
	"fmt"
)

const (
	// Endpoint is the public path for the instances service.
	Endpoint = "/instances"
)

// InstanceResponse holds instance description, as returned by the API.
type InstanceResponse struct {
	Status string `json:"status"`
	Data   struct {
		Instances []Instance `json:"instances"`
	} `json:"data"`
}

// NetworkInterface holds interface data returned by the Instance API.
type NetworkInterface struct {
	UUID      string `json:"uuid"`
	PrivateIP string `json:"private_ip"`
	MAC       string `json:"mac"`
}

// Instance holds the description of the KraftCloud compute instance, as
// understood by the API server.
//
// See: https://docs.kraft.cloud/002-rest-api-v1-instances.html#response_2
type Instance struct {
	// UUID of the instance.
	UUID string `json:"uuid" pretty:"UUID"`

	// Publicly accessible DNS name of the instance.
	DNS string `json:"dns" pretty:"DNS"`

	// Private IPv4 of the instance in CIDR notation for communication between
	// instances of the same user. This is equivalent to the IPv4 address of the
	// first network interface.
	PrivateIP string `json:"private_ip" pretty:"PrivateIP"`

	// Current state of the instance or error if the request failed.
	Status string `json:"status" pretty:"Status"`

	// Date and time of creation in ISO8601.
	CreatedAt string `json:"created_at" pretty:"Created At"`

	// Digest of the image that the instance uses.  Note that the image tag (e.g.,
	// latest) is translated by KraftCloud to the image digest that was assigned
	// the tag at the time of instance creation. The image is pinned to this
	// particular version.
	Image string `json:"image" pretty:"Image"`

	// Amount of memory assigned to the instance in megabytes.
	MemoryMB int `json:"memory_mb" pretty:"Memory (MB)"`

	// Application arguments.
	Args []string `json:"args" pretty:"Args"`

	// Key/value pairs to be set as environment variables at boot time.
	Env map[string]string `json:"env" pretty:"Env"`

	// UUID of the service group that the instance is part of.
	ServiceGroup string `json:"service_group" pretty:"Service Group"`

	// List of network interfaces attached to the instance.
	NetworkInterfaces []NetworkInterface `json:"network_interfaces" pretty:"Network Interfaces"`

	// Time it took to start the instance including booting Unikraft in
	// microseconds.
	BootTimeUS int64 `json:"boot_time_us" pretty:"Boot Time (ms)"`

	// When an instance has a specific issue an accompanying message is included
	// to help diagnose the state of the instance.
	Message string `json:"message"`

	// An error response code dictating the specific error type.
	Error int64 `json:"error"`
}

func (i *Instance) GetFieldByPrettyTag(tag string) string {
	// TODO(jake-ciolek): Use reflection?
	switch tag {
	case "UUID":
		return i.UUID
	case "DNS":
		return i.DNS
	case "PrivateIP":
		return i.PrivateIP
	case "Status":
		return i.Status
	case "Created At":
		return i.CreatedAt
	case "Image":
		return i.Image
	case "Memory (MB)":
		return fmt.Sprintf("%d", i.MemoryMB)
	case "Args":
		return fmt.Sprintf("%v", i.Args)
	case "Env":
		return fmt.Sprintf("%v", i.Env)
	case "Service Group":
		return i.ServiceGroup
	case "Network Interfaces":
		return fmt.Sprintf("%v", i.NetworkInterfaces)
	case "Boot Time (ms)":
		return fmt.Sprintf("%d", i.BootTimeUS)
	default:
		return ""
	}
}

func firstInstanceOrErr(response *InstanceResponse) (*Instance, error) {
	if response == nil {
		return nil, errors.New("response is nil")
	}
	if response.Data.Instances == nil {
		return nil, errors.New("instances data is nil")
	}
	if len(response.Data.Instances) == 0 {
		return nil, errors.New("no instances data returned from the server")
	}
	if response.Data.Instances[0].Status == "error" {
		return nil, errors.New(response.Data.Instances[0].Message)
	}
	return &response.Data.Instances[0], nil
}
