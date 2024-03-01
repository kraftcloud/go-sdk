// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2022, Unikraft GmbH and The KraftKit Authors.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.
package images

const (
	// Endpoint is the public path for the images service.
	Endpoint = "/images"
)

// Image describes a KraftCloud image as returned by the API server.
type Image struct {
	// Image digest to uniquely identify this image.
	Digest string `json:"digest"`

	// Tags referencing this image. Can be used to create instances from this
	// image.
	Tags []string `json:"tags"`

	// Indicates if the image comes with an init ramdisk.
	Initrd bool `json:"initrd"`

	// Total size of the image on disk in bytes including the initrd, if any.
	SizeInBytes int64 `json:"size_in_bytes"`

	// Application arguments hardcoded into the image. Prepended to the arguments
	// of any instance running the image.
	Args string `json:"args"`

	// Unikraft kernel arguments hardcoded into the image. Prepended to the kernel
	// arguments set by KraftCloud
	KernelArgs string `json:"kernel_args"`

	// Message contains the error message either on `partial_success` or `error`.
	Message string `json:"message,omitempty"`
}
