// Copyright 2020 The Inflion Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package matcher

import (
	"github.com/inflion/inflion/flow/event"
)

type Conditions struct {
	Conditions []Condition
}

func (r *Conditions) match(event event.InflionEvent) bool {
	var matches []bool
	for _, rule := range r.Conditions {
		matches = append(matches, rule.match(event))
	}

	result := true
	for _, b := range matches {
		if b == false {
			result = false
			break
		}
	}
	return result
}
