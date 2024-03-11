// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2024, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package metros

import (
	"time"

	kcclient "sdk.kraft.cloud/client"
)

// ListResponseItem is a data item from a response to a /metros/list request.
// https://docs.kraft.cloud/api/v1/metros/#list-existing-metros
type ListResponseItem struct {
	Code     string        `json:"code"`
	Delay    time.Duration `json:"delay"`
	Ipv4     string        `json:"ipv4"`
	Location string        `json:"location"`
	Online   bool          `json:"online"`
	Proxy    string        `json:"proxy"`

	kcclient.APIResponseCommon
}
