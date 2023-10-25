// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package instances

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"

	"sdk.kraft.cloud/client"
)

// Logs returns the console output of the specified instance.
//
// See: https://docs.kraft.cloud/002-rest-api-v1-instances.html#console
func (c *instancesClient) Logs(ctx context.Context, uuid string, maxLines int, latest bool) (string, error) {
	endpoint := Endpoint + "/" + uuid + "/console"

	var resp client.ServiceResponse[Instance]
	if err := c.request.DoRequest(ctx, http.MethodGet, endpoint, nil, &resp); err != nil {
		return "", fmt.Errorf("performing the request: %w", err)
	}

	instance, err := resp.FirstOrErr()
	if instance != nil && instance.Message != "" {
		return "", fmt.Errorf("%w: %s", err, instance.Message)
	}

	output, err := base64.StdEncoding.DecodeString(instance.Output)
	if err != nil {
		return "", fmt.Errorf("decoding base64 console output: %w", err)
	}

	return string(output), nil
}
