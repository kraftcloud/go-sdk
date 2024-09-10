// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2023, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package instances

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	ukcclient "sdk.kraft.cloud/client"
	"sdk.kraft.cloud/uuid"
)

// Log implements InstancesService.
func (c *client) Log(ctx context.Context, id string, offset int, limit int) (*ukcclient.ServiceResponse[LogResponseItem], error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("identifier cannot be empty")
	}

	reqItem := make(map[string]any, 3)
	if uuid.IsValid(id) {
		reqItem["uuid"] = id
	} else {
		reqItem["name"] = id
	}
	reqItem["offset"] = offset
	reqItem["limit"] = limit

	body, err := json.Marshal([]map[string]any{reqItem})
	if err != nil {
		return nil, fmt.Errorf("marshalling request body: %w", err)
	}

	resp := &ukcclient.ServiceResponse[LogResponseItem]{}
	if err := c.request.DoRequest(ctx, http.MethodGet, Endpoint+"/log", bytes.NewReader(body), resp); err != nil {
		return nil, fmt.Errorf("performing the request: %w", err)
	}

	return resp, nil
}

// TailLogs implements InstancesService.
func (c *client) TailLogs(ctx context.Context, id string, follow bool, tail int, delay time.Duration) (chan string, chan error, error) {
	var (
		logChan = make(chan string)
		errChan = make(chan error)
	)

	// Start a goroutine to concurrently fetch logs and send them over the logs
	// channel.
	go func() {
		var startOffset int

		// If tail is set greater than zero, start by fetching the logs in reverse
		// order, separating by newline and appending to a buffer. Once the buffer
		// size reaches the length of tail, print the lines by reversing the buffer
		// again.
		if tail > 0 {
			buf := make([]string, 0)
			var reverseOffset int

		poll:
			for {
				reverseOffset -= LogMaxPageSize

				// Always request the largest page size when iterating backwards through
				// the logs as we cannot determine the number of lines ahead of time and
				// this ultimately reduces the number of requests.
				resp, err := c.Log(ctx, id, reverseOffset, LogMaxPageSize)
				if err != nil {
					errChan <- err
					continue
				}

				item, err := resp.FirstOrErr()
				if err != nil {
					errChan <- err
					continue
				}

				output, err := base64.StdEncoding.DecodeString(item.Output)
				if err != nil {
					errChan <- err
					continue
				}

				// Iterate through each character of the `output` in reverse order and
				// take a note of the start ("limit"), `i`, and end ("offset"), `j`, of
				// a line by reading the newline `\n`.  When the newline characters
				// are read, it represents the end of the current line but also the
				// start of the next.
				//
				// For example, if we have the log output set to:
				//
				// ```
				// hello
				// world
				// ```
				//
				// The value of `output` would a byte array with `total` equal to 12
				// and would take the form:
				//
				//   ◀─────────────────────────────────────────────────────────┤ time
				//    0    1    2    3    4    5    6    7    8    9   10   11
				// ┌────┬────┬────┬────┬────┬────┬────┬────┬────┬────┬────┬────┐
				// │  h │  e │  l │  l │  o │ \n │  w │  o │  r │  l │  d │ \n │
				// └────┴────┴────┴────┴────┴────┴────┴────┴────┴────┴────┴────┘
				//    ▲                        ▲                            ▲
				//    │                        │                            │  (step 1)
				//    │                        └j=6                    i=0 ─┘
				//    │                        ▲
				//    └ j=0               i=6 ─┘ (i is now set to j)         (step 2)
				//

				total := len(output)

				// Set the start offset to the end of total available logs. Continuously
				// set this value over the loop, as the value may change by during
				// iteration (as new logs have come in).
				startOffset = item.Available.End

				var i int
				var hasNewlineOnFirstLine bool
				for j := total - 1; j >= 0; j-- {
					// Check whether a newline sequence is detected.
					if j > 0 && output[j] == '\n' {
						// Only save the line to the buffer if it is not the last character
						// of the `output`, since there would be nothing to add to the
						// buffer.
						if j+1 != total {
							if i > 0 {
								buf = append(buf, string(output[j+1:i]))
							}
						} else {
							hasNewlineOnFirstLine = true
						}

						// Update i to the end of the current line.
						i = j

						// Break early if we have enough lines saved to the buffer.
						if len(buf) == tail {
							break poll
						}
					}
				}

				if hasNewlineOnFirstLine {
					buf = append(buf, string(output[0:i]))
				}

				// Break if we have enough lines saved to the buffer or there is no more
				// logs.
				if total < LogMaxPageSize {
					break poll
				}
			}

			// Print the buffer in reverse, since lines were saved in reverse order.
			for i := len(buf) - 1; i >= 0; i-- {
				logChan <- buf[i]
			}
		}

		// If we've tailed, then we've reached the latest logs, and if following has
		// not been set, we can safely return early.
		// return early.
		if !follow && tail > 0 {
			close(logChan)
			close(errChan)
			return
		}

		// Set the page size to the maximum.  We are now moving forwards through the
		// logs and we can reduce the number of remote calls by first making larger
		// calls.
		pageSize := LogMaxPageSize

		// Now iterate through the logs, starting at the `startOffset`.
		for {
			resp, err := c.Log(ctx, id, startOffset, pageSize)
			if err != nil {
				errChan <- err
				continue
			}

			item, err := resp.FirstOrErr()
			if err != nil {
				errChan <- err
				continue
			}

			output, err := base64.StdEncoding.DecodeString(item.Output)
			if err != nil {
				errChan <- err
				continue
			}

			// Reduce the page size now that we've hit the latest logs.
			if len(output) < LogMaxPageSize {
				pageSize = LogDefaultPageSize
			}

			// Iterate through each character and take a note of the start ("offset"),
			// `i`, and end ("limit"), `j`, of a line by reading the newline `\n`
			// character.  When the newline character is read, send this offset and
			// limit via the logs channel.  Update the startOffset to the end of the
			// last line.
			if len(output) > 0 {
				var j int

				for i := 0; i < len(output); i++ {
					if i > 0 && output[i] == '\n' {
						logChan <- string(output[j:i])
						j = i + 1
					}

					// Upon reaching the end of the output, update the startOffset so that
					// the next page can be reached.
					if i == len(output)-1 {
						startOffset += j
					}
				}
			}

			// We've received the last payload of logs if the size of the logs is less
			// than the page size.  If tailing has been disabled, we can close the
			// channel now.
			if !follow && startOffset == item.Range.End {
				close(logChan)
				close(errChan)
				return
			}

			time.Sleep(delay)
		}
	}()

	return logChan, errChan, nil
}
