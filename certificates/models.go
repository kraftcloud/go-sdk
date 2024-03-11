// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2022, Unikraft GmbH and The KraftKit Authors.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package certificates

import kcclient "sdk.kraft.cloud/client"

// GetResponseItem is a data item from a response to a GET /certificates request.
// https://docs.kraft.cloud/api/v1/certificates/#getting-the-status-of-a-certificate
type GetResponseItem struct {
	Status        string                    `json:"status"`
	UUID          string                    `json:"uuid"`
	Name          string                    `json:"name"`
	CreatedAt     string                    `json:"created_at"`
	CommonName    string                    `json:"common_name"`
	State         string                    `json:"state"`
	Validation    *GetResponseValidation    `json:"validation"`
	Subject       string                    `json:"subject"`
	Issuer        string                    `json:"issuer"`
	SerialNumber  string                    `json:"serial_number"`
	NotBefore     string                    `json:"not_before"`
	NotAfter      string                    `json:"not_after"`
	ServiceGroups []GetResponseServiceGroup `json:"service_groups"`

	kcclient.APIResponseCommon
}

type GetResponseValidation struct {
	Attempt int    `json:"attempt"`
	Next    string `json:"next"`
}

type GetResponseServiceGroup struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

// DeleteResponseItem is a data item from a response to a DELETE /certificates request.
// https://docs.kraft.cloud/api/v1/certificates/#deleting-a-certificate
type DeleteResponseItem struct {
	Status string `json:"status"`
	UUID   string `json:"uuid"`
	Name   string `json:"name"`

	kcclient.APIResponseCommon
}

// ListResponseItem is a data item from a response to a GET /certificates/list request.
// https://docs.kraft.cloud/api/v1/certificates/#list-existing-certificates
type ListResponseItem struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`

	kcclient.APIResponseCommon
}
