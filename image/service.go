// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package image

import (
	"context"

	"sdk.kraft.cloud/client"
)

type ImagesService interface {
	client.ServiceClient[ImagesService]

	// Lists all existing images. You can filter by digest, tag and based on
	// whether the image is public or not. The returned groups fulfill all
	// provided filter criteria. No particular value is assumed if a filter is not
	// part of the request.
	//
	// See: https://docs.kraft.cloud/004-rest-api-v1-images.html#list-existing-images
	List(ctx context.Context, filter map[string]interface{}) ([]Image, error)
}
