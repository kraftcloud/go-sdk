// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2024, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package services

import (
	"sdk.kraft.cloud/instances"
	"sdk.kraft.cloud/services"
)

type AutoscalePolicyType string

const (
	AutoscalePolicyTypeStep = "step"
)

const (
	// Endpoint is the public path for the autoscale service.
	Endpoint = "/services/autoscale"

	// AutoscaleEndpoint is the public path for the autoscale service.
	AutoscaleEndpoint = "/autoscale"

	// AutoscalePolicyEndpoint is the public path for the autoscale policy service.
	AutoscalePolicyEndpoint = "/policies"
)

// A service helps describe the the load balancing and autoscale groups.
type Service struct {
	// Public-facing port.
	Port int `json:"port"`

	// Application port to which inbound traffic is redirected.
	DestinationPort int `json:"destination_port,omitempty"`

	// List of handlers.
	Handlers []services.Handler `json:"handlers,omitempty"`
}

// AutoscalePolicyMetric is the autoscale policy metric.
type AutoscalePolicyMetric string

const (
	// AutoscalePolicyMetricCPU is the CPU autoscale policy metric.
	AutoscalePolicyTypeCPU = "cpu"
)

// AutoscaleAdjustmentType is the autoscale adjustment type.
type AutoscaleAdjustmentType string

const (
	// AutoscaleAdjustmentTypePercent is the percent autoscale adjustment type.
	AutoscaleAdjustmentTypePercent = "percent"

	// AutoscaleAdjustmentTypeAbsolute is the absolute autoscale adjustment type.
	AutoscaleAdjustmentTypeAbsolute = "absolute"

	// AutoscaleAdjustmentTypeChange is the change autoscale adjustment type.
	AutoscaleAdjustmentTypeChange = "change"
)

type AutoscaleStepPolicyStep struct {
	// LowerBound is the lower bound of the autoscale policy step.
	LowerBound *int `json:"lower_bound,omitempty"`

	// UpperBound is the upper bound of the autoscale policy step.
	UpperBound *int `json:"upper_bound,omitempty"`

	// Adjustment is the adjustment of the autoscale policy step.
	Adjustment *int `json:"adjustment,omitempty"`
}

type AutoscaleStepPolicy struct {
	// Name of the policy.
	Name string `json:"name,omitempty"`

	// Type of the autoscale policy.
	Type AutoscalePolicyType `json:"type,omitempty"`

	// Metric is the metric of the autoscale policy.
	Metric AutoscalePolicyMetric `json:"metric,omitempty"`

	// AdjustmentType is the adjustment type of the autoscale policy.
	AdjustmentType AutoscaleAdjustmentType `json:"adjustment_type,omitempty"`

	// Steps are the steps of the autoscale policy.
	Steps []AutoscaleStepPolicyStep `json:"steps,omitempty"`
}

type AutoscaleConfiguration struct {
	// The status of the autoscale configuration.
	Status string `json:"status,omitempty"`

	// UUID is the UUID of the autoscale configuration.
	UUID string `json:"uuid,omitempty"`

	// Name is the name of the autoscale configuration.
	Name string `json:"name,omitempty"`

	// Enabled indicates if the autoscale configuration is enabled.
	Enabled bool `json:"enabled,omitempty"`

	// MinSize is the minimum number of instances of the autoscale configuration.
	MinSize uint `json:"min_size"` // 'omitempty' is not used here because 0 is a valid value.

	// MaxSize is the maximum number of instances of the autoscale configuration.
	MaxSize uint `json:"max_size,omitempty"`

	// WarmupTimeMs is the length of the warmup period in milliseconds.
	WarmupTimeMs uint `json:"warmup_time_ms,omitempty"`

	// CooldownTimeMs is the length of the cooldown period in milliseconds.
	CooldownTimeMs uint `json:"cooldown_time_ms,omitempty"`

	// Message contains the error message either on `partial_success` or `error`.
	Message string `json:"message,omitempty"`

	// Master is the master instance of the autoscale configuration.
	Master *instances.Instance `json:"master,omitempty"`

	// Policies is the list of autoscale policies of the autoscale configuration.
	Policies []map[string]interface{} `json:"policies,omitempty"`
}
