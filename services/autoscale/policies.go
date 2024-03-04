// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2022, Unikraft GmbH and The KraftKit Authors.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package autoscale

// Policy is implemented by all types that represent an autoscale policy.
type Policy interface {
	Type() PolicyType
}

// JSON attributes
const (
	attrPolicies = "policies"
	polAttrName  = "name"
	polAttrType  = "type"
)
