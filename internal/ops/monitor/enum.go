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

type EventType struct{ value string }

var CPUUtilization = EventType{"CPUUtilization"}
var OpenPortDetected = EventType{"OpenPortDetected"}

func (e EventType) String() string {
	if e.value == "" {
		return "未定義"
	}
	return e.value
}