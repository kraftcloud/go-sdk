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

// StartByUUID starts a previously stopped instance based on its UUID. Does
// nothing for an instance that is already running.
//
// See: https://docs.kraft.cloud/api/v1/instances/#start
func (c *instancesClient) StartByUUID(ctx context.Context, uuid string, waitTimeoutMS int) (*Instance, error) {
	if uuid == "" {
		return nil, errors.New("UUID cannot be empty")
	}

	body, err := json.Marshal(map[string]interface{}{
		"wait_timeout_ms": waitTimeoutMS,
	})
	if err != nil {
		return nil, fmt.Errorf("marshalling request body: %w", err)
	}

	var response client.ServiceResponse[Instance]
	if err := c.request.DoRequest(ctx, http.MethodPut, Endpoint+"/"+uuid+"/start", bytes.NewBuffer(body), &response); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	instance, err := response.FirstOrErr()
	if instance != nil && instance.Message != "" {
		err = errors.Join(err, fmt.Errorf(instance.Message))
	}

	return instance, err
}
