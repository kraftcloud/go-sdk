// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// ErrorResponse is the list of errors that have occurred during the invocation
// of the API call.
type ErrorResponse struct {
	Status int `json:"status"`
}

// ServiceResponse embodies the the API response for an invocation to a service
// on UnikraftCloud.  It uses standard HTTP response codes to indicate success or
// failure. In addition, the response body contains more details about the
// result of the operation in a JSON object. On success the data member contains
// an array of objects with one object per result. The array is named according
// to the type of object that the request is operating on. For example, when
// working with instances the response contains an instances array.
//
// See: https://docs.kraft.cloud/api/v1/
type ServiceResponse[T APIResponseDataEntry] struct {
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

	// Buffer holding the raw API response body.
	body bytes.Buffer
}

// ServiceResponseData is the embedded list of structures defined by T.  The
// results are always available at the attribute `Entries` and uses a custom
// JSON unmarshaler to determine the JSON tag associated with these entries.
type ServiceResponseData[T APIResponseDataEntry] struct {
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

// AllOrErr returns the all data entrypoints or an error if it is not
// available.
func (r *ServiceResponse[T]) AllOrErr() ([]T, error) {
	if r == nil {
		return nil, errors.New("response is nil")
	}

	if r.Data.Entries == nil {
		return nil, errors.New("data entries are nil")
	}
	return r.Data.Entries, r.aggregateErrors()
}

// RawBody returns the raw API response body.
func (r *ServiceResponse[T]) RawBody() []byte {
	b := make([]byte, r.body.Len())
	_, _ = r.body.Read(b) // cannot fail on a bytes.Buffer
	return b
}

// rawBodyHolder is implemented by types that can hold raw API response bodies.
type rawBodyHolder interface {
	storeBody(io.Reader) (n int64, err error)
}

// storeBody implements rawBodyHolder.
func (r *ServiceResponse[T]) storeBody(br io.Reader) (int64, error) {
	return io.Copy(&r.body, br)
}

// aggregateErrors returns an aggregate of all errors returned in an API response.
func (r *ServiceResponse[T]) aggregateErrors() error {
	if !(r.Status == "error" || r.Status == "partial_success") {
		return nil
	}
	errs := make([]error, 0, len(r.Data.Entries)+1)
	errs = append(errs, errors.New(r.Message))
	for _, entry := range r.Data.Entries {
		if entry := entry.ErrorAttributes(); entry.Error != nil {
			errs = append(errs, fmt.Errorf("%s (code=%d)", entry.Message, *entry.Error))
		}
	}
	return errors.Join(errs...)
}

// APIResponseCommon contains attributes common to all API responses, namely
// the attributes which are returned either on error or partial success.
// https://docs.kraft.cloud/api/v1/#api-responses
type APIResponseCommon struct {
	Status  string        `json:"status"`
	Message string        `json:"message"`
	Error   *APIHTTPError `json:"error"`
}

// ErrorAttributes implements APIResponseDataEntry.
func (c APIResponseCommon) ErrorAttributes() APIResponseCommon {
	return c
}

// APIResponseDataEntry provides access to data common to all data entries
// returned in API responses.
type APIResponseDataEntry interface {
	ErrorAttributes() APIResponseCommon
}
