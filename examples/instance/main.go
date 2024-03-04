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

	kraftcloud "sdk.kraft.cloud"
	"sdk.kraft.cloud/instances"
	"sdk.kraft.cloud/services"
)

// Here, you'll learn how to create an instance and display its console output.
// Subsequent actions include stopping and starting the instance, listing all instances in the project, and, ultimately, deleting the created instance.
func main() {
	token := os.Getenv("KRAFTCLOUD_TOKEN")
	if token == "" {
		fmt.Println("Please set the KRAFTCLOUD_TOKEN environment variable")
		os.Exit(1)
	}

	client := kraftcloud.NewInstancesClient(
		kraftcloud.WithToken(token),
	)

	ctx := context.Background()

	instance, err := client.Create(ctx, instances.CreateInstanceRequest{
		Image:    "nginx:latest",
		MemoryMB: 32,
		ServiceGroup: &instances.CreateInstanceServiceGroupRequest{
			Services: []services.Service{{
				Port: 443,
				Handlers: []services.Handler{
					services.HandlerTLS,
					services.HandlerHTTP,
				},
				DestinationPort: 80,
			}},
		},
		Autostart: instances.DefaultAutoStart,
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if instance, err = client.GetByUUID(ctx, instance.UUID); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("[+] instance", instance.UUID, "was created")

	displayInstanceDetails(*instance)

	time.Sleep(time.Second * 1)

	// get and print the console logs
	logs, err := client.LogsByUUID(ctx, instance.UUID, -1, true)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("[+] instance logs:")
	fmt.Println(logs)

	// stop
	if instance, err = client.StopByUUID(ctx, instance.UUID, 0); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("[+] instance", instance.UUID, "was stopped")

	// delete
	err = client.DeleteByUUID(ctx, instance.UUID)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("[+] instance", instance.UUID, "was deleted")
}

// displayInstanceDetails pretty prints the result of an instance status call.
func displayInstanceDetails(instance instances.Instance) {
	fmt.Println("=====================================")
	fmt.Println("          Instance Details           ")
	fmt.Println("=====================================")

	fmt.Println("UUID         ", instance.UUID)
	fmt.Println("State        ", instance.State)
	fmt.Println("Created At   ", instance.CreatedAt)
	fmt.Println("Image        ", instance.Image)
	fmt.Println("Memory       ", instance.MemoryMB, "MB")
	fmt.Println("FQDN         ", instance.FQDN)
	fmt.Println("Private IP   ", instance.PrivateIP)
	fmt.Println("Boot Time    ", bootTimeToString(instance.BootTimeUS))

	fmt.Println("=====================================")
}

func bootTimeToString(bootTimeUS int64) string {
	bootTimeSec := float64(bootTimeUS) / 1_000_000.0
	if bootTimeSec < 1.0 {
		return fmt.Sprintf("%.2f ms", bootTimeSec*1_000)
	}
	return fmt.Sprintf("%.2f s", bootTimeSec)
}
