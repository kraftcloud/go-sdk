// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package client

import (
	"time"

	"sdk.kraft.cloud/client/httpclient"
)

// ServiceClient is an interface of mandatory methods that a service must
// implement.  These methods are used to customize the request just-in-time,
// enabling deep customization of each request without having to re-instantiate
// the client.
type ServiceClient[T any] interface {
	// WithMetro sets the just-in-time metro to use when connecting to the
	// KraftCloud API.
	WithMetro(string) T

	// WithTimeout sets the timeout when making the request.
	WithTimeout(time.Duration) T

	// WithHTTPClient overwrites the base HTTP client.
	WithHTTPClient(httpclient.HTTPClient) T
}
