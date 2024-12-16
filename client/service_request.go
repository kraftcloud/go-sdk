// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"sdk.kraft.cloud/client/errors"
	"sdk.kraft.cloud/client/httpclient"
	"sdk.kraft.cloud/client/options"
)

// ServiceRequest is the utility structure for performing individual requests to
// a service location at UnikraftCloud.
type ServiceRequest struct {
	// constructors must ensure that opts is non-nil, and that all its nested
	// fields are populated to at least a default value
	opts *options.Options

	metro      string
	httpClient httpclient.HTTPClient
	timeout    time.Duration
}

// NewServiceRequestFromDefaultOptions is a constructor method which uses the
// prebuilt set of options as part of the request.
func NewServiceRequestFromDefaultOptions(opts *options.Options) *ServiceRequest {
	return &ServiceRequest{
		opts: opts,
	}
}

// WithMetro returns a ServiceRequest that uses the given metro in API
// requests.
func (r *ServiceRequest) WithMetro(m string) *ServiceRequest {
	rcpy := r.clone()
	rcpy.metro = m
	return rcpy
}

// WithTimeout returns a ServiceRequest that uses the specified timeout
// duration in API requests.
func (r *ServiceRequest) WithTimeout(to time.Duration) *ServiceRequest {
	rcpy := r.clone()
	rcpy.timeout = to
	return rcpy
}

// WithHTTPClient returns a ServiceRequest which performs API requests using
// the given HTTPClient.
func (r *ServiceRequest) WithHTTPClient(hc httpclient.HTTPClient) *ServiceRequest {
	rcpy := r.clone()
	rcpy.httpClient = hc
	return rcpy
}

// Metro returns the metro that this request will perform against.
func (r *ServiceRequest) Metro() string {
	return r.metro
}

// Metrolink returns the full URI representing the API endpoint of a UnikraftCloud
// metro.
func (r *ServiceRequest) Metrolink(path string) string {
	// If the metro contains a full URL, quantified by the presence of a scheme,
	// we assume it is a full URL to a metro and we use it as is.
	if strings.Contains(r.metro, "://") {
		p, _ := url.JoinPath(r.metro, path)
		return p
	}

	m := r.opts.DefaultMetro()
	if r.metro != "" {
		m = r.metro
	}

	// We discard the error because we are working with well-known path constant.
	u, _ := url.Parse(fmt.Sprintf(BaseV1FormatURL, m))
	return u.JoinPath(path).String()
}

// DoRequest performs the request and hydrates a target type with response body.
func (r *ServiceRequest) DoRequest(ctx context.Context, method, url string, body io.Reader, target any) error {
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

	bodyReader, err := maybeStoreRawBody(target, resp.Body)
	if err != nil {
		return err
	}

	if err := json.NewDecoder(bodyReader).Decode(target); err != nil {
		return fmt.Errorf("error parsing response: %w", err)
	}

	return nil
}

// DoWithAuth performs a request with headers defining the content type.  We
// also inject the authentication details.
func (r *ServiceRequest) DoWithAuth(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", r.GetBearerToken())
	req.Header.Set("Content-Type", "application/json")

	hc := r.opts.HTTPClient()
	if r.httpClient != nil {
		hc = r.httpClient
	}

	return hc.Do(req)
}

// GetBearerToken uses the pre-defined token to construct the header used for
// authenticating requests.
func (r *ServiceRequest) GetBearerToken() string {
	return "Bearer " + r.opts.Token()
}

// GetToken uses the pre-defined token to construct the Bearer token used
// for authenticating requests.
func (r *ServiceRequest) GetToken() string {
	return r.opts.Token()
}

// clone returns a shallow copy of r.
func (r *ServiceRequest) clone() *ServiceRequest {
	rcpy := *r
	return &rcpy
}

func checkResponse(resp *http.Response) error {
	if resp.StatusCode == http.StatusOK {
		return nil
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("reading response body: %w", err)
	}
	return &errors.Error{StatusCode: resp.StatusCode, Message: string(bodyBytes)}
}

// maybeStoreRawBody stores the response body from b into target if target
// supports it, and returns an io.Reader that can be used for future reads of
// the given response body.
// After passing b to this function, b should no longer be read from.
func maybeStoreRawBody(target any, b io.Reader) (io.Reader, error) {
	bh, ok := target.(rawBodyHolder)
	if !ok {
		return b, nil
	}

	body, err := io.ReadAll(b)
	if err != nil {
		return nil, fmt.Errorf("reading response body: %w", err)
	}
	if _, err = bh.storeBody(bytes.NewReader(body)); err != nil {
		return nil, fmt.Errorf("storing response body: %w", err)
	}
	return bytes.NewReader(body), nil
}
