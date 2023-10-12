// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package kraftcloud

// Options contain necessary information for connecting to a KraftCloud service
// endpoint.
type Options struct {
	user         string
	token        string
	defaultMetro string
	http         HTTPClient
}

// Option is an option function used during initialization of a client.
type Option func(*Options) error

// NewDefaultOptions is a constructor method for instantiation a new set of
// default options for underlying requests to the KraftCloud API.
func NewDefaultOptions(opts ...Option) *Options {
	options := Options{}

	for _, opt := range opts {
		opt(&options)
	}

	if options.defaultMetro == "" {
		options.defaultMetro = DefaultMetro
	}

	if options.http == nil {
		options.http = NewHTTPClient()
	}

	return &options
}

// WithUser sets the username of the client connecting to KraftCloud.
func WithUser(user string) Option {
	return func(client *Options) error {
		client.user = user
		return nil
	}
}

// WithToken sets the access token of the client connecting to KraftCloud.
func WithToken(token string) Option {
	return func(client *Options) error {
		client.token = token
		return nil
	}
}

// WithHTTPClient sets the HTTP client that's used to customize the connection
// to KraftCloud's API.
func WithHTTPClient(httpClient HTTPClient) Option {
	return func(client *Options) error {
		client.http = httpClient
		return nil
	}
}

// WithDefaultMetro sets a KraftCloud metro, e.g. `fra0` which is based in
// Frankfurt.
func WithDefaultMetro(metro string) Option {
	return func(client *Options) error {
		client.defaultMetro = metro
		return nil
	}
}
