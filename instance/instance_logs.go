// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package instance

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
)

// LogsResponse holds console output data, as returned by the API.
type LogsResponse struct {
	Status string `json:"status"`
	Data   struct {
		Instances []struct {
			Status string `json:"status"`
			UUID   string `json:"uuid"`
			Output string `json:"output"`
		} `json:"instances"`
	} `json:"data"`
}

// Logs returns the console output of the specified instance.
//
// See: https://docs.kraft.cloud/002-rest-api-v1-instances.html#console
func (i *InstanceClient) Logs(ctx context.Context, uuid string, maxLines int, latest bool) (string, error) {
	base := i.BaseURL + Endpoint
	endpoint := fmt.Sprintf("%s/%s/console", base, uuid)

	response := &LogsResponse{}

	if err := i.DoRequest(ctx, http.MethodGet, endpoint, nil, response); err != nil {
		return "", fmt.Errorf("performing the request: %w", err)
	}

	if response.Data.Instances == nil {
		return "", errors.New("instances data is nil")
	}

	if len(response.Data.Instances) == 0 {
		return "", errors.New("no instances data returned from the server")
	}

	outputB64 := response.Data.Instances[0].Output
	output, err := base64.StdEncoding.DecodeString(outputB64)
	if err != nil {
		return "", fmt.Errorf("decoding base64 console output: %w", err)
	}

	return string(output), nil
}
