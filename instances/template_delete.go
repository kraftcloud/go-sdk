// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
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

	kcclient "sdk.kraft.cloud/client"
	"sdk.kraft.cloud/uuid"
)

// Delete implements InstancesService.
func (c *client) DeleteTemplate(ctx context.Context, ids ...string) (*kcclient.ServiceResponse[TemplateDeleteResponseItem], error) {
	if len(ids) == 0 {
		return nil, errors.New("requires at least one identifier")
	}

	reqItems := make([]map[string]string, 0, len(ids))
	for _, id := range ids {
		if uuid.IsValid(id) {
			reqItems = append(reqItems, map[string]string{"uuid": id})
		} else {
			reqItems = append(reqItems, map[string]string{"name": id})
		}
	}

	body, err := json.Marshal(reqItems)
	if err != nil {
		return nil, fmt.Errorf("encoding JSON object: %w", err)
	}

	resp := &kcclient.ServiceResponse[TemplateDeleteResponseItem]{}
	if err := c.request.DoRequest(ctx, http.MethodDelete, Endpoint+"/templates", bytes.NewReader(body), resp); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	return resp, nil
}
