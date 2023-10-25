// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// Package options provides the structure representing the instantiated client
// options.
package options

import "sdk.kraft.cloud/client/httpclient"

// Options contain necessary information for connecting to a KraftCloud service
// endpoint.
type Options struct {
	token        string
	defaultMetro string
	httpClient   httpclient.HTTPClient
}

func (opts *Options) SetToken(token string) {
	opts.token = token
}

func (opts *Options) Token() string {
	return opts.token
}

func (opts *Options) SetDefaultMetro(metro string) {
	opts.defaultMetro = metro
}

func (opts *Options) DefaultMetro() string {
	return opts.defaultMetro
}

func (opts *Options) SetHTTPClient(httpClient httpclient.HTTPClient) {
	opts.httpClient = httpClient
}

func (opts *Options) HTTPClient() httpclient.HTTPClient {
	return opts.httpClient
}
