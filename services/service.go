// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package services

const (
	// Endpoint is the public path for the services service.
	Endpoint = "/services"
)

// Connection Handlers. KraftCloud uses connection handlers to decide how to
// forward connections from the Internet to your application. You configure the
// handlers for every published service port individually.
//
// Currently, there is a set of constraints when publishing ports:
//   - Port 80: Must have http and must not have tls set;
//   - Port 443: Must have http and tls set;
//   - The `redirect` handler can only be set on port 80 (HTTP) to redirect to
//     port 443 (HTTPS);
//   - All other ports must have tls and must not have http set.
type Handler string

const (
	// Terminate the TLS connection at the KraftCloud gateway using our wildcard
	// certificate issued for the kraft.cloud domain. The gateway forwards the
	// unencrypted traffic to your application.
	HandlerTLS = Handler("tls")

	// Enable HTTP mode on the load balancer to load balance on the level of
	// individual HTTP requests. In this mode, only HTTP connections are accepted.
	// If this option is not set the load balancer works in TCP mode and
	// distributes TCP connections.
	HandlerHTTP = Handler("http")

	// Redirect traffic from the source port to the destination port.
	HandlerRedirect = Handler("redirect")
)

// Handlers returns all possible service handlers.
func Handlers() []Handler {
	return []Handler{
		HandlerTLS,
		HandlerHTTP,
		HandlerRedirect,
	}
}

// A service helps describe the the load balancing and autoscale groups.
type Service struct {
	// Public-facing port.
	Port int `json:"port"`

	// Application port to which inbound traffic is redirected.
	DestinationPort int `json:"destination_port,omitempty"`

	// List of handlers.
	Handlers []Handler `json:"handlers,omitempty"`
}

// A service group has a public DNS name such as
// young-monkey-uq6dxq0u.fra0.kraft.cloud. When you assign an instance to a
// service group, the instance becomes accessible from the Internet using this
// DNS name. KraftCloud generates a random DNS name of that form for every
// service group.
//
// Except for unencrypted HTTP traffic on port 80, KraftCloud accepts only TLS
// connections from the Internet. It uses Server Name Indication (SNI) to
// forward inbound traffic to your application.
//
// By default, a service group does not publish any services. To allow traffic
// to pass to the instances in the service group, you specify the network ports
// to publish. For example, if you run a web server you would publish port 80
// (HTTP) and/or port 443 (HTTPS).
type ServiceGroup struct {
	// The state of the service group.
	State string `json:"state,omitempty"`

	// Name of the service group.
	Name string `json:"name,omitempty"`

	// UUID of the group.
	UUID string `json:"uuid,omitempty"`

	// Public FQDN name under which the group is accessible from the Internet.
	FQDN string `json:"fqdn,omitempty"`

	// Instances contains a list of UUID representing instances attached to this
	// group.
	Instances []string `json:"instances,omitempty"`

	// Services contains the descriptions of exposed network services.
	Services []Service `json:"services,omitempty"`

	// Persistent indicates if the group will stay alive even after the last
	// instance detached.
	Persistent bool `json:"persistent,omitempty"`

	// Date and time of creation in ISO8601.
	CreatedAt string `json:"created_at,omitempty"`
}
