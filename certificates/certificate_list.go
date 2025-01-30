// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2024, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package certificates

import (
	"context"
	"fmt"
	"net/http"

	ukcclient "sdk.kraft.cloud/client"
)

// List implements CertificatesService.
func (c *client) List(ctx context.Context) (*ukcclient.ServiceResponse[GetResponseItem], error) {
	resp := &ukcclient.ServiceResponse[GetResponseItem]{}
	if err := c.request.DoRequest(ctx, http.MethodGet, Endpoint, nil, resp); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	return resp, nil
}
