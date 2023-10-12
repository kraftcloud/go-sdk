// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2022, Unikraft GmbH and The KraftKit Authors.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.
package image

import (
	"fmt"
)

const (
	// Endpoint is the public path for the images service.
	Endpoint = "/images"
)

// Image describes a Kraftcloud image as returned by the API server.
type Image struct {
	Digest      string   `json:"digest" pretty:"Digest"`
	Tags        []string `json:"tags" pretty:"Tags"`
	Public      bool     `json:"public" pretty:"Public"`
	Initrd      bool     `json:"initrd" pretty:"Initrd"`
	SizeInBytes int64    `json:"size_in_bytes" pretty:"Size (bytes)"`
	Args        string   `json:"args" pretty:"Args"`
	KernelArgs  string   `json:"kernel_args" pretty:"Kernel Args"`
}

func (i *Image) GetFieldByPrettyTag(tag string) string {
	switch tag {
	case "Digest":
		return i.Digest
	case "Tags":
		return fmt.Sprintf("%v", i.Tags)
	case "Public":
		if i.Public {
			return "true"
		}
		return "false"
	case "Initrd":
		if i.Public {
			return "true"
		}
		return "false"
	case "SizeInBytes":
		return fmt.Sprintf("%d", i.SizeInBytes)
	case "Args":
		return i.Args
	case "KernelArgs":
		return i.KernelArgs
	default:
		return ""
	}
}
