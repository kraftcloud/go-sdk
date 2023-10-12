// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package instance

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"

	kraftcloud "sdk.kraft.cloud/v0"
)

// Logs returns the console output of the specified instance.
//
// See: https://docs.kraft.cloud/002-rest-api-v1-instances.html#console
func (i *instancesClient) Logs(ctx context.Context, uuid string, maxLines int, latest bool) (string, error) {
	endpoint := fmt.Sprintf("%s/%s/console", Endpoint, uuid)

	if i.request == nil {
		i.request = kraftcloud.NewServiceRequestFromDefaultOptions(i.opts)
	}

	defer func() { i.request = nil }()

	var resp kraftcloud.ServiceResponse[Instance]
	if err := i.request.DoRequest(ctx, http.MethodGet, endpoint, nil, resp); err != nil {
		return "", fmt.Errorf("performing the request: %w", err)
	}

	instance, err := resp.FirstOrErr()
	if err != nil {
		return "", err
	}

	output, err := base64.StdEncoding.DecodeString(instance.Output)
	if err != nil {
		return "", fmt.Errorf("decoding base64 console output: %w", err)
	}

	return string(output), nil
}
