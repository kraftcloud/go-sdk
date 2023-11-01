// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package users

import (
	"context"

	"sdk.kraft.cloud/client"
)

type UsersService interface {
	client.ServiceClient[UsersService]

	// Lists quota usage and limits of your user account. Limits are hard limits
	// that cannot be exceeded.
	Quotas(ctx context.Context) (*Quotas, error)
}
