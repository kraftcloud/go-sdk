// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package kraftcloud

const (
	// BaseURL defines the default location of the Kraftcloud API.
	BaseURL = "https://api.fra0.kraft.cloud/v1"
	// BaseV1FormatURL defines the default location of the KraftCloud API which is
	// formatted to allow setting the metro.
	BaseV1FormatURL = "https://api.%s.kraft.cloud/v1"
	// DefaultPort is the port the instance will listen on externally by default.
	DefaultPort = 443
	// DefaultHandler sets the connection handler. The API only accepts "tls" for
	// now.
	DefaultHandler = "tls"
	// DefaultAutoStart is the default autostart value - whether the instance will
	// start immediately after creation
	DefaultAutoStart = true
)
