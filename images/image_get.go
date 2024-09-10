// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package images

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	ukcclient "sdk.kraft.cloud/client"
)

// Get implements ImagesService.
func (c *client) Get(ctx context.Context, ids ...string) (*ukcclient.ServiceResponse[GetResponseItem], error) {
	if len(ids) == 0 {
		return nil, errors.New("requires at least one identifier")
	}

	var reqItems []map[string]string
	for _, id := range ids {
		if strings.Contains(id, "@") {
			reqItems = append(reqItems, map[string]string{"digest": id})
		} else {
			reqItems = append(reqItems, map[string]string{"tag": id})
		}
	}

	body, err := json.Marshal(reqItems)
	if err != nil {
		return nil, fmt.Errorf("encoding JSON object: %w", err)
	}

	resp := &ukcclient.ServiceResponse[GetResponseItem]{}
	if err := c.request.DoRequest(ctx, http.MethodGet, Endpoint+"/list", bytes.NewReader(body), resp); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	return resp, nil
}
