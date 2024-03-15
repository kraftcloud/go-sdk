// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package volumes

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	kcclient "sdk.kraft.cloud/client"
)

// DetachByUUID implements VolumesService.
func (c *client) DetachByUUID(ctx context.Context, uuid string) (*kcclient.ServiceResponse[DetachResponseItem], error) {
	if uuid == "" {
		return nil, errors.New("UUID cannot be empty")
	}

	resp := &kcclient.ServiceResponse[DetachResponseItem]{}
	if err := c.request.DoRequest(ctx, http.MethodPut, Endpoint+"/"+uuid+"/detach", nil, resp); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	return resp, nil
}
