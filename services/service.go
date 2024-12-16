// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package services

// Endpoint is the public path for the services service.
const Endpoint = "/services"

// Connection Handlers. UnikraftCloud uses connection handlers to decide how to
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
	// Terminate the TLS connection at the UnikraftCloud gateway using our wildcard
	// certificate issued for the kraft.cloud domain. The gateway forwards the
	// unencrypted traffic to your application.
	HandlerTLS Handler = "tls"

	// Enable HTTP mode on the load balancer to load balance on the level of
	// individual HTTP requests. In this mode, only HTTP connections are accepted.
	// If this option is not set the load balancer works in TCP mode and
	// distributes TCP connections.
	HandlerHTTP Handler = "http"

	// Redirect traffic from the source port to the destination port.
	HandlerRedirect Handler = "redirect"
)

// Handlers returns all possible service handlers.
func Handlers() []Handler {
	return []Handler{
		HandlerTLS,
		HandlerHTTP,
		HandlerRedirect,
	}
}
