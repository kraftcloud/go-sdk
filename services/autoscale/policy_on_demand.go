// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH and The KraftKit Authors.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package autoscale

import "encoding/json"

type OnDemandPolicy struct {
	Name      string `json:"name"`
	Enabled   bool   `json:"enabled,omitzero"`
	Exclusive bool   `json:"exclusive,omitzero"`
}

// Type implements Policy.
func (p OnDemandPolicy) Type() PolicyType {
	return PolicyTypeOnDemand
}

func (p OnDemandPolicy) MarshalJSON() ([]byte, error) {
	data := make(map[string]interface{})

	data["name"] = p.Name
	data["type"] = p.Type()

	return json.Marshal(data)
}
