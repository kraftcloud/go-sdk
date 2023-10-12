// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2022, Unikraft GmbH and The KraftKit Authors.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.
package image

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	kraftcloud "sdk.kraft.cloud/v0"
)

const (
	// Endpoint is the public path for the images service.
	Endpoint = "/images"
)

// ImageClient wraps the v1 Image client of Kraftcloud.
type ImageClient struct {
	kraftcloud.RESTClient
}

// NewDefaultClient creates a sensible, default Kraftcloud image API client.
func NewDefaultImageClient(user, token string) *ImageClient {
	return NewImageClient(kraftcloud.NewHTTPClient(), kraftcloud.BaseURL, user, token)
}

func NewImageClient(httpClient kraftcloud.HTTPClient, baseURL, user, token string) *ImageClient {
	return &ImageClient{
		RESTClient: kraftcloud.RESTClient{
			HTTPClient: httpClient,
			BaseURL:    baseURL,
			User:       user,
			Token:      token,
		},
	}
}

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

// ImageListResponse holds the list of images description, as returned by the API.
type ImageListResponse struct {
	Status string `json:"status"`
	Data   struct {
		Images []Image `json:"images"`
	} `json:"data"`
}

// ListImages fetches all images from the Kraftcloud API.
// see: https://docs.kraft.cloud/004-rest-api-v1-images.html#list
func (i *ImageClient) ListImages(ctx context.Context, filter map[string]interface{}) ([]Image, error) {
	body, err := json.Marshal(filter)
	if err != nil {
		return nil, fmt.Errorf("marshalling request body: %w", err)
	}

	endpoint := i.BaseURL + Endpoint + "/list"

	var response ImageListResponse

	if err := i.DoRequest(ctx, http.MethodGet, endpoint, bytes.NewBuffer(body), &response); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	return response.Data.Images, nil
}
