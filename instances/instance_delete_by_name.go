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

// DeleteByName deletes the specified instance based on its name. After this
// call the UUID of the instance is no longer valid. If the instance is
// currently running it is force stopped.
//
// See: https://docs.kraft.cloud/002-rest-api-v1-instances.html#delete
func (c *instancesClient) DeleteByName(ctx context.Context, name string) error {
	if name == "" {
		return errors.New("name cannot be empty")
	}

	body, err := json.Marshal([]map[string]interface{}{{"name": name}})
	if err != nil {
		return fmt.Errorf("encoding JSON object: %w", err)
	}

	var response client.ServiceResponse[Instance]
	if err := c.request.DoRequest(ctx, http.MethodDelete, Endpoint, bytes.NewBuffer(body), &response); err != nil {
		return fmt.Errorf("performing the request: %w", err)
	}

	instance, err := response.FirstOrErr()
	if instance != nil && instance.Message != "" {
		err = fmt.Errorf("%w: %s", err, instance.Message)
	}
	return err
}
