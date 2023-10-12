// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2022, Unikraft GmbH and The KraftKit Authors.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.
package instance

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	kraftcloud "sdk.kraft.cloud/v0"
	"sdk.kraft.cloud/v0/util"
)

const (
	instancesEndpoint = "/instances"
)

// InstanceClient is a basic wrapper around the v1 Instance client of Kraftcloud.
// see: https://docs.kraft.cloud/002-rest-api-v1-instances.html
type InstanceClient struct {
	kraftcloud.Client
}

// NewDefaultClient creates a sensible, default Kraftcloud instance API client.
func NewDefaultInstanceClient(user, token string) *InstanceClient {
	return NewInstanceClient(kraftcloud.NewHTTPClient(), kraftcloud.BaseURL, user, token)
}

func NewInstanceClient(httpClient kraftcloud.HTTPClient, baseURL, user, token string) *InstanceClient {
	return &InstanceClient{
		Client: kraftcloud.Client{
			HTTPClient: httpClient,
			BaseURL:    baseURL,
			User:       user,
			Token:      token,
		},
	}
}

// CreateInstancePayload holds all the data necessary to create an instance via the API.
// see: https://docs.kraft.cloud/002-rest-api-v1-instances.html#create
type CreateInstancePayload struct {
	// Name of the Unikraft image to instantiate. Private images will be available under your user's namespace
	Image string
	// Application arguments
	Args []string
	// Amount of memory to assign to the instance in megabytes
	Memory int64
	// Connection handlers. Must be [ "tls" ]
	Handlers []string
	// Public-facing Port
	Port int64
	// Port that the image listens on
	InternalPort int64
	// Autostart behavior. If true the instance will start immediately after creation
	Autostart bool
}

// StopInstancePayload carries the data used by stop instance requests.
type StopInstancePayload struct {
	DrainTimeoutMS int64 `json:"drain_timeout_ms,omitempty"`
}

// InstanceResponse holds instance description, as returned by the API.
type InstanceResponse struct {
	Status string `json:"status"`
	Data   struct {
		Instances []Instance `json:"instances"`
	} `json:"data"`
}

// NetworkInterface holds interface data returned by the Instance API.
type NetworkInterface struct {
	UUID      string `json:"uuid"`
	PrivateIP string `json:"private_ip"`
	MAC       string `json:"mac"`
}

// ConsoleOutputResponse holds console output data, as returned by the API.
type ConsoleOutputResponse struct {
	Status string `json:"status"`
	Data   struct {
		Instances []struct {
			Status string `json:"status"`
			UUID   string `json:"uuid"`
			Output string `json:"output"`
		} `json:"instances"`
	} `json:"data"`
}

// Instance holds the description of the Kraftcloud compute instance, as understood by the API server.
type Instance struct {
	UUID              string             `json:"uuid" pretty:"UUID"`
	DNS               string             `json:"dns" pretty:"DNS"`
	PrivateIP         string             `json:"private_ip" pretty:"PrivateIP"`
	Status            string             `json:"status" pretty:"Status"`
	CreatedAt         string             `json:"created_at" pretty:"Created At"`
	Image             string             `json:"image" pretty:"Image"`
	MemoryMB          int                `json:"memory_mb" pretty:"Memory (MB)"`
	Args              []string           `json:"args" pretty:"Args"`
	Env               map[string]string  `json:"env" pretty:"Env"`
	ServiceGroup      string             `json:"service_group" pretty:"Service Group"`
	NetworkInterfaces []NetworkInterface `json:"network_interfaces" pretty:"Network Interfaces"`
	BootTimeUS        int64              `json:"boot_time_us" pretty:"Boot Time (ms)"`
	Message           string             `json:"message"`
	Error             int64              `json:"error"`
}

func (i *Instance) GetFieldByPrettyTag(tag string) string {
	// TODO(jake-ciolek): Use reflection?
	switch tag {
	case "UUID":
		return i.UUID
	case "DNS":
		return i.DNS
	case "PrivateIP":
		return i.PrivateIP
	case "Status":
		return i.Status
	case "Created At":
		return i.CreatedAt
	case "Image":
		return i.Image
	case "Memory (MB)":
		return fmt.Sprintf("%d", i.MemoryMB)
	case "Args":
		return fmt.Sprintf("%v", i.Args)
	case "Env":
		return fmt.Sprintf("%v", i.Env)
	case "Service Group":
		return i.ServiceGroup
	case "Network Interfaces":
		return fmt.Sprintf("%v", i.NetworkInterfaces)
	case "Boot Time (ms)":
		return fmt.Sprintf("%d", i.BootTimeUS)
	default:
		return ""
	}
}

// CreateInstance dispatches the request to create a Kraftcloud compute instance.
// see: https://docs.kraft.cloud/002-rest-api-v1-instances.html#create
func (i *InstanceClient) CreateInstance(ctx context.Context, data CreateInstancePayload) (*Instance, error) {
	// normalize into the from kraftcloud API expects:
	image, err := util.NormalizeImageName(data.Image)
	if err != nil {
		return nil, fmt.Errorf("normalizing image name: %w", err)
	}

	requestBody := map[string]interface{}{
		"image":     image,
		"args":      data.Args,
		"memory_mb": data.Memory,
		"services": []map[string]interface{}{
			{
				"port":          data.Port,
				"internal_port": data.InternalPort,
				"handlers":      data.Handlers,
			},
		},
		"autostart": data.Autostart,
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("marshalling request body: %w", err)
	}

	endpoint := i.BaseURL + instancesEndpoint

	var response InstanceResponse
	if err := i.DoRequest(ctx, http.MethodPost, endpoint, bytes.NewBuffer(body), &response); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	return firstInstanceOrErr(&response)
}

