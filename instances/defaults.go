// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package instances

const (
	// DefaultHandler sets the connection handler. The API only accepts "tls" for
	// now.
	DefaultHandler = "tls"

	// DefaultAutoStart is the default autostart value - whether the instance will
	// start immediately after creation
	DefaultAutoStart = true
)
