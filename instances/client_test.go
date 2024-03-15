// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package instances_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"testing"

	kraftcloud "sdk.kraft.cloud"
)

const (
	uuid1 = "00000000-0000-0000-0000-000000000001"
	uuid2 = "00000000-0000-0000-0000-000000000002"
	uuid3 = "00000000-0000-0000-0000-000000000003"
)

func TestClientThreadSafety(t *testing.T) {
	const requests = 100

	cli := kraftcloud.NewInstancesClient(
		kraftcloud.WithHTTPClient(httpMockStatusClient()),
	)

	errCh := make(chan error)
	t.Cleanup(func() { close(errCh) })

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	go func() {
		defer cancel()

		var wg sync.WaitGroup

		for _, uuid := range []string{uuid1, uuid2, uuid3} {
			wg.Add(1)
			go func(uuid string) {
				defer wg.Done()
				for i := 0; i < requests; i++ {
					resp, err := cli.GetByUUIDs(ctx, uuid)
					if err != nil {
						select {
						case <-ctx.Done():
							return
						default:
							errCh <- fmt.Errorf("[%s] Status: %v", uuid, err)
						}
						continue
					}

					ins, err := resp.FirstOrErr()
					if err != nil {
						select {
						case <-ctx.Done():
							return
						default:
							errCh <- fmt.Errorf("[%s] Status: %v", uuid, err)
						}
						continue
					}

					if ins.UUID != uuid {
						select {
						case <-ctx.Done():
							return
						default:
							errCh <- fmt.Errorf("[%s] Got unexpected uuid %s", uuid, ins.UUID)
						}
					}
				}
			}(uuid)
		}

		wg.Wait()
	}()

	select {
	case err := <-errCh:
		t.Fatal(err)
	case <-ctx.Done():
	}
}

// httpClientMockStatusRoundTripper returns a http.Client that returns static
// responses to requests towards the default GET endpoint of the instance API.
func httpMockStatusClient() *http.Client {
	return &http.Client{
		Transport: (*mockStatusRoundTripper)(nil),
	}
}

type mockStatusRoundTripper struct{}

var _ http.RoundTripper = (*mockStatusRoundTripper)(nil)

func (*mockStatusRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       instanceStatusRespBody(r),
	}
	return resp, nil
}

func instanceStatusRespBody(r *http.Request) io.ReadCloser {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	reqData := make([]map[string]string, 0, 1)
	if err := json.Unmarshal(b, &reqData); err != nil {
		fmt.Printf("%v: %s\n", err, b)
		return nil
	}
	if len(reqData) != 1 {
		fmt.Printf("Expected 1 item in request data, got %d\n", len(reqData))
		return nil
	}

	return io.NopCloser(strings.NewReader(`` +
		`{` +
		`  "data": {` +
		`    "instances": [` +
		`      {` +
		`        "uuid":"` + reqData[0]["uuid"] + `"` +
		`      }` +
		`    ]` +
		`  }` +
		`}`,
	))
}
