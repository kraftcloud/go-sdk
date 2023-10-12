// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2022, Unikraft GmbH and The KraftKit Authors.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.
package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	kraftcloud "sdk.kraft.cloud/v0"
	"sdk.kraft.cloud/v0/image"
)

// This demonstrates how to list images in your project.
func main() {
	user := os.Getenv("KRAFTCLOUD_USER")
	token := os.Getenv("KRAFTCLOUD_TOKEN")

	if user == "" || token == "" {
		fmt.Println("Please set KRAFTCLOUD_USER and KRAFTCLOUD_TOKEN environment variables")
		return
	}

	client := image.NewImagesClient(
		kraftcloud.WithUser(user),
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

	printImages(images, []string{"Tags", "Digest", "SizeInBytes"})
}

func printImages(images []image.Image, fields []string) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	// print headers
	fmt.Fprintln(w, strings.Join(fields, "\t"))

	for _, image := range images {
		values := []string{}
		for _, field := range fields {
			values = append(values, image.GetFieldByPrettyTag(field))
		}
		fmt.Fprintln(w, strings.Join(values, "\t"))
	}

	w.Flush()
}
