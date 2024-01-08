// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package instances

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"

	"sdk.kraft.cloud/client"
)

// LogsByUUID returns the console output of the specified instance based on its
// UUID.
//
// See: https://docs.kraft.cloud/002-rest-api-v1-instances.html#console
func (c *instancesClient) LogsByUUID(ctx context.Context, uuid string, maxLines int, latest bool) (string, error) {
	if len(uuid) == 0 {
		return "", fmt.Errorf("UUID cannot be empty")
	}

	var resp client.ServiceResponse[Instance]
	if err := c.request.DoRequest(ctx, http.MethodGet, Endpoint+"/"+uuid+"/console", nil, &resp); err != nil {
		return "", fmt.Errorf("performing the request: %w", err)
	}

	instance, err := resp.FirstOrErr()
	if instance != nil && instance.Message != "" {
		return "", errors.Join(err, fmt.Errorf(instance.Message))
	}

	output, err := base64.StdEncoding.DecodeString(instance.Output)
	if err != nil {
		return "", fmt.Errorf("decoding base64 console output: %w", err)
	}

	return string(output), nil
}
