// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package image

import kraftcloud "sdk.kraft.cloud/v0"

// ImageClient wraps the v1 Image client of Kraftcloud.
type ImageClient struct {
	kraftcloud.RESTClient
}

// NewDefaultClient creates a sensible, default Kraftcloud image API client.
func NewDefaultImageClient(user, token string) *ImageClient {
	return NewImageClient(kraftcloud.NewHTTPClient(), kraftcloud.BaseURL, user, token)
}

func NewImageClient(httpClient kraftcloud.HTTPClient, baseURL, user, token string) *ImageClient {
	return &ImageClient{
		RESTClient: kraftcloud.RESTClient{
			HTTPClient: httpClient,
			BaseURL:    baseURL,
			User:       user,
			Token:      token,
		},
	}
}
