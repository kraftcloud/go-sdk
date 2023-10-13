# KraftCloud Go SDK

This SDK is an early version of a Go-based client designed to interface with the [KraftCloud](https://kraft.cloud) API.

> 📖 **Documentation**
>
> For a comprehensive list of all API endpoints and detailed usage, refer to the [official KraftCloud documentation](https://docs.kraft.cloud/).

## Requirements

- Go 1.20 or higher.
- Valid KraftCloud authentication credentials. [Sign up for the beta!](https://kraft.cloud)

## Quick start

```go
package main

import (
	"fmt"
	"context"

	kraftcloud "sdk.kraft.cloud"
	services "sdk.kraft.cloud/services"
)

func main() {
	client := services.NewServicesClient(
		kraftcloud.WithToken("token"),
	)

	filter := make(map[string]interface{})
	images, err := client.Images().List(context.Background(), filter)
	if err != nil {
		fmt.Printf("failed: %v", err)
		return
	}

	for _, i := range images {
		fmt.Println(i.Digest)
	}
}

```

## Examples

For additional practical implementations, check out the [examples directory](/examples):

- [Image Listing](/examples/image/list.go)

This example lists all images in your project

- [Instance Management](/examples/instance/instance.go)

Here, you'll learn how to create an instance and display its console output. Subsequent actions include stopping and starting the instance, listing all instances in the project, and, ultimately, deleting the created instance.


## Implemented functionality

Here's a breakdown of the available requests:

### Instance:
1. [Create](https://docs.kraft.cloud/002-rest-api-v1-instances.html#create)
2. [Status](https://docs.kraft.cloud/002-rest-api-v1-instances.html#status)
3. [List](https://docs.kraft.cloud/002-rest-api-v1-instances.html#list)
4. [Stop](https://docs.kraft.cloud/002-rest-api-v1-instances.html#stop)
5. [Start](https://docs.kraft.cloud/002-rest-api-v1-instances.html#start)
6. [Delete](https://docs.kraft.cloud/002-rest-api-v1-instances.html#delete)
7. [Console output](https://docs.kraft.cloud/002-rest-api-v1-instances.html#console)

### Image:
1. [List](https://docs.kraft.cloud/004-rest-api-v1-images.html#list)

