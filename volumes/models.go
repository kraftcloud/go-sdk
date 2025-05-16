// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2022, Unikraft GmbH and The KraftKit Authors.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package volumes

import kcclient "sdk.kraft.cloud/client"

// CreateRequest is a data structure for a request to a POST /volumes request.
// https://docs.kraft.cloud/api/v1/volumes/#creating-volumes-templates
type CreateRequestTemplate struct {
	Name *string `json:"name,omitempty"`
	UUID *string `json:"uuid,omitempty"`
}

// CreateRequest is a data structure for a request to a POST /volumes request.
// https://docs.kraft.cloud/api/v1/volumes/#creating-volumes
type CreateRequest struct {
	Name     *string                `json:"name"`
	Template *CreateRequestTemplate `json:"template,omitempty"`
	SizeMb   *int                   `json:"size_mb,omitempty"`
}

// CreateResponseItem is a data item from a response to a POST /volumes request.
// https://docs.kraft.cloud/api/v1/volumes/#creating-volumes
type CreateResponseItem struct {
	Status string `json:"status"`
	State  string `json:"state"`
	UUID   string `json:"uuid"`
	Name   string `json:"name"`

	kcclient.APIResponseCommon
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

	kcclient.APIResponseCommon
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

	kcclient.APIResponseCommon
}

// DetachResponseItem is a data item from a response to a PUT /volumes/detach request.
// https://docs.kraft.cloud/api/v1/volumes/#detaching-a-volume-from-an-instance
type DetachResponseItem struct {
	Status string `json:"status"`
	UUID   string `json:"uuid"`
	Name   string `json:"name"`

	kcclient.APIResponseCommon
}

// DeleteResponseItem is a data item from a response to a DELETE /volumes request.
// https://docs.kraft.cloud/api/v1/volumes/#deleting-a-volume
type DeleteResponseItem struct {
	Status string `json:"status"`
	UUID   string `json:"uuid"`
	Name   string `json:"name"`

	kcclient.APIResponseCommon
}

// ListResponseItem is a data item from a response to a GET /volumes/list request.
// https://docs.kraft.cloud/api/v1/volumes/#list-existing-volumes
type ListResponseItem struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`

	kcclient.APIResponseCommon
}

// TemplateCreateResponseItem is a data item from a response to a POST /volumes/templates request.
// https://docs.kraft.cloud/api/v1/volumes/templates#creating-templates
type TemplateCreateResponseItem struct {
	Status string `json:"status"`
	State  string `json:"state"`
	UUID   string `json:"uuid"`
	Name   string `json:"name"`

	kcclient.APIResponseCommon
}

// TemplateGetResponseItem is a data item from a response to a GET /volumes/templates request.
// https://docs.kraft.cloud/api/v1/volumes/templates#getting-the-status-of-a-template
type TemplateGetResponseItem struct {
	State       string               `json:"state"`
	UUID        string               `json:"uuid"`
	Name        string               `json:"name"`
	SizeMB      int                  `json:"size_mb"`
	FreeMB      int                  `json:"free_mb"`
	AttachedTo  []InstanceAttachment `json:"attached_to"`
	Persistent  bool                 `json:"persistent"`
	QuotaPolicy string               `json:"quota_policy"`
	CreatedAt   string               `json:"created_at"`

	kcclient.APIResponseCommon
}

// TemplateDeleteResponseItem is a data item from a response to a DELETE /volumes/templates request.
// https://docs.kraft.cloud/api/v1/volumes/templates#deleting-a-template
type TemplateDeleteResponseItem struct {
	Status string `json:"status"`
	UUID   string `json:"uuid"`
	Name   string `json:"name"`

	kcclient.APIResponseCommon
}
