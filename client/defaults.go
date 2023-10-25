// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package client

const (
	// BaseURL defines the default location of the Kraftcloud API.
	BaseURL = "https://api." + DefaultMetro + ".kraft.cloud/v1"
	// BaseV1FormatURL defines the default location of the KraftCloud API which is
	// formatted to allow setting the metro.
	BaseV1FormatURL = "https://api.%s.kraft.cloud/v1"
	// DefaultPort is the port the instance will listen on externally by default.
	DefaultPort = 443
	// DefaultMetro is set to a default node based in Frankfurt.
	DefaultMetro = "fra0"
)
