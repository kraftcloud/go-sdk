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

// Image describes a KraftCloud image as returned by the API server.
type Image struct {
	// Image digest to uniquely identify this image.
	Digest string `json:"digest" pretty:"Digest"`

	// Tags referencing this image. Can be used to create instances from this
	// image.
	Tags []string `json:"tags" pretty:"Tags"`

	// Indicates if this is a public image. If true every KraftCloud user can
	// access it
	Public bool `json:"public" pretty:"Public"`

	// Indicates if the image comes with an init ramdisk.
	Initrd bool `json:"initrd" pretty:"Initrd"`

	// Total size of the image on disk in bytes including the initrd, if any.
	SizeInBytes int64 `json:"size_in_bytes" pretty:"Size (bytes)"`

	// Application arguments hardcoded into the image. Prepended to the arguments
	// of any instance running the image.
	Args string `json:"args" pretty:"Args"`

	// Unikraft kernel arguments hardcoded into the image. Prepended to the kernel
	// arguments set by KraftCloud
	KernelArgs string `json:"kernel_args" pretty:"Kernel Args"`
}

func (c *Image) GetFieldByPrettyTag(tag string) string {
	switch tag {
	case "Digest":
		return c.Digest
	case "Tags":
		return fmt.Sprintf("%v", c.Tags)
	case "Public":
		if c.Public {
			return "true"
		}
		return "false"
	case "Initrd":
		if c.Public {
			return "true"
		}
		return "false"
	case "SizeInBytes":
		return fmt.Sprintf("%d", c.SizeInBytes)
	case "Args":
		return c.Args
	case "KernelArgs":
		return c.KernelArgs
	default:
		return ""
	}
}
