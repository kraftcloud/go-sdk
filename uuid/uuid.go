// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2024, Unikraft GmbH and The KraftKit Authors.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// Package uuid allows inspecting RFC 4122 UUIDs.
package uuid

import "github.com/google/uuid"

// IsValid returns whether the given string is a valid UUID.
func IsValid(s string) bool {
	if len(s) != 36 { // standard UUID format, with hyphens
		return false
	}
	return uuid.Validate(s) == nil
}
