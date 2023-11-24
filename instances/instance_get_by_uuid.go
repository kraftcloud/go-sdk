// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package instances

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"sdk.kraft.cloud/client"
)

// Get returns the current state and the configuration of an instance.
//
// See: https://docs.kraft.cloud/002-rest-api-v1-instances.html#state
func (c *instancesClient) GetByUUID(ctx context.Context, uuid string) (*Instance, error) {
	if uuid == "" {
		return nil, errors.New("UUID cannot be empty")
	}

	endpoint := Endpoint + "/" + uuid

	var response client.ServiceResponse[Instance]
	if err := c.request.DoRequest(ctx, http.MethodGet, endpoint, nil, &response); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	instance, err := response.FirstOrErr()
	if instance != nil && instance.Message != "" {
		err = errors.Join(err, fmt.Errorf(instance.Message))
	}

	// Clean FQDN with trailing dot
	instance.FQDN = strings.TrimSuffix(instance.FQDN, ".")

	return instance, err
}
