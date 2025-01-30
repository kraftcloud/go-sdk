// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2024, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package metros

import (
	"context"

	ukcclient "sdk.kraft.cloud/client"
)

type MetrosService interface {
	ukcclient.ServiceClient[MetrosService]

	// Lists all existing metros. This list is currently hard-coded so the
	// returned result is always the same.
	//
	// See: https://docs.kraft.cloud/api/v1/metros/#list-existing-metros
	List(ctx context.Context, status bool) ([]ListResponseItem, error)
}
