// Copyright 2020 The Inflion Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package monitor

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

type MonitoringEvent struct {
	Project string                 `json:"project"`
	Body    map[string]interface{} `json:"body"`
	RawBody json.RawMessage
}

func (e *MonitoringEvent) HasAttribute(key string) bool {
	if _, ok := e.Body[key]; ok {
		return true
	} else {
		return false
	}
}

func (e *MonitoringEvent) GetValue(key string) (interface{}, bool) {
	if _, ok := e.Body[key]; ok {
		value, ok := e.Body[key]
		return value, ok
	} else {
		return nil, false
	}
}

func (e *MonitoringEvent) Hash(ignoreKeys []string) string {
	var buf string

	for k, v := range e.Body {
		for _, ignore := range ignoreKeys {
			if k == ignore {
				continue
			}
		}

		if val, ok := v.(string); ok {
			buf += val
		}
	}

	h := sha256.New()
	h.Write([]byte(buf))

	return fmt.Sprintf("%x", h.Sum(nil))
}
