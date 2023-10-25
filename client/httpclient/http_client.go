// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2022, Unikraft GmbH and The KraftKit Authors.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// Package httpclient provides an interface for enabling manipulating the
// underelying HTTP request performed by a client.
package httpclient

import "net/http"

// HTTPClient interface abstracts a generic HTTP request issuing client.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// NewHTTPClient creates a default Go HTTP client.
func NewHTTPClient() *http.Client {
	// We disable KeepAlive due to issues with the proxy in front of the API.
	return &http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: true,
		},
	}
}
