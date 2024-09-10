// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2022, Unikraft GmbH and The KraftKit Authors.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package services

import ukcclient "sdk.kraft.cloud/client"

// CreateRequest is the payload for a POST /services request.
// https://docs.kraft.cloud/api/v1/services/#creating-a-new-service-group
type CreateRequest struct {
	Name      *string                `json:"name,omitempty"`
	Domains   []CreateRequestDomain  `json:"domains,omitempty"`
	Services  []CreateRequestService `json:"services,omitempty"`
	SoftLimit *int                   `json:"soft_limit,omitempty"`
	HardLimit *int                   `json:"hard_limit,omitempty"`
}

type CreateRequestDomain struct {
	Name        string                          `json:"name"`
	Certificate *CreateRequestDomainCertificate `json:"certificate,omitempty"`
}

type CreateRequestDomainCertificate struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

type CreateRequestService struct {
	Port            int       `json:"port"`
	DestinationPort *int      `json:"destination_port,omitempty"`
	Handlers        []Handler `json:"handlers,omitempty"`
}

// PatchRequest is the payload for a PATCH /services request.
// https://docs.kraft.cloud/api/v1/services/#update-a-service
type PatchRequest struct {
	UUID  string  `json:"uuid,omitempty"`
	Name  *string `json:"name,omitempty"`
	Op    *string `json:"op,omitempty"`
	Value *any    `json:"value,omitempty"`
	ID    *any    `json:"id,omitempty"`
}

// CreateResponseItem is a data item from a response to a POST /services request.
// https://docs.kraft.cloud/api/v1/services/#creating-a-new-service-group
type CreateResponseItem struct {
	Status  string                    `json:"status"`
	UUID    string                    `json:"uuid"`
	Name    string                    `json:"name"`
	Domains []GetCreateResponseDomain `json:"domains"`

	ukcclient.APIResponseCommon
}

// GetResponseItem is a data item from a response to a GET /services request.
// https://docs.kraft.cloud/api/v1/services/#getting-the-status-of-a-service-group
type GetResponseItem struct {
	Status     string                    `json:"status"`
	UUID       string                    `json:"uuid"`
	Name       string                    `json:"name"`
	CreatedAt  string                    `json:"created_at"`
	Persistent bool                      `json:"persistent"`
	Autoscale  bool                      `json:"autoscale"`
	Services   []GetResponseService      `json:"services"`
	Domains    []GetCreateResponseDomain `json:"domains"`
	Instances  []GetResponseInstance     `json:"instances"`
	SoftLimit  int                       `json:"soft_limit"`
	HardLimit  int                       `json:"hard_limit"`

	ukcclient.APIResponseCommon
}

type GetCreateResponseDomain struct {
	FQDN        string                              `json:"fqdn"`
	Certificate *GetCreateResponseDomainCertificate `json:"certificate"`
}

type GetCreateResponseDomainCertificate struct {
	UUID  string `json:"uuid"`
	Name  string `json:"name"`
	State string `json:"state"`
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

	ukcclient.APIResponseCommon
}

// ListResponseItem is a data item from a response to a GET /services/list request.
// https://docs.kraft.cloud/api/v1/services/#list-existing-service-groups
type ListResponseItem struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`

	ukcclient.APIResponseCommon
}