// InstanceStatus queries the Kraftcloud compute API for details of a running instance.
// see: https://docs.kraft.cloud/002-rest-api-v1-instances.html#status
func (i *InstanceClient) InstanceStatus(ctx context.Context, uuid string) (*Instance, error) {
	if uuid == "" {
		return nil, errors.New("UUID cannot be empty")
	}
	base := i.BaseURL + instancesEndpoint
	endpoint := fmt.Sprintf("%s/%s", base, uuid)

	var response InstanceResponse

	if err := i.DoRequest(ctx, http.MethodGet, endpoint, nil, &response); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	return firstInstanceOrErr(&response)
}

// ListInstances fetches all instances from the Kraftcloud compute API.
// see: https://docs.kraft.cloud/002-rest-api-v1-instances.html#list
func (i *InstanceClient) ListInstances(ctx context.Context) ([]Instance, error) {
	base := i.BaseURL + instancesEndpoint
	endpoint := fmt.Sprintf("%s/list", base)

	var response InstanceResponse

	if err := i.DoRequest(ctx, http.MethodGet, endpoint, nil, &response); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	return response.Data.Instances, nil
}

// StopInstance stops the specified instance, but does not destroy it. All volatile state (e.g., RAM contents) is lost.
// Does nothing for an instance that is already stopped. The instance can be started again with the start endpoint.
// see https://docs.kraft.cloud/002-rest-api-v1-instances.html#stop
func (i *InstanceClient) StopInstance(ctx context.Context, uuid string, drainTimeoutMS int64) (*Instance, error) {
	if uuid == "" {
		return nil, errors.New("UUID cannot be empty")
	}
	base := i.BaseURL + instancesEndpoint
	endpoint := fmt.Sprintf("%s/%s/stop", base, uuid)

	requestBody := StopInstancePayload{
		DrainTimeoutMS: drainTimeoutMS,
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("marshalling request body: %w", err)
	}

	var response InstanceResponse
	if err := i.DoRequest(ctx, http.MethodPut, endpoint, bytes.NewBuffer(body), &response); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	return firstInstanceOrErr(&response)
}

// StartInstance starts a previously stopped instance. Does nothing for an instance that is already running.
// see: https://docs.kraft.cloud/002-rest-api-v1-instances.html#start
func (i *InstanceClient) StartInstance(ctx context.Context, uuid string, waitTimeoutMS int) (*Instance, error) {
	if uuid == "" {
		return nil, errors.New("UUID cannot be empty")
	}
	base := i.BaseURL + instancesEndpoint
	endpoint := fmt.Sprintf("%s/%s/start", base, uuid)

	requestBody := map[string]interface{}{
		"wait_timeout_ms": waitTimeoutMS,
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("marshalling request body: %w", err)
	}

	var response InstanceResponse

	if err := i.DoRequest(ctx, http.MethodPut, endpoint, bytes.NewBuffer(body), &response); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	return firstInstanceOrErr(&response)
}

// DeleteInstance deletes the specified instance. After this call the UUID of the instance is no longer valid. If the instance is currently running it is force stopped.
// see: https://docs.kraft.cloud/002-rest-api-v1-instances.html#delete
func (i *InstanceClient) DeleteInstance(ctx context.Context, uuid string) error {
	if uuid == "" {
		return errors.New("UUID cannot be empty")
	}
	base := i.BaseURL + instancesEndpoint
	endpoint := fmt.Sprintf("%s/%s", base, uuid)

	var response InstanceResponse

	if err := i.DoRequest(ctx, http.MethodDelete, endpoint, nil, &response); err != nil {
		return fmt.Errorf("performing the request: %w", err)
	}

	return nil
}

// ConsoleOutput fetches console output of the specified instance.
// see: https://docs.kraft.cloud/002-rest-api-v1-instances.html#console
func (i *InstanceClient) ConsoleOutput(ctx context.Context, uuid string, maxLines int, latest bool) (string, error) {
	base := i.BaseURL + instancesEndpoint
	endpoint := fmt.Sprintf("%s/%s/console", base, uuid)

	response := &ConsoleOutputResponse{}

	if err := i.DoRequest(ctx, http.MethodGet, endpoint, nil, response); err != nil {
		return "", fmt.Errorf("performing the request: %w", err)
	}

	if response.Data.Instances == nil {
		return "", errors.New("instances data is nil")
	}

	if len(response.Data.Instances) == 0 {
		return "", errors.New("no instances data returned from the server")
	}

	outputB64 := response.Data.Instances[0].Output
	output, err := base64.StdEncoding.DecodeString(outputB64)
	if err != nil {
		return "", fmt.Errorf("decoding base64 console output: %w", err)
	}

	return string(output), nil
}

func firstInstanceOrErr(response *InstanceResponse) (*Instance, error) {
	if response == nil {
		return nil, errors.New("response is nil")
	}
	if response.Data.Instances == nil {
		return nil, errors.New("instances data is nil")
	}
	if len(response.Data.Instances) == 0 {
		return nil, errors.New("no instances data returned from the server")
	}
	if response.Data.Instances[0].Status == "error" {
		return nil, errors.New(response.Data.Instances[0].Message)
	}
	return &response.Data.Instances[0], nil
}
