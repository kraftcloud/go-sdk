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

// Deletes the specified instance. After this call the UUID of the instance is
// no longer valid. If the instance is currently running it is force stopped.
//
// See: https://docs.kraft.cloud/002-rest-api-v1-instances.html#delete
func (i *instancesClient) Delete(ctx context.Context, uuid string) error {
	if uuid == "" {
		return errors.New("UUID cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/%s", Endpoint, uuid)

	if i.request == nil {
		i.request = kraftcloud.NewServiceRequestFromDefaultOptions(i.opts)
	}

	defer func() { i.request = nil }()

	var response kraftcloud.ServiceResponse[Instance]
	if err := i.request.DoRequest(ctx, http.MethodDelete, endpoint, nil, &response); err != nil {
		return fmt.Errorf("performing the request: %w", err)
	}

	return nil
}
