// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2022, Unikraft GmbH and The KraftKit Authors.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.
package util

import (
	"fmt"
	"strings"
)

// The API expects image names in the format of index.unikraft.io/jayc.unikraft.io/python
// and not the usual index.unikraft.io/jayc.unikraft.io:python
// We use this function to normalize the name.
// See: https://docs.kraft.cloud/002-rest-api-v1-instances.html#create
func NormalizeImageName(name string) (string, error) {
	if strings.Count(name, ":") > 1 {
		return "", fmt.Errorf("more than one semicolon is present in the image name: %s", name)
	}

	return strings.Replace(name, ":", "/", 1), nil
}
