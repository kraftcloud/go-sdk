// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package instance

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	kraftcloud "sdk.kraft.cloud"
)

// Status returns the current status and the configuration of an instance.
//
// See: https://docs.kraft.cloud/002-rest-api-v1-instances.html#status
func (c *instancesClient) Status(ctx context.Context, uuid string) (*Instance, error) {
	if uuid == "" {
		return nil, errors.New("UUID cannot be empty")
	}

	endpoint := Endpoint + "/" + uuid

	var response kraftcloud.ServiceResponse[Instance]
	if err := c.request.DoRequest(ctx, http.MethodGet, endpoint, nil, &response); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	instance, err := response.FirstOrErr()
	if instance != nil && instance.Message != "" {
		err = fmt.Errorf("%w: %s", err, instance.Message)
	}
	return instance, err
}
