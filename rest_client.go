// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2022, Unikraft GmbH and The KraftKit Authors.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package kraftcloud

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// HTTPClient interface abstracts a generic HTTP request issuing client.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// RESTClient is our base API Client implementing the common HTTP/authentication
// behaviours.
type RESTClient struct {
	HTTPClient HTTPClient
	BaseURL    string
	User       string
	Token      string
}

// DoRequest performs the request and hydrates a target type with response body.
func (c *RESTClient) DoRequest(ctx context.Context, method, endpoint string, body io.Reader, v interface{}) error {
	req, err := http.NewRequestWithContext(ctx, method, endpoint, body)
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}

	resp, err := c.doWithAuth(req)
	if err != nil {
		return fmt.Errorf("issuing request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return checkResponse(resp)
	}

	return json.NewDecoder(resp.Body).Decode(v)
}

func (c *RESTClient) doWithAuth(req *http.Request) (*http.Response, error) {
	bearer, err := c.getBearerToken()
	if err != nil {
		return nil, fmt.Errorf("getting the bearer token: %w", err)
	}

	req.Header.Set("Authorization", bearer)
	req.Header.Set("Content-Type", "application/json")

	return c.HTTPClient.Do(req)
}

func (c *RESTClient) getBearerToken() (string, error) {
	if c.User == "" || c.Token == "" {
		return "", errors.New("no auth details provided")
	}
	bearer := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", c.User, c.Token)))
	return fmt.Sprintf("Bearer %s", bearer), nil
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

func checkResponse(resp *http.Response) error {
	if resp.StatusCode == http.StatusOK {
		return nil
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("reading response body: %w", err)
	}
	return &Error{StatusCode: resp.StatusCode, Message: string(bodyBytes)}
}
