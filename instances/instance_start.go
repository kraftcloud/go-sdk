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

// Starts a previously stopped instance. Does nothing for an instance that is
// already running.
//
// See: https://docs.kraft.cloud/002-rest-api-v1-instances.html#start
func (c *instancesClient) Start(ctx context.Context, uuid string, waitTimeoutMS int) (*Instance, error) {
	if uuid == "" {
		return nil, errors.New("UUID cannot be empty")
	}

	endpoint := Endpoint + "/" + uuid + "/start"

	requestBody := map[string]interface{}{
		"wait_timeout_ms": waitTimeoutMS,
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("marshalling request body: %w", err)
	}

	var response client.ServiceResponse[Instance]
	if err := c.request.DoRequest(ctx, http.MethodPut, endpoint, bytes.NewBuffer(body), &response); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	instance, err := response.FirstOrErr()
	if instance != nil && instance.Message != "" {
		err = fmt.Errorf("%w: %s", err, instance.Message)
	}
	return instance, err
}
