// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package instances

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"sdk.kraft.cloud/client"
)

// StopByName the specified instance based on its name, but does not destroy it. All volatile state
// (e.g., RAM contents) is lost. Does nothing for an instance that is already
// stopped. The instance can be started again with the start endpoint.
//
// See: https://docs.kraft.cloud/002-rest-api-v1-instances.html#stop
func (c *instancesClient) StopByName(ctx context.Context, name string, drainTimeoutMS int64) (*Instance, error) {
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}

	body, err := json.Marshal([]map[string]interface{}{{
		"name":             name,
		"drain_timeout_ms": drainTimeoutMS,
	}})
	if err != nil {
		return nil, fmt.Errorf("marshalling request body: %w", err)
	}

	var response client.ServiceResponse[Instance]
	if err := c.request.DoRequest(ctx, http.MethodPut, Endpoint+"/stop", bytes.NewBuffer(body), &response); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	instance, err := response.FirstOrErr()
	if instance != nil && instance.Message != "" {
		err = fmt.Errorf("%w: %s", err, instance.Message)
	}
	return instance, err
}
