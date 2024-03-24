// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package volumes

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	kcclient "sdk.kraft.cloud/client"
)

// DeleteByUUIDs implements VolumesService.
func (c *client) DeleteByUUIDs(ctx context.Context, uuids ...string) (*kcclient.ServiceResponse[DeleteResponseItem], error) {
	if len(uuids) == 0 {
		return nil, errors.New("requires at least one uuid")
	}

	reqItems := make([]map[string]string, 0, len(uuids))
	for _, uuid := range uuids {
		reqItems = append(reqItems, map[string]string{"uuid": uuid})
	}

	body, err := json.Marshal(reqItems)
	if err != nil {
		return nil, fmt.Errorf("encoding JSON object: %w", err)
	}

	resp := &kcclient.ServiceResponse[DeleteResponseItem]{}
	if err := c.request.DoRequest(ctx, http.MethodDelete, Endpoint, bytes.NewReader(body), resp); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	return resp, nil
}
