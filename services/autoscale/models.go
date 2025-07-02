// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2022, Unikraft GmbH and The KraftKit Authors.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package autoscale

import (
	"encoding/json"
	"fmt"

	"github.com/go-viper/mapstructure/v2"

	kcclient "sdk.kraft.cloud/client"
)

// CreateRequest is the payload for a POST /services/<uuid>/autoscale request.
// https://docs.kraft.cloud/api/v1/autoscale/#creating-an-autoscale-configuration
type CreateRequest struct {
	UUID           *string                 `json:"uuid,omitempty"` // mutually exclusive with name
	Name           *string                 `json:"name,omitempty"` // mutually exclusive with uuid
	MinSize        *int                    `json:"min_size,omitempty"`
	MaxSize        *int                    `json:"max_size,omitempty"`
	WarmupTimeMs   *int                    `json:"warmup_time_ms,omitempty"`
	CooldownTimeMs *int                    `json:"cooldown_time_ms,omitempty"`
	CreateArgs     CreateRequestCreateArgs `json:"create_args"`
	Policies       []Policy                `json:"policies,omitempty"`
}

// CreateRequestCreateArgs will eventually be replaced with
// instances.CreateRequest{}.
type CreateRequestCreateArgs struct {
	Template *CreateRequestTemplate `json:"template,omitempty"` // mutually exclusive with roms
	Roms     *CreateRequestRoms     `json:"roms,omitempty"`     // mutually exclusive with template
}

type CreateRequestRoms struct {
	Name  *string `json:"name,omitempty"`
	Image *string `json:"image,omitempty"`
}

type CreateRequestTemplate struct {
	UUID *string `json:"uuid,omitempty"` // mutually exclusive with name
	Name *string `json:"name,omitempty"` // mutually exclusive with uuid
}

// CreateResponseItem is a data item from a response to a POST /services/<uuid>/autoscale request.
// https://docs.kraft.cloud/api/v1/autoscale/#creating-an-autoscale-configuration
type CreateResponseItem struct {
	Status string `json:"status"`
	UUID   string `json:"uuid"`
	Name   string `json:"name"`

	kcclient.APIResponseCommon
}

// GetResponseItem is a data item from a response to a GET /services/<uuid>/autoscale request.
// https://docs.kraft.cloud/api/v1/autoscale/#getting-an-existing-autoscale-configuration
type GetResponseItem struct {
	Status         string               `json:"status"`
	UUID           string               `json:"uuid"`
	Name           string               `json:"name"`
	Enabled        bool                 `json:"enabled"`
	MinSize        *int                 `json:"min_size"`         // only if enabled
	MaxSize        *int                 `json:"max_size"`         // only if enabled
	WarmupTimeMs   *int                 `json:"warmup_time_ms"`   // only if enabled
	CooldownTimeMs *int                 `json:"cooldown_time_ms"` // only if enabled
	Template       *GetResponseTemplate `json:"template"`         // only if enabled
	Policies       []Policy             `json:"policies"`

	kcclient.APIResponseCommon
}

type GetResponseTemplate struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (i *GetResponseItem) UnmarshalJSON(data []byte) error {
	unstructured := make(map[string]interface{})
	err := json.Unmarshal(data, &unstructured)
	if err != nil {
		return fmt.Errorf("deserializing response JSON data: %w", err)
	}

	if policies, ok := unstructured["policies"].([]interface{}); ok {
		for _, p := range policies {
			policy, ok := p.(map[string]interface{})
			if !ok {
				continue
			}

			policyType, ok := policy["type"]
			if !ok {
				continue
			}

			switch PolicyType(policyType.(string)) {
			case PolicyTypeStep:
				var stepPolicy StepPolicy
				if err := mapstructure.Decode(policy, &stepPolicy); err != nil {
					return fmt.Errorf("initializing step policy from response data: %w", err)
				}

				i.Policies = append(i.Policies, stepPolicy)
			case PolicyTypeOnDemand:
				var onDemandPolicy OnDemandPolicy
				if err := mapstructure.Decode(policy, &onDemandPolicy); err != nil {
					return fmt.Errorf("initializing on-demand policy from response data: %w", err)
				}

				i.Policies = append(i.Policies, onDemandPolicy)
			}
		}
	}

	// Do not decode the "policies" field into the struct as it has already been
	// processed above.
	delete(unstructured, "policies")

	return mapstructure.Decode(unstructured, &i)
}

