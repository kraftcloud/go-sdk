// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2022, Unikraft GmbH and The KraftKit Authors.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.
package main

import (
	"context"
	"fmt"
	"os"

	kraftcloud "sdk.kraft.cloud"
)

// This demonstrates how to list images in your project.
func main() {
	token := os.Getenv("UNIKRAFTCLOUD_TOKEN")

	if token == "" {
		token = os.Getenv("KRAFTCLOUD_TOKEN")
	}

	if token == "" {
		token = os.Getenv("UKC_TOKEN")
	}

	if token == "" {
		token = os.Getenv("KC_TOKEN")
	}

	if token == "" {
		fmt.Println("Please set the UNIKRAFTCLOUD_TOKEN environment variable")
		os.Exit(1)
	}

	client := kraftcloud.NewImagesClient(
		kraftcloud.WithToken(token),
	)
	images, err := client.List(context.Background())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, i := range images {
		fmt.Println(i.Digest)
	}
}
