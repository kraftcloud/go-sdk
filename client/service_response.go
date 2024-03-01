// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package client

import (
	"encoding/json"
	"errors"
	"fmt"
)

// ErrorResponse is the list of errors that have occurred during the invocation
// of the API call.
type ErrorResponse struct {
	Status int `json:"status"`
}

// ServiceResponse embodies the the API response for an invocation to a service
// on KraftCloud.  It uses standard HTTP response codes to indicate success or
// failure. In addition, the response body contains more details about the
// result of the operation in a JSON object. On success the data member contains
// an array of objects with one object per result. The array is named according
// to the type of object that the request is operating on. For example, when
// working with instances the response contains an instances array.
//
// See: https://docs.kraft.cloud/api/v1/
type ServiceResponse[T any] struct {
	// Status contains the top-level information about a server response, and
	// returns either `success`, `partial_success` or `error`.
	Status string `json:"status,omitempty"`

	// Message contains the error message either on `partial_success` or `error`.
	Message string `json:"message,omitempty"`

	// Errors are the list of errors that have occurred.
	Errors []ErrorResponse `json:"errors,omitempty"`

	// On a successful response, the data element is returned with relevant
	// information.
	Data ServiceResponseData[T] `json:"data,omitempty"`
}

// ServiceResponseData is the embedded list of structures defined by T.  The
// results are always available at the attribute `Entries` and uses a custom
// JSON unmarshaler to determine the JSON tag associated with these entries.
type ServiceResponseData[T any] struct {
	Entries []T
}

// UnmarshalJSON implements json.Unmarshaler.
func (d *ServiceResponseData[T]) UnmarshalJSON(b []byte) error {
	var res map[string]json.RawMessage

	if err := json.Unmarshal(b, &res); err != nil {
		return err
	}

	if len(res) == 0 {
		d.Entries = make([]T, 0)
		return nil
	}

	if len(res) > 1 {
		var keys []string
		for key := range res {
			keys = append(keys, key)
		}
		return fmt.Errorf("cannot unmarshal service response data with multiple top-level keys %v", keys)
	}

	for _, entries := range res {
		return json.Unmarshal(entries, &d.Entries)
	}

	return nil // Unreachable
}

// FirstOrErr returns the first data entrypoint or an error if it is not
// available.
func (r *ServiceResponse[T]) FirstOrErr() (*T, error) {
	entries, err := r.AllOrErr()
	if len(entries) > 0 {
		return &entries[0], err
	}

	return nil, err
}

// FirstOrErr returns the all data entrypoints or an error if it is not
// available.
func (r *ServiceResponse[T]) AllOrErr() ([]T, error) {
	if r == nil {
		return nil, errors.New("response is nil")
	}

	if r.Status == "error" {
		return r.Data.Entries, fmt.Errorf(r.Message)
	}

	if r.Data.Entries == nil {
		return nil, errors.New("data entries are nil")
	}

	return r.Data.Entries, nil
}
