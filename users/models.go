// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package users

import kcclient "sdk.kraft.cloud/client"

// QuotasResponseItem is a data item from a response to a GET /users/quotas request.
// https://docs.kraft.cloud/api/v1/users/#list-quota-usage-and-limits
type QuotasResponseItem struct {
	UUID   string               `json:"uuid"`
	Used   QuotasResponseUsed   `json:"used"`
	Hard   QuotasResponseHard   `json:"hard"`
	Limits QuotasResponseLimits `json:"limits"`

	kcclient.APIResponseCommon
}

type QuotasResponseUsed struct {
	Instances     int `json:"instances"`
	LiveInstances int `json:"live_instances"`
	LiveVcpus     int `json:"live_vcpus"`
	LiveMemoryMb  int `json:"live_memory_mb"`
	ServiceGroups int `json:"service_groups"`
	Services      int `json:"services"`
	Volumes       int `json:"volumes"`
	TotalVolumeMb int `json:"total_volume_mb"`
}

type QuotasResponseHard struct {
	Instances     int `json:"instances"`
	LiveVcpus     int `json:"live_vcpus"`
	LiveMemoryMb  int `json:"live_memory_mb"`
	ServiceGroups int `json:"service_groups"`
	Services      int `json:"services"`
	Volumes       int `json:"volumes"`
	TotalVolumeMb int `json:"total_volume_mb"`
}

type QuotasResponseLimits struct {
	MinMemoryMb      int `json:"min_memory_mb"`
	MaxMemoryMb      int `json:"max_memory_mb"`
	MinVcpus         int `json:"min_vcpus"`
	MaxVcpus         int `json:"max_vcpus"`
	MinVolumeMb      int `json:"min_volume_mb"`
	MaxVolumeMb      int `json:"max_volume_mb"`
	MinAutoscaleSize int `json:"min_autoscale_size"`
	MaxAutoscaleSize int `json:"max_autoscale_size"`
}
