// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package users

import (
	"context"

	kcclient "sdk.kraft.cloud/client"
)

type UsersService interface {
	kcclient.ServiceClient[UsersService]

	// Lists quota usage and limits of your user account. Limits are hard limits
	// that cannot be exceeded.
	// https://docs.kraft.cloud/api/v1/users/#list-quota-usage-and-limits
	Quotas(ctx context.Context) (*kcclient.ServiceResponse[QuotasResponseItem], error)
}
