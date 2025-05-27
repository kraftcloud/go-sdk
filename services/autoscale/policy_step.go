// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2022, Unikraft GmbH and The KraftKit Authors.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package autoscale

import (
	"bytes"
	"encoding/json"
)

// StepPolicy JSON attributes
const (
	spolAttrMetric  = "metric"
	spolAttrAdjType = "adjustment_type"
	spolAttrSteps   = "steps"
)

// StepPolicy is a Step autoscale policy.
// https://docs.kraft.cloud/api/v1/autoscale/#step-policy
type StepPolicy struct {
	Name           string
	Metric         PolicyMetric
	AdjustmentType AdjustmentType
	Steps          []Step
}

// Type implements Policy.
func (p StepPolicy) Type() PolicyType {
	return PolicyTypeStep
}

// AddSteps implements Policy.
func (p *StepPolicy) AddSteps(steps ...Step) {
	p.Steps = append(p.Steps, steps...)
}

// MarshalJSON implements json.Marshaler.
func (p StepPolicy) MarshalJSON() ([]byte, error) {
	var jsonData bytes.Buffer
	jsonData.WriteByte('{')
	jsonData.WriteString(`"name":"` + p.Name + `"`)
	jsonData.WriteString(`,"type":"` + string(PolicyTypeStep) + `"`)
	jsonData.WriteString(`,"` + spolAttrMetric + `":"` + string(p.Metric) + `"`)
	jsonData.WriteString(`,"` + spolAttrAdjType + `":"` + string(p.AdjustmentType) + `"`)
	jsonData.WriteString(`,"` + spolAttrSteps + `":[`)
	for i, s := range p.Steps {
		b, _ := json.Marshal(s) // cannot error
		jsonData.Write(b)
		if i < len(p.Steps)-1 {
			jsonData.WriteByte(',')
		}
	}
	jsonData.WriteByte(']')
	jsonData.WriteByte('}')
	return jsonData.Bytes(), nil
}

// Step is a step in a StepPolicy.
type Step struct {
	Adjustment int  `json:"adjustment"`
	LowerBound *int `json:"lower_bound,omitempty"` // optional, can be nil
	UpperBound *int `json:"upper_bound,omitempty"` // optional, can be nil
}
