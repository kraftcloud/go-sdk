// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2022, Unikraft GmbH and The KraftKit Authors.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package instances

import (
	"fmt"
	"math"
	"strings"
	"syscall"
	"time"

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

type InstanceState string

const (
	// The instance is not running and does not count against live resource
	// quotas. Connections cannot be established.
	InstanceStateStopped InstanceState = "stopped"

	// The instance is booting up. This usually takes just a few milliseconds.
	InstanceStateStarting InstanceState = "starting"

	// Your application’s main entry point has been reached.
	InstanceStateRunning InstanceState = "running"

	// The instance is draining connections before shutting down. No new connections can be established.
	InstanceStateDraining InstanceState = "draining"

	// The instance is shutting down.
	InstanceStateStopping InstanceState = "stopping"

	// The instance has scale-to-zero enabled. The instance is not running, but will be automatically started when there are incoming requests.
	InstanceStateStandby InstanceState = "standby"
)

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
	StopCode          *uint                          `json:"stop_code,omitempty"`
	StopReason        *StopReason                    `json:"stop_reason,omitempty"`
	StartCount        uint                           `json:"start_count"`
	RestartCount      uint                           `json:"restart_count"`
	ExitCode          *uint                          `json:"exit_code,omitempty"`
	State             InstanceState                  `json:"state"`
	Image             string                         `json:"image"`
	MemoryMB          uint                           `json:"memory_mb"`
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

// Stop code of the kernel.  This value encodes multiple details about the stop
// irrespective of the application.
//
// MSB                                                     LSB
// ┌──────────────┬──────────┬──────────┬───────────┬────────┐
// │ 31 ────── 24 │ 23 ── 16 │    15    │ 14 ──── 8 │ 7 ── 0 │
// ├──────────────┼──────────┼──────────┼───────────┼────────┤
// │ reserved[^1] │ errno    │ shutdown │ initlevel │ reason │
// └──────────────┴──────────┴──────────┴───────────┴────────┘
//
// [^1]:      Reserved for future use.
// errno:     The application errno, using Linux's errno.h values.  (Optional,
//            can be 0.)
// shutdown:  Whether the shutdown originated from the inittable (0) or from the
//            termtable (1).
// initlevel: The initlevel at the time of the stop.
// reason:    The reason for the stop. See `StopCodeReason`.

type StopCodeMask uint

const (
	StopCodeMaskErrno     StopCodeMask = 0xFF0000
	StopCodeMaskShutdown  StopCodeMask = 0x008000
	StopCodeMaskInitLevel StopCodeMask = 0x007F00
	StopCodeMaskReason    StopCodeMask = 0x0000FF
)

// StopCodeErrno returns the application errno, using Linux's errno.h values.
func (item *GetResponseItem) StopCodeErrno() uint8 {
	if item.StopCode == nil {
		return 0
	}

	return uint8((*item.StopCode & uint(StopCodeMaskErrno)) >> 16)
}

// StopCodeShutdownTable returns whether the stop originated from the inittable
// (0) or from the termtable (1).
func (item *GetResponseItem) StopCodeShutdownTable() uint8 {
	if item.StopCode == nil {
		return 0
	}

	return uint8((*item.StopCode & uint(StopCodeMaskShutdown)) >> 15)
}

// StopCodeInitLevel returns the initlevel at the time of the stop.
func (item *GetResponseItem) StopCodeInitLevel() uint8 {
	if item.StopCode == nil {
		return 0
	}

	return uint8((*item.StopCode & uint(StopCodeMaskInitLevel)) >> 8)
}

type StopCodeReason uint8

const (
	// 0 - Success
	StopCodeReasonOK StopCodeReason = iota

	// 1 - Explicit crash (bugon/assert/crash/unhandled breakpoint)
	StopCodeReasonEXP

	// 2 - Arithmetic error
	StopCodeReasonMATH

	// 3 - Invalid CPU instruction or instruction error (e.g., operand alignment)
	StopCodeReasonINVLOP

	// 4 - Page fault - see errno (out of mem, EFAULT)
	StopCodeReasonPGFAULT

	// 5 - Segmentation fault
	StopCodeReasonSEGFAULT

	// 6 - Hardware error, NMI, alignment checks
	StopCodeReasonHWERR

	// 7 - Security violation, control protection (MTE, shadow stacks, PKU?)
	StopCodeReasonSECERR
)

// StopCodeReason returns all the reasons for the stop.
func StopCodeReasons() []string {
	return []string{
		"OK",
		"EXP",
		"MATH",
		"INVLOP",
		"PGFAULT",
		"SEGFAULT",
		"HWERR",
		"SECERR",
	}
}

// StopCodeReason provides the identity value for the reason for the stop.
func (item *GetResponseItem) StopCodeReason() StopCodeReason {
	if item.StopCode == nil {
		return StopCodeReasonOK
	}

	return StopCodeReason(*item.StopCode & uint(StopCodeMaskReason))
}

type StopReason uint

const (
	StopReasonKernel      StopReason = 1 << iota // 0b00001
	StopReasonApplication                        // 0b00010
	StopReasonPlatform                           // 0b00100
	StopReasonUser                               // 0b01000
	StopReasonForced                             // 0b10000
)

// DescribeStopOrigin provides a human-readable interpretation of the stop
// reason.
func (item *GetResponseItem) DescribeStopOrigin() string {
	if item.StopReason == nil || *item.StopReason == 0 {
		return "unknown"
	}

	var ret strings.Builder

	if *item.StopReason&StopReasonForced != 0 {
		ret.WriteString("force ")
	}

	ret.WriteString("initiated by ")

	switch true {
	case *item.StopReason&StopReasonPlatform == StopReasonPlatform && *item.StopReason&StopReasonUser != StopReasonUser:
		ret.WriteString("platform")
	case *item.StopReason&StopReasonUser == StopReasonUser:
		ret.WriteString("user")
	case *item.StopReason&StopReasonApplication == StopReasonApplication:
		ret.WriteString("app")
	case *item.StopReason&StopReasonKernel == StopReasonKernel:
		ret.WriteString("kernel")
	}

	return ret.String()
}

// StopOriginCode provides a human-readable interpretation of the stop reason in
// the form of a short-code.
func (item *GetResponseItem) StopOriginCode() string {
	if item.StopReason == nil || *item.StopReason == 0 {
		return "-----"
	}

	var ret strings.Builder

	if *item.StopReason&StopReasonForced == StopReasonForced {
		ret.WriteString("f")
	} else {
		ret.WriteString("-")
	}
	if *item.StopReason&StopReasonUser == StopReasonUser {
		ret.WriteString("u")
	} else {
		ret.WriteString("-")
	}
	if *item.StopReason&StopReasonPlatform == StopReasonPlatform {
		ret.WriteString("p")
	} else {
		ret.WriteString("-")
	}
	if *item.StopReason&StopReasonApplication == StopReasonApplication {
		ret.WriteString("a")
	} else {
		ret.WriteString("-")
	}
	if *item.StopReason&StopReasonKernel == StopReasonKernel {
		ret.WriteString("k")
	} else {
		ret.WriteString("-")
	}

	return ret.String()
}

// DescribeStopReason provides a human-readable description of the stop reason.
func (item *GetResponseItem) DescribeStopReason() string {
	if item.StopCode == nil || *item.StopCode == 0 {
		return ""
	}

	var ret strings.Builder

	switch true {
	case item.StopCodeShutdownTable() == 1 && (item.StopCodeInitLevel() == 0 || item.StopCodeInitLevel() == 1) && item.StopCodeReason() == StopCodeReasonOK:
		ret.WriteString("shutdown")
	case item.StopCodeReason() == StopCodeReasonEXP:
		ret.WriteString("assertion error")
	case item.StopCodeReason() == StopCodeReasonPGFAULT && item.StopCodeErrno() == 0xc:
		ret.WriteString("out of memory")
	case item.StopCodeReason() == StopCodeReasonPGFAULT && (item.StopCodeErrno() == 0xe || item.StopCodeErrno() == 0x1):
		ret.WriteString("illegal memory access")
	case item.StopCodeReason() == StopCodeReasonSEGFAULT:
		ret.WriteString("segmentation fault")
	case item.StopCodeReason() == StopCodeReasonPGFAULT:
		ret.WriteString("page fault")
	case item.StopCodeReason() == StopCodeReasonMATH:
		ret.WriteString("arithmetic error")
	case item.StopCodeReason() == StopCodeReasonINVLOP:
		ret.WriteString("instruction error")
	case item.StopCodeReason() == StopCodeReasonHWERR:
		ret.WriteString("hardware error")
	case item.StopCodeReason() == StopCodeReasonSECERR:
		ret.WriteString("security violation")
	default:
		ret.WriteString("unexpected error")
	}

	return ret.String()
}

// StopReasonCode returns a human-readable short-code representation of the stop
// reason.
func (item *GetResponseItem) StopReasonCode() string {
	if item.StopCode == nil || *item.StopCode == 0 {
		return ""
	}

	var ret strings.Builder

	if item.StopCodeShutdownTable() == 0 {
		ret.WriteString("i")
	} else {
		ret.WriteString("t")
	}

	ret.WriteString(fmt.Sprintf("%d", item.StopCodeInitLevel()))

	ret.WriteString(" ")

	ret.WriteString(StopCodeReasons()[item.StopCodeReason()])

	if item.StopCodeErrno() != 0 {
		ret.WriteString(" ")
		errno, ok := ErrnoNames()[syscall.Errno(item.StopCodeErrno())]
		if ok {
			ret.WriteString(errno)
		} else {
			ret.WriteString(fmt.Sprintf("%d", item.StopCodeErrno()))
		}
	}

	return ret.String()
}

// DescribeStatus returns a human-readable description of the instance's status.
func (item *GetResponseItem) DescribeStatus() string {
	switch item.State {
	case InstanceStateRunning:
		dur, err := time.ParseDuration(fmt.Sprintf("%dms", item.UptimeMs))
		if err != nil {
			return err.Error()
		}

		days := int64(dur.Hours() / 24)
		hours := int64(math.Mod(dur.Hours(), 24))
		minutes := int64(math.Mod(dur.Minutes(), 60))
		seconds := int64(math.Mod(dur.Seconds(), 60))

		chunks := []struct {
			singularName string
			amount       int64
		}{
			{"day", days},
			{"hr", hours},
			{"min", minutes},
			{"sec", seconds},
		}

		parts := []string{}

		for i, chunk := range chunks {
			if len(parts) > 0 && i+1 == len(chunks) { // Skip secs if greater than 1m
				continue
			}
			switch chunk.amount {
			case 0:
				continue
			case 1:
				parts = append(parts, fmt.Sprintf("%d%s", chunk.amount, chunk.singularName))
			default:
				parts = append(parts, fmt.Sprintf("%d%ss", chunk.amount, chunk.singularName))
			}
		}

		return fmt.Sprintf("since %s", strings.Join(parts, " "))
	case InstanceStateStopped:
		reason := item.DescribeStopReason()
		if reason == "shutdown" {
			return ""
		}

		return reason
	default:
		return string(item.State)
	}
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
