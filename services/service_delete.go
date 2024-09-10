// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package services

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	ukcclient "sdk.kraft.cloud/client"
	"sdk.kraft.cloud/uuid"
)

// Delete implements ServicesService.
func (c *client) Delete(ctx context.Context, ids ...string) (*ukcclient.ServiceResponse[DeleteResponseItem], error) {
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

	resp := &ukcclient.ServiceResponse[DeleteResponseItem]{}
	if err := c.request.DoRequest(ctx, http.MethodDelete, Endpoint, bytes.NewReader(body), resp); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	return resp, nil
}
