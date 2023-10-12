// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package instance

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	kraftcloud "sdk.kraft.cloud/v0"
)

// StopInstancePayload carries the data used by stop instance requests.
type StopInstancePayload struct {
	DrainTimeoutMS int64 `json:"drain_timeout_ms,omitempty"`
}

// Stops the specified instance, but does not destroy it. All volatile state
// (e.g., RAM contents) is lost. Does nothing for an instance that is already
// stopped. The instance can be started again with the start endpoint.
//
// See: https://docs.kraft.cloud/002-rest-api-v1-instances.html#stop
func (i *InstanceClient) Stop(ctx context.Context, uuid string, drainTimeoutMS int64) (*Instance, error) {
	if uuid == "" {
		return nil, errors.New("UUID cannot be empty")
	}
	base := i.BaseURL + Endpoint
	endpoint := fmt.Sprintf("%s/%s/stop", base, uuid)

	requestBody := StopInstancePayload{
		DrainTimeoutMS: drainTimeoutMS,
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("marshalling request body: %w", err)
	}

	var response kraftcloud.ServiceResponse[Instance]
	if err := i.DoRequest(ctx, http.MethodPut, endpoint, bytes.NewBuffer(body), &response); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	return response.FirstOrErr()
}
