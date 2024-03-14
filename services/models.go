// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2022, Unikraft GmbH and The KraftKit Authors.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package services

import kcclient "sdk.kraft.cloud/client"

// CreateRequest is the payload for a POST /services request.
// https://docs.kraft.cloud/api/v1/services/#creating-a-new-service-group
type CreateRequest struct {
	Name     *string                `json:"name"`
	DNSName  *string                `json:"dns_name"`
	Services []CreateRequestService `json:"services"`
}

type CreateRequestService struct {
	Port            int       `json:"port"`
	DestinationPort *int      `json:"destination_port"`
	Handlers        []Handler `json:"handlers"`
}

// CreateResponseItem is a data item from a response to a POST /services request.
// https://docs.kraft.cloud/api/v1/services/#creating-a-new-service-group
type CreateResponseItem struct {
	Status string `json:"status"`
	UUID   string `json:"uuid"`
	Name   string `json:"name"`
	FQDN   string `json:"fqdn"`

	kcclient.APIResponseCommon
}

// GetResponseItem is a data item from a response to a GET /services request.
// https://docs.kraft.cloud/api/v1/services/#getting-the-status-of-a-service-group
type GetResponseItem struct {
	Status     string                `json:"status"`
	UUID       string                `json:"uuid"`
	Name       string                `json:"name"`
	CreatedAt  string                `json:"created_at"`
	FQDN       string                `json:"fqdn"`
	Instances  []GetResponseInstance `json:"instances"`
	Services   []GetResponseService  `json:"services"`
	Persistent bool                  `json:"persistent"`
	Autoscale  bool                  `json:"autoscale"`

	kcclient.APIResponseCommon
}

type GetResponseService struct {
	Port            int       `json:"port"`
	DestinationPort int       `json:"destination_port"`
	Handlers        []Handler `json:"handlers"`
}

type GetResponseInstance struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

// DeleteResponseItem is a data item from a response to a DELETE /services request.
// https://docs.kraft.cloud/api/v1/services/#deleting-a-service-group
type DeleteResponseItem struct {
	Status string `json:"status"`
	UUID   string `json:"uuid"`
	Name   string `json:"name"`

	kcclient.APIResponseCommon
}

// ListResponseItem is a data item from a response to a GET /services/list request.
// https://docs.kraft.cloud/api/v1/services/#list-existing-service-groups
type ListResponseItem struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`

	kcclient.APIResponseCommon
}
