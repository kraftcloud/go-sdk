// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2022, Unikraft GmbH and The KraftKit Authors.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package volumes

import ukcclient "sdk.kraft.cloud/client"

// CreateResponseItem is a data item from a response to a POST /volumes request.
// https://docs.kraft.cloud/api/v1/volumes/#creating-volumes
type CreateResponseItem struct {
	Status string `json:"status"`
	UUID   string `json:"uuid"`
	Name   string `json:"name"`

	ukcclient.APIResponseCommon
}

// GetResponseItem is a data item from a response to a GET /volumes request.
// https://docs.kraft.cloud/api/v1/volumes/#getting-the-status-of-a-volume
type GetResponseItem struct {
	Status     string               `json:"status"`
	State      string               `json:"state"`
	UUID       string               `json:"uuid"`
	Name       string               `json:"name"`
	SizeMB     int                  `json:"size_mb"`
	AttachedTo []InstanceAttachment `json:"attached_to"`
	MountedBy  []InstanceMounting   `json:"mounted_by"`
	Persistent bool                 `json:"persistent"`
	CreatedAt  string               `json:"created_at"`

	ukcclient.APIResponseCommon
}

type InstanceAttachment struct {
	UUID string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
}

type InstanceMounting struct {
	UUID     string `json:"uuid"`
	Name     string `json:"name"`
	ReadOnly bool   `json:"readonly"`
}

// AttachResponseItem is a data item from a response to a PUT /volumes/attach request.
// https://docs.kraft.cloud/api/v1/volumes/#attaching-a-volume-to-an-instance
type AttachResponseItem struct {
	Status string `json:"status"`
	UUID   string `json:"uuid"`
	Name   string `json:"name"`

	ukcclient.APIResponseCommon
}

// DetachResponseItem is a data item from a response to a PUT /volumes/detach request.
// https://docs.kraft.cloud/api/v1/volumes/#detaching-a-volume-from-an-instance
type DetachResponseItem struct {
	Status string `json:"status"`
	UUID   string `json:"uuid"`
	Name   string `json:"name"`

	ukcclient.APIResponseCommon
}

// DeleteResponseItem is a data item from a response to a DELETE /volumes request.
// https://docs.kraft.cloud/api/v1/volumes/#deleting-a-volume
type DeleteResponseItem struct {
	Status string `json:"status"`
	UUID   string `json:"uuid"`
	Name   string `json:"name"`

	ukcclient.APIResponseCommon
}

// ListResponseItem is a data item from a response to a GET /volumes/list request.
// https://docs.kraft.cloud/api/v1/volumes/#list-existing-volumes
type ListResponseItem struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`

	ukcclient.APIResponseCommon
}
