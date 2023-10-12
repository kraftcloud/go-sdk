// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package instance

import (
	kraftcloud "sdk.kraft.cloud/v0"
)

// InstanceClient is a basic wrapper around for the v1 instance services of
// KraftCloud.
type InstanceClient struct {
	kraftcloud.RESTClient
}

// NewDefaultClient creates a sensible, default Kraftcloud instance API client.
func NewDefaultInstanceClient(user, token string) *InstanceClient {
	return NewInstanceClient(kraftcloud.NewHTTPClient(), kraftcloud.BaseURL, user, token)
}

func NewInstanceClient(httpClient kraftcloud.HTTPClient, baseURL, user, token string) *InstanceClient {
	return &InstanceClient{
		RESTClient: kraftcloud.RESTClient{
			HTTPClient: httpClient,
			BaseURL:    baseURL,
			User:       user,
			Token:      token,
		},
	}
}
