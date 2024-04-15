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
	Name          *string                    `json:"name,omitempty"`
	Image         string                     `json:"image"`
	Args          []string                   `json:"args,omitempty"`
	Env           map[string]string          `json:"env,omitempty"`
	MemoryMB      *int                       `json:"memory_mb,omitempty"`
	ServiceGroup  *CreateRequestServiceGroup `json:"service_group,omitempty"`
	Volumes       []CreateRequestVolume      `json:"volumes,omitempty"`
	Autostart     *bool                      `json:"autostart,omitempty"`
	Replicas      *int                       `json:"replicas,omitempty"`
	RestartPolicy *RestartPolicy             `json:"restart_policy,omitempty"`
	WaitTimeoutMs *int                       `json:"wait_timeout_ms,omitempty"`
	Features      []Feature                  `json:"features,omitempty"`
}

type CreateRequestServiceGroup struct {
	UUID     *string                         `json:"uuid,omitempty"`
	Name     *string                         `json:"name,omitempty"`
	Services []services.CreateRequestService `json:"services,omitempty"`
	Domains  []services.CreateRequestDomain  `json:"domains,omitempty"`
}

type CreateRequestVolume struct {
	UUID     *string `json:"uuid,omitempty"`
	Name     *string `json:"name,omitempty"`
	SizeMB   *int    `json:"size_mb,omitempty"`
	At       *string `json:"at,omitempty"`
	ReadOnly *bool   `json:"readonly,omitempty"`
}

// CreateResponseItem is a data item from a response to a POST /instances request.
// https://docs.kraft.cloud/api/v1/instances/#creating-a-new-instance
type CreateResponseItem struct {
	Status       string                         `json:"status"`
	State        string                         `json:"state"`
	UUID         string                         `json:"uuid"`
	Name         string                         `json:"name"`
	PrivateFQDN  string                         `json:"private_fqdn"`
	PrivateIP    string                         `json:"private_ip"`
	ServiceGroup *GetCreateResponseServiceGroup `json:"service_group"` // only if service_group was set in the request
	BootTimeUs   *int                           `json:"boot_time_us"`  // only if wait_timeout_ms was set in the request

	kcclient.APIResponseCommon
}

// GetResponseItem is a data item from a response to a GET /instances request.
// https://docs.kraft.cloud/api/v1/instances/#getting-the-status-of-an-instance
type GetResponseItem struct {
	Status            string                         `json:"status"`
	UUID              string                         `json:"uuid"`
	Name              string                         `json:"name"`
	CreatedAt         string                         `json:"created_at"`
	StartedAt         string                         `json:"started_at"`
	StoppedAt         string                         `json:"stopped_at"`
	UptimeMs          int                            `json:"uptime_ms"`
	Restart           *GetResponseRestart            `json:"restart,omitempty"`
	RestartPolicy     RestartPolicy                  `json:"restart_policy"`
	StopCode          *int                           `json:"stop_code,omitempty"`
	StopReason        *int                           `json:"stop_reason,omitempty"`
	StartCount        int                            `json:"start_count"`
	RestartCount      int                            `json:"restart_count"`
	ExitCode          *int                           `json:"exit_code,omitempty"`
	State             string                         `json:"state"`
	Image             string                         `json:"image"`
	MemoryMB          int                            `json:"memory_mb"`
	Args              []string                       `json:"args"`
	Env               map[string]string              `json:"env"`
	PrivateFQDN       string                         `json:"private_fqdn"`
	PrivateIP         string                         `json:"private_ip"`
	ServiceGroup      *GetCreateResponseServiceGroup `json:"service_group"`
	Volumes           []GetResponseVolume            `json:"volumes"`
	NetworkInterfaces []GetResponseNetworkInterface  `json:"network_interfaces"`
	BootTimeUs        int                            `json:"boot_time_us"` // always returned, even if never started

	kcclient.APIResponseCommon
}

type GetResponseRestart struct {
	Attempt int    `json:"attempt"`
	NextAt  string `json:"next_at"`
}

type GetCreateResponseServiceGroup struct {
	UUID    string                             `json:"uuid"`
	Name    string                             `json:"name"`
	Domains []services.GetCreateResponseDomain `json:"domains"`
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

// LogResponseItem is a data item from a response to a GET /instances/log request.
// https://docs.kraft.cloud/api/v1/instances/#retrieve-the-console-output
type LogResponseItem struct {
	Status    string               `json:"status"`
	UUID      string               `json:"uuid"`
	Name      string               `json:"name"`
	Output    string               `json:"output"`
	Range     LogResponseRange     `json:"range"`
	Available LogResponseAvailable `json:"available"`

	kcclient.APIResponseCommon
}

type LogResponseRange struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

type LogResponseAvailable struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

// WaitResponseItem is a data item from a response to a GET /instances/wait request.
// https://docs.kraft.cloud/api/v1/instances/#waiting-for-an-instance-to-reach-a-desired-state
type WaitResponseItem struct {
	Status string `json:"status"`
	UUID   string `json:"uuid"`
	Name   string `json:"name"`
	State  string `json:"state"`

	kcclient.APIResponseCommon
}
