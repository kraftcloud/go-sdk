// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2024, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package certificates

import (
	"context"

	kcclient "sdk.kraft.cloud/client"
)

type CertificatesService interface {
	kcclient.ServiceClient[CertificatesService]

	// GetByUUIDs returns the current status and the properties of one or more
	// certificate(s) by their UUID(s).
	//
	// See: https://docs.kraft.cloud/api/v1/certificates/#getting-the-status-of-a-certificate
	GetByUUIDs(ctx context.Context, uuids ...string) (*kcclient.ServiceResponse[GetResponseItem], error)

	// GetByNames returns the current status and the properties of one or more
	// certificate(s) by their names(s).
	//
	// See: https://docs.kraft.cloud/api/v1/certificates/#getting-the-status-of-a-certificate
	GetByNames(ctx context.Context, names ...string) (*kcclient.ServiceResponse[GetResponseItem], error)

	// DeleteByUUIDs deletes one or more certificate(s) by their UUID(s).
	//
	// See: https://docs.kraft.cloud/api/v1/certificates/#deleting-a-certificate
	DeleteByUUIDs(ctx context.Context, uuids ...string) (*kcclient.ServiceResponse[DeleteResponseItem], error)

	// DeleteByNames deletes one or more certificate(s) by their UUID(s).
	//
	// See: https://docs.kraft.cloud/api/v1/certificates/#deleting-a-certificate
	DeleteByNames(ctx context.Context, names ...string) (*kcclient.ServiceResponse[DeleteResponseItem], error)

	// List all existing certificates.
	//
	// See: https://docs.kraft.cloud/api/v1/certificates/#list-existing-certificates
	List(ctx context.Context) (*kcclient.ServiceResponse[ListResponseItem], error)
}
