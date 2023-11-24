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

// StartByName starts a previously stopped instance based on its name. Does
// nothing for an instance that is already running.
//
// See: https://docs.kraft.cloud/002-rest-api-v1-instances.html#start
func (c *instancesClient) StartByName(ctx context.Context, name string, waitTimeoutMS int) (*Instance, error) {
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}

	body, err := json.Marshal([]map[string]interface{}{{
		"name":            name,
		"wait_timeout_ms": waitTimeoutMS,
	}})
	if err != nil {
		return nil, fmt.Errorf("marshalling request body: %w", err)
	}

	var response client.ServiceResponse[Instance]
	if err := c.request.DoRequest(ctx, http.MethodPut, Endpoint+"/start", bytes.NewBuffer(body), &response); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	instance, err := response.FirstOrErr()
	if instance != nil && instance.Message != "" {
		err = errors.Join(err, fmt.Errorf(instance.Message))
	}

	return instance, err
}
