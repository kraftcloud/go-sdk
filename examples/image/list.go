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
	token := os.Getenv("KRAFTCLOUD_TOKEN")
	if token == "" {
		fmt.Println("Please set KRAFTCLOUD_TOKEN environment variable")
		return
	}

	client := kraftcloud.NewImagesClient(
		kraftcloud.WithToken(token),
	)
	filter := make(map[string]interface{})
	images, err := client.List(context.Background(), filter)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	for _, i := range images {
		fmt.Println(i.Digest)
	}
}
