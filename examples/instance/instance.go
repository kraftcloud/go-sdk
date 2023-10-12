// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2022, Unikraft GmbH and The KraftKit Authors.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.
package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
	kraftcloud "sdk.kraft.cloud/v0"
	"sdk.kraft.cloud/v0/instance"
)

// Here, you'll learn how to create an instance and display its console output.
// Subsequent actions include stopping and starting the instance, listing all instances in the project, and, ultimately, deleting the created instance.
func main() {
	user := os.Getenv("KRAFTCLOUD_USER")
	token := os.Getenv("KRAFTCLOUD_TOKEN")

	if user == "" || token == "" {
		fmt.Println("Please set KRAFTCLOUD_USER and KRAFTCLOUD_TOKEN environment variables")
		return
	}

	apiClient := instance.NewInstancesClient(
		kraftcloud.WithUser(user),
		kraftcloud.WithToken(token),
	)
	ctx := context.Background()
	instance, err := apiClient.Create(ctx, instance.CreateInstanceRequest{
		// You have to build the kraft.cloud.yaml target from https://github.com/unikraft/app-nginx
		// and upload it with kraft pkg push to make this image available to your account.
		Image:    "unikraft.io/jayc.unikraft.io/nginx:latest",
		Args:     []string{"-c", "/nginx/conf/nginx.conf"},
		MemoryMB: 16,
		Services: []instance.CreateInstanceServicesRequest{
			{
				Port:         443,
				Handlers:     []string{instance.DefaultHandler},
				InternalPort: 80,
			},
		},
		Autostart: instance.DefaultAutoStart,
	})
	if err != nil {
		fmt.Printf("erred: %v\n", err)
		return
	}

	result, err := apiClient.Status(ctx, instance.UUID)
	if err != nil {
		fmt.Printf("erred: %v\n", err)
		return
	}

	DisplayInstanceDetails(*result)

	time.Sleep(time.Second * 1)

	// get and print the console output
	output, err := apiClient.Logs(ctx, instance.UUID, -1, true)
	if err != nil {
		fmt.Printf("erred: %v\n", err)
	}

	fmt.Println(output)

	// stop
	instance, err = apiClient.Stop(ctx, instance.UUID, 0)
	if err != nil {
		fmt.Printf("%v", err)
	}
	// start
	instance, err = apiClient.Start(ctx, instance.UUID, 0)
	if err != nil {
		fmt.Printf("%v", err)
	}
	// list
	instances, err := apiClient.List(ctx)
	if err != nil {
		fmt.Printf("%v", err)
	}
	// print the retrieved instances
	for _, i := range instances {
		fmt.Println(i.UUID)
	}
	// delete
	err = apiClient.Delete(ctx, instance.UUID)
	if err != nil {
		fmt.Printf("%v", err)
	}
}

// DisplayInstanceDetails pretty prints the result of an instance status call.
func DisplayInstanceDetails(instance instance.Instance) {
	headerColor := color.New(color.FgCyan, color.Bold)
	dataColor := color.New(color.FgWhite)

	headerColor.Println("=====================================")
	headerColor.Println("          Instance Details           ")
	headerColor.Println("=====================================")

	dataColor.Printf("UUID:         %s\n", instance.UUID)
	dataColor.Printf("Status:       %s\n", instance.Status)
	dataColor.Printf("Created At:   %s\n", instance.CreatedAt)
	dataColor.Printf("Image:        %s\n", instance.Image)
	dataColor.Printf("Memory (MB):  %d\n", instance.MemoryMB)
	dataColor.Printf("DNS:          %s\n", instance.DNS)
	dataColor.Printf("Private IP:   %s\n", instance.PrivateIP)
	dataColor.Printf("Boot Time:    %s\n", bootTimeToString(instance.BootTimeUS))

	headerColor.Println("=====================================")
}

func bootTimeToString(bootTimeUS int64) string {
	bootTimeSec := float64(bootTimeUS) / 1_000_000.0
	if bootTimeSec < 1.0 {
		return fmt.Sprintf("%.2f ms", bootTimeSec*1_000)
	}
	return fmt.Sprintf("%.2f s", bootTimeSec)
}
