// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2022, Unikraft GmbH and The KraftKit Authors.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package kraftcloud

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestCheckResponse(t *testing.T) {
	tests := []struct {
		name            string
		inputStatusCode int
		inputBody       string
		expectErr       bool
		expectErrMsg    string
	}{
		{
			name:            "test with status OK",
			inputStatusCode: http.StatusOK,
			expectErr:       false,
		},
		{
			name:            "test with status BadRequest",
			inputStatusCode: http.StatusBadRequest,
			inputBody:       "Bad Request",
			expectErr:       true,
			expectErrMsg:    "API error: status code 400, message: Bad Request",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &http.Response{
				StatusCode: tt.inputStatusCode,
				Body:       io.NopCloser(strings.NewReader(tt.inputBody)),
			}

			err := checkResponse(resp)
			if tt.expectErr {
				if err == nil {
					t.Errorf("expected an error, got nil")
					return
				}

				if err.Error() != tt.expectErrMsg {
					t.Errorf("expected error message %q, got %q", tt.expectErrMsg, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("expected no error, got %v", err)
				}
			}
		})
	}
}
