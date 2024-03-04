// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2022, Unikraft GmbH and The KraftKit Authors.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package autoscale

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

// StepPolicy JSON attributes
const (
	spolAttrMetric  = "metric"
	spolAttrAdjType = "adjustment_type"
	spolAttrSteps   = "steps"

	stepAttrAdj     = "adjustment"
	stepAttrLoBound = "lower_bound"
	stepAttrUpBound = "upper_bound"
)

// StepPolicy is a Step autoscale policy.
// https://docs.kraft.cloud/api/v1/autoscale/#step-policy
type StepPolicy struct {
	Name           string
	Metric         PolicyMetric
	AdjustmentType AdjustmentType
	Steps          []Step
}

var _ Policy = (*StepPolicy)(nil)

// Type implements Policy.
func (p StepPolicy) Type() PolicyType {
	return PolicyTypeStep
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
	Adjustment int
	LowerBound *int
	UpperBound *int
}

// MarshalJSON implements json.Marshaler.
func (s Step) MarshalJSON() ([]byte, error) {
	var jsonData bytes.Buffer
	jsonData.WriteByte('{')
	jsonData.WriteString(`"` + stepAttrAdj + `":` + strconv.Itoa(s.Adjustment))
	if s.LowerBound != nil {
		jsonData.WriteString(`,"` + stepAttrLoBound + `":` + strconv.Itoa(*s.LowerBound))
	}
	if s.UpperBound != nil {
		jsonData.WriteString(`,"` + stepAttrUpBound + `":` + strconv.Itoa(*s.UpperBound))
	}
	jsonData.WriteByte('}')
	return jsonData.Bytes(), nil
}

// stepPolicyFromUnstructured returns a StepPolicy which attributes are
// populated from the given unstructured data.
func stepPolicyFromUnstructured(data map[string]any) (*StepPolicy, error) {
	name, err := readStringAttribute(data, polAttrName)
	if err != nil {
		return nil, err
	}
	metric, err := readStringAttribute(data, spolAttrMetric)
	if err != nil {
		return nil, err
	}
	adjType, err := readStringAttribute(data, spolAttrAdjType)
	if err != nil {
		return nil, err
	}

	p := &StepPolicy{
		Name:           name,
		Metric:         PolicyMetric(metric),
		AdjustmentType: AdjustmentType(adjType),
	}

	stepsData, err := readSliceAttribute(data, spolAttrSteps)
	if err != nil {
		return nil, err
	}

	p.Steps = make([]Step, 0, len(stepsData))
	for _, s := range stepsData {
		stepData, ok := s.(map[string]any)
		if !ok {
			return nil, fmt.Errorf("'%s' attribute item is not a map (%T)", spolAttrSteps, s)
		}

		adj, err := readIntAttribute(stepData, stepAttrAdj)
		if err != nil {
			return nil, err
		}

		s := Step{
			Adjustment: adj,
		}

		if _, ok := stepData[stepAttrLoBound]; ok {
			b, err := readIntAttribute(stepData, stepAttrLoBound)
			if err != nil {
				return nil, err
			}
			s.LowerBound = &b
		}
		if _, ok := stepData[stepAttrUpBound]; ok {
			b, err := readIntAttribute(stepData, stepAttrUpBound)
			if err != nil {
				return nil, err
			}
			s.UpperBound = &b
		}

		p.Steps = append(p.Steps, s)
	}

	return p, nil
}

func readStringAttribute(data map[string]any, attrName string) (string, error) {
	v, ok := data[attrName]
	if !ok {
		return "", errors.New("missing '" + attrName + "' attribute")
	}
	strV, ok := v.(string)
	if !ok {
		return "", fmt.Errorf("'%s' attribute is not a string (%T)", attrName, v)
	}

	return strV, nil
}

func readIntAttribute(data map[string]any, attrName string) (int, error) {
	v, ok := data[attrName]
	if !ok {
		return -1, errors.New("missing '" + attrName + "' attribute")
	}
	floatV, ok := v.(float64) // numbers deserialized from JSON as 'any' are float64
	if !ok {
		return -1, fmt.Errorf("'%s' attribute is not a number (%T)", attrName, v)
	}

	return int(floatV), nil
}

func readBoolAttribute(data map[string]any, attrName string) (bool, error) {
	v, ok := data[attrName]
	if !ok {
		return false, errors.New("missing '" + attrName + "' attribute")
	}
	boolV, ok := v.(bool)
	if !ok {
		return false, fmt.Errorf("'%s' attribute is not a boolean (%T)", attrName, v)
	}

	return boolV, nil
}

func readSliceAttribute(data map[string]any, attrName string) ([]any, error) {
	v, ok := data[attrName]
	if !ok {
		return nil, errors.New("missing '" + attrName + "' attribute")
	}
	sliceV, ok := v.([]any)
	if !ok {
		return nil, fmt.Errorf("'%s' attribute is not a slice (%T)", attrName, v)
	}

	return sliceV, nil
}
