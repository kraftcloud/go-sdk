// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2022, Unikraft GmbH and The KraftKit Authors.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package instances

// Endpoint is the public path for the instances service.
const Endpoint = "/instances"

// Feature is a special feature of an instance.
type Feature string

// FeatureScaleToZero indicates that the instance can be scaled to zero.
const FeatureScaleToZero Feature = "scale-to-zero"
