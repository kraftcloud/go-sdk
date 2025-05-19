// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2024, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package autoscale

// PolicyType is the type of the autoscale policy.
type PolicyType string

const (
	// PolicyTypeStep is the step autoscale policy type.
	PolicyTypeStep PolicyType = "step"

	// PolicyTypeOnDemand is the on-demand autoscale policy type.
	PolicyTypeOnDemand PolicyType = "on_demand"
)

const (
	// Endpoint is the public path for the autoscale service.
	Endpoint = "/services/autoscale"

	// AutoscaleEndpoint is the public path for the autoscale service.
	AutoscaleEndpoint = "/autoscale"

	// AutoscalePolicyEndpoint is the public path for the autoscale policy service.
	AutoscalePolicyEndpoint = "/policies"
)

// PolicyMetric is the autoscale policy metric.
type PolicyMetric string

const (
	// PolicyMetricCPU is the CPU autoscale policy metric.
	PolicyMetricCPU = "cpu"

	// PolicyMetricInflightRequests is the inflight requests autoscale policy
	// metric.
	PolicyMetricInflightRequests = "inflight_reqs"
)

// AdjustmentType is the autoscale adjustment type.
type AdjustmentType string

const (
	// AdjustmentTypePercent is the percent autoscale adjustment type.
	AdjustmentTypePercent AdjustmentType = "percent"
	// AdjustmentTypeAbsolute is the absolute autoscale adjustment type.
	AdjustmentTypeAbsolute AdjustmentType = "absolute"
	// AdjustmentTypeChange is the change autoscale adjustment type.
	AdjustmentTypeChange AdjustmentType = "change"
)
