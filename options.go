// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package kraftcloud

import (
	"os"
	"time"

	"sdk.kraft.cloud/client"
	"sdk.kraft.cloud/client/httpclient"
	"sdk.kraft.cloud/client/options"
)

// Option is an option function used during initialization of a client.
type Option func(*options.Options)

// NewDefaultOptions is a constructor method for instantiation a new set of
// default options for underlying requests to the KraftCloud API.
func NewDefaultOptions(opts ...Option) *options.Options {
	options := options.Options{}

	for _, opt := range opts {
		opt(&options)
	}

	if options.Token() == "" {
		options.SetToken(os.Getenv("UNIKRAFTCLOUD_TOKEN"))
	}

	if options.Token() == "" {
		options.SetToken(os.Getenv("KRAFTCLOUD_TOKEN"))
	}

	if options.Token() == "" {
		options.SetToken(os.Getenv("KC_TOKEN"))
	}

	if options.Token() == "" {
		options.SetToken(os.Getenv("UKC_TOKEN"))
	}

	if options.DefaultMetro() == "" {
		options.SetDefaultMetro(client.DefaultMetro)
	}

	if options.DefaultTimeout() == 0 {
		options.SetDefaultTimeout(30 * time.Second)
	}

	if options.DefaultRetries() == 0 {
		options.SetDefaultRetries(5)
	}

	if options.HTTPClient() == nil {
		options.SetHTTPClient(httpclient.NewHTTPClient())
	}

	return &options
}

// WithToken sets the access token of the client connecting to KraftCloud.
func WithToken(token string) Option {
	return func(client *options.Options) {
		client.SetToken(token)
	}
}

// WithHTTPClient sets the HTTP client that's used to customize the connection
// to KraftCloud's API.
func WithHTTPClient(httpClient httpclient.HTTPClient) Option {
	return func(client *options.Options) {
		client.SetHTTPClient(httpClient)
	}
}

// WithDefaultMetro sets a KraftCloud metro, e.g. `fra0` which is based in
// Frankfurt.
func WithDefaultMetro(metro string) Option {
	return func(client *options.Options) {
		client.SetDefaultMetro(metro)
	}
}
