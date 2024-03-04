// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2022, Unikraft GmbH and The KraftKit Authors.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package instances

import (
	kcclient "sdk.kraft.cloud/client"
	"sdk.kraft.cloud/services"
)

// CreateRequest is the payload for a POST /instances request.
// https://docs.kraft.cloud/api/v1/instances/#creating-a-new-instance
type CreateRequest struct {
	Name          *string                    `json:"name"`
	Image         string                     `json:"image"`
	Args          []string                   `json:"args"`
	Env           map[string]string          `json:"env"`
	MemoryMB      *int                       `json:"memory_mb"`
	ServiceGroup  *CreateRequestServiceGroup `json:"service_group"`
	Volumes       []CreateRequestVolume      `json:"volumes"`
	Autostart     *bool                      `json:"autostart"`
	Replicas      *int                       `json:"replicas"`
	WaitTimeoutMs *int                       `json:"wait_timeout_ms"`
	Features      []Feature                  `json:"features"`
}

type CreateRequestServiceGroup struct {
	UUID     *string                         `json:"uuid"`
	Name     *string                         `json:"name"`
	Services []services.CreateRequestService `json:"services"`
	DNSName  *string                         `json:"dns_name"`
}

type CreateRequestVolume struct {
	UUID     string `json:"uuid"`
	Name     string `json:"name"`
	SizeMB   int    `json:"size_mb"`
	At       string `json:"at"`
	ReadOnly *bool  `json:"readonly"`
}

// CreateResponseItem is a data item from a response to a POST /instances request.
// https://docs.kraft.cloud/api/v1/instances/#creating-a-new-instance
type CreateResponseItem struct {
	Status      string `json:"status"`
	State       string `json:"state"`
	UUID        string `json:"uuid"`
	Name        string `json:"name"`
	FQDN        string `json:"fqdn"`
	PrivateFQDN string `json:"private_fqdn"`
	PrivateIP   string `json:"private_ip"`
	BootTimeUs  *int   `json:"boot_time_us"` // only if wait_timeout_ms was set in the request

	kcclient.APIResponseCommon
}

// GetResponseItem is a data item from a response to a GET /instances request.
// https://docs.kraft.cloud/api/v1/instances/#getting-the-status-of-an-instance
type GetResponseItem struct {
	Status            string                        `json:"status"`
	UUID              string                        `json:"uuid"`
	Name              string                        `json:"name"`
	CreatedAt         string                        `json:"created_at"`
	State             string                        `json:"state"`
	Image             string                        `json:"image"`
	MemoryMB          int                           `json:"memory_mb"`
	Args              []string                      `json:"args"`
	Env               map[string]string             `json:"env"`
	FQDN              string                        `json:"fqdn"`
	PrivateFQDN       string                        `json:"private_fqdn"`
	PrivateIP         string                        `json:"private_ip"`
	ServiceGroup      *GetResponseServiceGroup      `json:"service_group"`
	Volumes           []GetResponseVolume           `json:"volumes"`
	NetworkInterfaces []GetResponseNetworkInterface `json:"network_interfaces"`
	BootTimeUs        int                           `json:"boot_time_us"` // always returned, even if never started

	kcclient.APIResponseCommon
}

type GetResponseServiceGroup struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

type GetResponseVolume struct {
	UUID     string `json:"uuid"`
	Name     string `json:"name"`
	At       string `json:"at"`
	ReadOnly bool   `json:"readonly"`
}

type GetResponseNetworkInterface struct {
	UUID      string `json:"uuid"`
	PrivateIP string `json:"private_ip"`
	MAC       string `json:"mac"`
}

// DeleteResponseItem is a data item from a response to a DELETE /instances request.
// https://docs.kraft.cloud/api/v1/instances/#deleting-an-instance
type DeleteResponseItem struct {
	Status        string `json:"status"`
	UUID          string `json:"uuid"`
	Name          string `json:"name"`
	PreviousState string `json:"previous_state"`

	kcclient.APIResponseCommon
}

// ListResponseItem is a data item from a response to a GET /instances/list request.
// https://docs.kraft.cloud/api/v1/instances/#list-existing-instances
type ListResponseItem struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`

	kcclient.APIResponseCommon
}

// StartResponseItem is a data item from a response to a POST /instances/start request.
// https://docs.kraft.cloud/api/v1/instances/#starting-an-instance
type StartResponseItem struct {
	Status        string `json:"status"`
	UUID          string `json:"uuid"`
	Name          string `json:"name"`
	State         string `json:"state"`
	PreviousState string `json:"previous_state"`

	kcclient.APIResponseCommon
}

// StopResponseItem is a data item from a response to a POST /instances/stop request.
// https://docs.kraft.cloud/api/v1/instances/#stopping-an-instance
type StopResponseItem struct {
	Status        string `json:"status"`
	UUID          string `json:"uuid"`
	Name          string `json:"name"`
	State         string `json:"state"`
	PreviousState string `json:"previous_state"`

	kcclient.APIResponseCommon
}

// ConsoleResponseItem is a data item from a response to a GET /instances/console request.
// https://docs.kraft.cloud/api/v1/instances/#retrieve-the-console-output
type ConsoleResponseItem struct {
	Status string `json:"status"`
	UUID   string `json:"uuid"`
	Name   string `json:"name"`
	Output string `json:"output"`

	kcclient.APIResponseCommon
}
