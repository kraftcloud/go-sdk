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

	"sdk.kraft.cloud/client"
)

// Deletes the specified instance. After this call the UUID of the instance is
// no longer valid. If the instance is currently running it is force stopped.
//
// See: https://docs.kraft.cloud/002-rest-api-v1-instances.html#delete
func (c *instancesClient) Delete(ctx context.Context, uuidOrName string) error {
	if uuidOrName == "" {
		return errors.New("UUID or Name cannot be empty")
	}

	endpoint := Endpoint + "/" + uuidOrName

	var response client.ServiceResponse[Instance]
	if err := c.request.DoRequest(ctx, http.MethodDelete, endpoint, nil, &response); err != nil {
		return fmt.Errorf("performing the request: %w", err)
	}

	return nil
}
