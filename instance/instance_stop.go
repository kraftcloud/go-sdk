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

	kraftcloud "sdk.kraft.cloud"
)

// InstanceStopRequest carries the data used by stop instance requests.
//
// See: https://docs.kraft.cloud/002-rest-api-v1-instances.html#request_3
type InstanceStopRequest struct {
	// Timeout for draining connections in milliseconds. The instance does not
	// receive new connections in the draining phase. The instance is stopped when
	// the last connection has been closed or the timeout expired.
	DrainTimeoutMS int64 `json:"drain_timeout_ms,omitempty"`
}

// Stops the specified instance, but does not destroy it. All volatile state
// (e.g., RAM contents) is lost. Does nothing for an instance that is already
// stopped. The instance can be started again with the start endpoint.
//
// See: https://docs.kraft.cloud/002-rest-api-v1-instances.html#stop
func (i *instancesClient) Stop(ctx context.Context, uuid string, drainTimeoutMS int64) (*Instance, error) {
	if uuid == "" {
		return nil, errors.New("UUID cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/%s/stop", Endpoint, uuid)

	requestBody := InstanceStopRequest{
		DrainTimeoutMS: drainTimeoutMS,
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("marshalling request body: %w", err)
	}

	if i.request == nil {
		i.request = kraftcloud.NewServiceRequestFromDefaultOptions(i.opts)
	}

	defer func() { i.request = nil }()

	var response kraftcloud.ServiceResponse[Instance]
	if err := i.request.DoRequest(ctx, http.MethodPut, endpoint, bytes.NewBuffer(body), &response); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	instance, err := response.FirstOrErr()
	if instance != nil && instance.Message != "" {
		err = fmt.Errorf("%w: %s", err, instance.Message)
	}
	return instance, err
}
