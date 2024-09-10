// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2022, Unikraft GmbH and The KraftKit Authors.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package images

import ukcclient "sdk.kraft.cloud/client"

// GetResponseItem is a data item from a response to a /images/list request.
// https://docs.kraft.cloud/api/v1/images/#list-existing-images
type GetResponseItem struct {
	Digest      string            `json:"digest"`
	Tags        []string          `json:"tags"`
	Initrd      bool              `json:"initrd"`
	SizeInBytes int64             `json:"size_in_bytes"`
	Args        string            `json:"args"`
	KernelArgs  string            `json:"kernel_args"`
	Labels      map[string]string `json:"labels"`

	ukcclient.APIResponseCommon
}

// QuotasResponseItem is a data item from a parsed response to a registry request at
// /projects/{project_name_or_id}/summary
type QuotasResponseItem struct {
	Used int64 `json:"quota"`
	Hard int64 `json:"hard"`
}
