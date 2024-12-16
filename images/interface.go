// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package images

import (
	"context"

	ukcclient "sdk.kraft.cloud/client"
)

type ImagesService interface {
	ukcclient.ServiceClient[ImagesService]

	// Lists all existing images.
	//
	// See: https://docs.kraft.cloud/api/v1/images/#list-existing-images
	List(ctx context.Context) (*ukcclient.ServiceResponse[GetResponseItem], error)

	// Get an existing image. You can filter by digest, tag and based on
	// whether the image is public or not. The returned groups fulfill all
	// provided filter criteria. No particular value is assumed if a filter is not
	// part of the request.
	//
	// See: https://docs.kraft.cloud/api/v1/images/#list-existing-images
	Get(ctx context.Context, ids ...string) (*ukcclient.ServiceResponse[GetResponseItem], error)

	// Delete an image by its provided name.
	DeleteByName(ctx context.Context, name string) error

	// Quotas returns the storage quotas for the user.
	Quotas(ctx context.Context) (*QuotasResponseItem, error)
}
