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

	"sdk.kraft.cloud/client"
)

// Deletes the specified instance. After this call the UUID of the instance is
// no longer valid. If the instance is currently running it is force stopped.
//
// See: https://docs.kraft.cloud/002-rest-api-v1-instances.html#delete
func (c *instancesClient) Delete(ctx context.Context, uuid string) error {
	if uuid == "" {
		return errors.New("UUID cannot be empty")
	}

	endpoint := Endpoint + "/" + uuid

	var response client.ServiceResponse[Instance]
	if err := c.request.DoRequest(ctx, http.MethodDelete, endpoint, nil, &response); err != nil {
		return fmt.Errorf("performing the request: %w", err)
	}

	return nil
}
