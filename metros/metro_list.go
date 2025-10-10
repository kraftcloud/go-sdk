// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2024, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package metros

import (
	"context"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"
)

// testMetroAlive sends a request to https://api.<metro>.unikraft.cloud/ and
// checks if a response is received is received and the time to dial the tcp
// connection.
func testMetroAlive(metro, ip string) time.Duration {
	var url string
	if metro != "" && strings.ContainsAny(metro[len(metro)-1:], "0123456789") {
		url = "https://api." + metro + ".unikraft.cloud/"
	} else {
		url = "https://api." + metro + ".unikraft.cloud/"
	}

	client := http.Client{
		Timeout: 3 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return 0
	} else {
		if resp != nil && resp.StatusCode == http.StatusUnauthorized {
			address := ip + ":443"

			start := time.Now()
			conn, err := net.DialTimeout("tcp", address, 3*time.Second)
			elapsed := time.Since(start)

			if err != nil {
				return 0
			} else {
				conn.Close()
				return elapsed
			}
		} else {
			return 0
		}
	}
}

// fillMetroIP looks up the IP address of the metro using the DNS name.
func fillMetroIP(metro string) string {
	var url string
	if metro != "" && strings.ContainsAny(metro[len(metro)-1:], "0123456789") {
		url = metro + ".kraft.host"
	} else {
		url = metro + ".unikraft.app"
	}

	ips, err := net.LookupIP(url)
	if err != nil {
		return ""
	}

	return ips[0].String()
}

// List implements MetrosService.
func (c *client) List(ctx context.Context, status bool) ([]ListResponseItem, error) {
	items := []ListResponseItem{
		{
			Code:     "fra",
			Location: "Frankfurt, DE",
			Proxy:    "fra.unikraft.app",
		},
		{
			Code:     "dal",
			Location: "Dallas, TX",
			Proxy:    "dal.unikraft.app",
		},
		{
			Code:     "sin",
			Location: "Singapore",
			Proxy:    "sin.unikraft.app",
		},
		{
			Code:     "sfo",
			Location: "San Francisco, CA",
			Proxy:    "sfo.unikraft.app",
		},
		{
			Code:     "was",
			Location: "Washington, DC",
			Proxy:    "was.unikraft.app",
		},
	}

	var wg sync.WaitGroup
	for i := range items {
		wg.Add(1)
		go func(i int) {
			items[i].Ipv4 = fillMetroIP(items[i].Code)
			items[i].Online = items[i].Ipv4 != ""

			if items[i].Online && status {
				items[i].Delay = testMetroAlive(items[i].Code, items[i].Ipv4)
				items[i].Online = items[i].Delay != 0
			}
			wg.Done()
		}(i)
	}
	wg.Wait()

	return items, nil
}
