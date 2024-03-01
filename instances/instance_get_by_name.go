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
	"strings"

	"sdk.kraft.cloud/client"
)

// Get returns the current state and the configuration of an instance.
//
// See: https://docs.kraft.cloud/api/v1/instances/#status
func (c *instancesClient) GetByName(ctx context.Context, name string) (*Instance, error) {
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}

	body, err := json.Marshal([]map[string]interface{}{{"name": name}})
	if err != nil {
		return nil, fmt.Errorf("encoding JSON object: %w", err)
	}

	var response client.ServiceResponse[Instance]
	if err := c.request.DoRequest(ctx, http.MethodGet, Endpoint, bytes.NewBuffer(body), &response); err != nil {
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
