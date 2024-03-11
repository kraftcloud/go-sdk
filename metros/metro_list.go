// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2024, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package metros

import (
	"context"
	"net"
	"net/http"
	"sync"
	"time"
)

// testMetroAlive sends a request to https://api.<metro>.kraft.cloud/ and checks
// if a response is received is received and the time to dial the tcp connection.
func testMetroAlive(metro, ip string) time.Duration {
	url := "https://api." + metro + ".kraft.cloud/"

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

// List implements MetrosService.
func (c *client) List(ctx context.Context, status bool) ([]ListResponseItem, error) {
	items := []ListResponseItem{
		{
			Code:     "fra0",
			Location: "Frankfurt, DE",
			Proxy:    "fra0.kraft.host",
			Ipv4:     "145.40.93.137",
			Online:   true,
		},
		{
			Code:     "dal0",
			Location: "Dallas, TX",
			Proxy:    "dal0.kraft.host",
			Ipv4:     "147.28.196.53",
			Online:   true,
		},
		{
			Code:     "sin0",
			Location: "Singapore",
			Proxy:    "sin0.kraft.host",
			Ipv4:     "145.40.71.141",
			Online:   true,
		},
		{
			Code:     "was1",
			Location: "Washington, DC",
			Proxy:    "was1.kraft.host",
			Ipv4:     "3.211.205.241",
			Online:   true,
		},
	}

	if status {
		var wg sync.WaitGroup
		for i := range items {
			wg.Add(1)
			go func(i int) {
				items[i].Delay = testMetroAlive(items[i].Code, items[i].Ipv4)
				if items[i].Delay == 0 {
					items[i].Online = false
				}
				wg.Done()
			}(i)
		}
		wg.Wait()
	}

	return items, nil
}
