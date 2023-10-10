// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2022, Unikraft GmbH and The KraftKit Authors.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.
package util

import "testing"

func TestNormalizeImageName(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expected      string
		expectedError string
	}{
		{
			name:     "Normal case",
			input:    "index.unikraft.io/jayc.unikraft.io:python",
			expected: "index.unikraft.io/jayc.unikraft.io/python",
		},
		{
			name:          "More than one colon",
			input:         "index.unikraft.io:jayc.unikraft.io:python",
			expectedError: "more than one semicolon is present in the image name: index.unikraft.io:jayc.unikraft.io:python",
		},
		{
			name:     "No colon case",
			input:    "index.unikraft.io/jayc.unikraft.io/python",
			expected: "index.unikraft.io/jayc.unikraft.io/python",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := NormalizeImageName(test.input)
			if err != nil {
				if test.expectedError == "" {
					t.Errorf("Unexpected error: %v", err)
				} else if err.Error() != test.expectedError {
					t.Errorf("Expected error %q but got %q", test.expectedError, err.Error())
				}
				return
			}

			if test.expectedError != "" {
				t.Errorf("Expected error %q but got nil", test.expectedError)
				return
			}

			if result != test.expected {
				t.Errorf("Expected %q but got %q", test.expected, result)
			}
		})
	}
}
