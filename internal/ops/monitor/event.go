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
	"fmt"
)

const MessageKey = "Message"

type MonitoringEvent struct {
	Type      string
	ProjectId int64
	Message   string
	Values    map[string]interface{} `json:"values"`
}

func (e *MonitoringEvent) HasAttribute(key string) bool {
	if key == MessageKey {
		return true
	} else {
		_, ok := e.Values[key]
		return ok
	}
}

func (e *MonitoringEvent) GetValue(key string) (interface{}, bool) {
	if key == MessageKey {
		return e.Message, true
	} else {
		value, ok := e.Values[key]
		return value, ok
	}
}

func (e *MonitoringEvent) Hash(ignoreKeys []string) string {
	var buf string

	for k, v := range e.Values {
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
