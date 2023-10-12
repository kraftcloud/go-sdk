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
)

// Starts a previously stopped instance. Does nothing for an instance that is
// already running.
//
// See: https://docs.kraft.cloud/002-rest-api-v1-instances.html#start
func (i *InstanceClient) Start(ctx context.Context, uuid string, waitTimeoutMS int) (*Instance, error) {
	if uuid == "" {
		return nil, errors.New("UUID cannot be empty")
	}
	base := i.BaseURL + Endpoint
	endpoint := fmt.Sprintf("%s/%s/start", base, uuid)

	requestBody := map[string]interface{}{
		"wait_timeout_ms": waitTimeoutMS,
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("marshalling request body: %w", err)
	}

	var response InstanceResponse

	if err := i.DoRequest(ctx, http.MethodPut, endpoint, bytes.NewBuffer(body), &response); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	return firstInstanceOrErr(&response)
}
