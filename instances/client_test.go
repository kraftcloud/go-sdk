// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package instances_test

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"path"
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
					ins, err := cli.Get(ctx, uuid)
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
// responses to requests towards the /status endpoint of the instance API.
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
	return io.NopCloser(strings.NewReader(`` +
		`{` +
		`  "data": {` +
		`    "instances": [` +
		`      {` +
		`        "uuid":"` + path.Base(r.URL.Path) + `"` +
		`      }` +
		`    ]` +
		`  }` +
		`}`,
	))
}
