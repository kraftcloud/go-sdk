// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package volumes

import (
	"context"
	"fmt"
	"net/http"

	kcclient "sdk.kraft.cloud/client"
)

// List implements VolumesService.
func (c *client) ListTemplate(ctx context.Context) (*kcclient.ServiceResponse[TemplateGetResponseItem], error) {
	resp := &kcclient.ServiceResponse[TemplateGetResponseItem]{}
	if err := c.request.DoRequest(ctx, http.MethodGet, Endpoint+"/templates", nil, resp); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	return resp, nil
}