// DeleteResponseItem is a data item from a response to a DELETE /services/<uuid>/autoscale request.
// https://docs.kraft.cloud/api/v1/autoscale/#deleting-an-autoscale-configuration
type DeleteResponseItem struct {
	Status string `json:"status"`
	UUID   string `json:"uuid"`
	Name   string `json:"name"`

	kcclient.APIResponseCommon
}

// AddPolicyResponseItem is a data item from a response to a POST /services/<uuid>/autoscale/policies request.
// https://docs.kraft.cloud/api/v1/autoscale/#adding-an-autoscale-policy
type AddPolicyResponseItem struct {
	Status string `json:"status"`
	UUID   string `json:"uuid"`
	Name   string `json:"name"`

	kcclient.APIResponseCommon
}

// GetPolicyResponseItem is a data item from a response to a GET /services/<uuid>/autoscale/policies request.
// https://docs.kraft.cloud/api/v1/autoscale/#getting-the-configuration-of-an-autoscale-policy
type GetPolicyResponseItem struct {
	Status   string   `json:"status"`
	Policies []Policy `json:"policies"`

	kcclient.APIResponseCommon
}

// UnmarshalJSON implements json.Unmarshaler.
func (i *GetPolicyResponseItem) UnmarshalJSON(data []byte) error {
	unstructured := make(map[string]interface{})
	err := json.Unmarshal(data, &unstructured)
	if err != nil {
		return fmt.Errorf("deserializing response JSON data: %w", err)
	}

	var policies []interface{}
	if _, ok := unstructured["policies"].([]interface{}); ok {
		policies = unstructured["policies"].([]interface{})
	} else {
		if len(unstructured) != 0 {
			policies = []interface{}{unstructured}
		} else {
			return fmt.Errorf("expected 'policies' field to be an array or object, got: %T", unstructured)
		}
	}

	for _, p := range policies {
		policy, ok := p.(map[string]interface{})
		if !ok {
			continue
		}

		policyType, ok := policy["type"]
		if !ok {
			continue
		}

		switch PolicyType(policyType.(string)) {
		case PolicyTypeStep:
			var stepPolicy StepPolicy

			if err := mapstructure.Decode(policy, &stepPolicy); err != nil {
				return fmt.Errorf("initializing step policy from response data: %w", err)
			}

			// NOTE(craciunoiuc): Dirty hack to handle fields that are pointers
			// This failed without it as integer responses could not be
			// unmarshalled into a pointer with the mapstructure package.
			remarshalled, err := json.Marshal(policy["steps"])
			if err != nil {
				return fmt.Errorf("marshalling steps from response data: %w", err)
			}
			json.Unmarshal(remarshalled, &stepPolicy.Steps)

			remarshalled, err = json.Marshal(policy["adjustment_type"])
			if err != nil {
				return fmt.Errorf("marshalling steps from response data: %w", err)
			}
			json.Unmarshal(remarshalled, &stepPolicy.AdjustmentType)

			i.Policies = append(i.Policies, stepPolicy)
		case PolicyTypeOnDemand:
			var onDemandPolicy OnDemandPolicy
			if err := mapstructure.Decode(policy, &onDemandPolicy); err != nil {
				return fmt.Errorf("initializing on-demand policy from response data: %w", err)
			}

			i.Policies = append(i.Policies, onDemandPolicy)
		}
	}

	// Do not decode the "policies" field into the struct as it has already been
	// processed above.
	delete(unstructured, "policies")

	return mapstructure.Decode(unstructured, &i)
}

// DeletePolicyResponseItem is a data item from a response to a DELETE /services/<uuid>/autoscale/policies request.
// https://docs.kraft.cloud/api/v1/autoscale/#deleting-an-autoscale-policy
type DeletePolicyResponseItem struct {
	Status string `json:"status"`
	Name   string `json:"name"`

	kcclient.APIResponseCommon
}
