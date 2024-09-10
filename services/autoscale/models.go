// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2022, Unikraft GmbH and The KraftKit Authors.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package autoscale

import (
	"encoding/json"
	"fmt"

	ukcclient "sdk.kraft.cloud/client"
)

// CreateRequest is the payload for a POST /services/<uuid>/autoscale request.
// https://docs.kraft.cloud/api/v1/autoscale/#creating-an-autoscale-configuration
type CreateRequest struct {
	UUID           *string             `json:"uuid,omitempty"` // mutually exclusive with name
	Name           *string             `json:"name,omitempty"` // mutually exclusive with uuid
	MinSize        *int                `json:"min_size,omitempty"`
	MaxSize        *int                `json:"max_size,omitempty"`
	WarmupTimeMs   *int                `json:"warmup_time_ms,omitempty"`
	CooldownTimeMs *int                `json:"cooldown_time_ms,omitempty"`
	Master         CreateRequestMaster `json:"master"`
	Policies       []Policy            `json:"policies,omitempty"`
}

type CreateRequestMaster struct {
	UUID *string `json:"uuid,omitempty"` // mutually exclusive with name
	Name *string `json:"name,omitempty"` // mutually exclusive with uuid
}

// CreateResponseItem is a data item from a response to a POST /services/<uuid>/autoscale request.
// https://docs.kraft.cloud/api/v1/autoscale/#creating-an-autoscale-configuration
type CreateResponseItem struct {
	Status string `json:"status"`
	UUID   string `json:"uuid"`
	Name   string `json:"name"`

	ukcclient.APIResponseCommon
}

// GetResponseItem is a data item from a response to a GET /services/<uuid>/autoscale request.
// https://docs.kraft.cloud/api/v1/autoscale/#getting-an-existing-autoscale-configuration
type GetResponseItem struct {
	Status         string             `json:"status"`
	UUID           string             `json:"uuid"`
	Name           string             `json:"name"`
	Enabled        bool               `json:"enabled"`
	MinSize        *int               `json:"min_size"`         // only if enabled
	MaxSize        *int               `json:"max_size"`         // only if enabled
	WarmupTimeMs   *int               `json:"warmup_time_ms"`   // only if enabled
	CooldownTimeMs *int               `json:"cooldown_time_ms"` // only if enabled
	Master         *GetResponseMaster `json:"master"`           // only if enabled
	Policies       []Policy           `json:"policies"`

	ukcclient.APIResponseCommon
}

type GetResponseMaster struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (i *GetResponseItem) UnmarshalJSON(data []byte) error {
	unstructured := make(map[string]any)
	err := json.Unmarshal(data, &unstructured)
	if err != nil {
		return fmt.Errorf("deserializing response JSON data: %w", err)
	}

	var pols []Policy
	if _, ok := unstructured[attrPolicies]; ok {
		polsData, err := readSliceAttribute(unstructured, attrPolicies)
		if err != nil {
			return err
		}

		pols = make([]Policy, 0, len(polsData))
		for _, p := range polsData {
			polData, ok := p.(map[string]any)
			if !ok {
				return fmt.Errorf("'%s' attribute item is not a map (%T)", attrPolicies, p)
			}

			typ, err := readStringAttribute(polData, polAttrType)
			if err != nil {
				return err
			}

			switch PolicyType(typ) {
			case PolicyTypeStep:
				p, err := stepPolicyFromUnstructured(polData)
				if err != nil {
					return fmt.Errorf("initializing step policy from response data: %w", err)
				}
				pols = append(pols, p)
			default:
				return fmt.Errorf("unsupported policy type %q", typ)
			}

		}
	}

	delete(unstructured, attrPolicies)
	if data, err = json.Marshal(unstructured); err != nil {
		return fmt.Errorf("re-serializing response data to JSON: %w", err)
	}

	type GetResponseItem_ GetResponseItem // prevent recursive call to (*GetResponseItem).UnmarshalJSON()
	ii := &GetResponseItem_{}

	if err = json.Unmarshal(data, ii); err != nil {
		return fmt.Errorf("re-deserializing response JSON data: %w", err)
	}

	*i = GetResponseItem(*ii)
	i.Policies = pols

	return nil
}

// DeleteResponseItem is a data item from a response to a DELETE /services/<uuid>/autoscale request.
// https://docs.kraft.cloud/api/v1/autoscale/#deleting-an-autoscale-configuration
type DeleteResponseItem struct {
	Status string `json:"status"`
	UUID   string `json:"uuid"`
	Name   string `json:"name"`

	ukcclient.APIResponseCommon
}

// AddPolicyResponseItem is a data item from a response to a POST /services/<uuid>/autoscale/policies request.
// https://docs.kraft.cloud/api/v1/autoscale/#adding-an-autoscale-policy
type AddPolicyResponseItem struct {
	Status string `json:"status"`
	UUID   string `json:"uuid"`
	Name   string `json:"name"`

	ukcclient.APIResponseCommon
}

// GetPolicyResponseItem is a data item from a response to a GET /services/<uuid>/autoscale/policies request.
// https://docs.kraft.cloud/api/v1/autoscale/#getting-the-configuration-of-an-autoscale-policy
type GetPolicyResponseItem struct {
	Status  string
	Enabled bool
	Details Policy

	ukcclient.APIResponseCommon
}

// UnmarshalJSON implements json.Unmarshaler.
func (i *GetPolicyResponseItem) UnmarshalJSON(data []byte) error {
	d := make(map[string]any)
	err := json.Unmarshal(data, &d)
	if err != nil {
		return fmt.Errorf("deserializing response JSON data: %w", err)
	}

	i.Status, err = readStringAttribute(d, "status")
	if err != nil {
		return err
	}

	i.Enabled, err = readBoolAttribute(d, "enabled")
	if err != nil {
		return err
	}

	typ, err := readStringAttribute(d, polAttrType)
	if err != nil {
		return err
	}

	switch PolicyType(typ) {
	case PolicyTypeStep:
		p, err := stepPolicyFromUnstructured(d)
		if err != nil {
			return fmt.Errorf("initializing step policy from response data: %w", err)
		}
		i.Details = p
	default:
		return fmt.Errorf("unsupported policy type %q", typ)
	}

	return nil
}

// DeletePolicyResponseItem is a data item from a response to a DELETE /services/<uuid>/autoscale/policies request.
// https://docs.kraft.cloud/api/v1/autoscale/#deleting-an-autoscale-policy
type DeletePolicyResponseItem struct {
	Status string `json:"status"`
	Name   string `json:"name"`

	ukcclient.APIResponseCommon
}
