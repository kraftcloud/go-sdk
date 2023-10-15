// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package kraftcloud

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

// ServiceRequest is the utility structure for performing individual requests to
// a service location at KraftCloud.
type ServiceRequest struct {
	opts       *Options
	metro      string
	httpClient HTTPClient
	timeout    time.Duration
}

// NewServiceRequest is a constructor method which prepares an individual
// request.
func NewServiceRequest(copts ...Option) *ServiceRequest {
	opts := Options{}

	for _, opt := range copts {
		opt(&opts)
	}

	return &ServiceRequest{}
}

// NewServiceRequest is a constructor method which uses the prebuilt set of
// options as part of the request.
func NewServiceRequestFromDefaultOptions(opts *Options) *ServiceRequest {
	return &ServiceRequest{
		opts: opts,
	}
}

// SetMetro to use on the request.
func (r *ServiceRequest) SetMetro(metro string) {
	r.metro = metro
}

// SetTimeout sets how long to wait for the request before erroring.
func (r *ServiceRequest) SetTimeout(timeout time.Duration) {
	r.timeout = timeout
}

// SetHTTPClient sets the underlying HTTP client to use when performing the
// request.
func (r *ServiceRequest) SetHTTPClient(httpClient HTTPClient) {
	r.httpClient = httpClient
}

// Metrolink returns the full URI representing the API endpoint of a KraftCloud
// metro.
func (r *ServiceRequest) Metrolink(path string) string {
	if r.metro == "" {
		r.metro = r.opts.defaultMetro
	}

	// We discard the error because we are working with well-known path constant.
	u, _ := url.Parse(fmt.Sprintf(BaseV1FormatURL, r.metro))
	return u.JoinPath(path).String()
}

// DoRequest performs the request and hydrates a target type with response body.
func (r *ServiceRequest) DoRequest(ctx context.Context, method, url string, body io.Reader, target interface{}) error {
	req, err := http.NewRequestWithContext(ctx, method, r.Metrolink(url), body)
	if err != nil {
		return fmt.Errorf("error creating the request: %w", err)
	}

	resp, err := r.DoWithAuth(req)
	if err != nil {
		return fmt.Errorf("error performing the request: %w", err)
	}

	defer resp.Body.Close()

	if err := checkResponse(resp); err != nil {
		return fmt.Errorf("received an error in the response: %w", err)
	}

	if err := json.NewDecoder(resp.Body).Decode(target); err != nil {
		return fmt.Errorf("error parsing response: %v", err)
	}

	return nil
}

// DoWithAuth performs a request with headers defining the content type.  We
// also inject the authentication details.
func (r *ServiceRequest) DoWithAuth(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", r.GetBearerToken())
	req.Header.Set("Content-Type", "application/json")

	if r.httpClient == nil {
		r.httpClient = r.opts.http
	}

	return r.httpClient.Do(req)
}

// GetBearerToken uses the pre-defined token to construct the Bearer token used
// for authenticating requests.
func (r *ServiceRequest) GetBearerToken() string {
	return "Bearer " + r.opts.token
}
