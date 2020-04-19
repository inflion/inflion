// Copyright 2020 The Inflion Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package notification

import "strings"

type ruleMatcherContains struct {
	value string
}

func (c *ruleMatcherContains) match(value interface{}) bool {
	if v, ok := value.(string); ok {
		return strings.Contains(v, c.value)
	}
	return false
}

type ruleMatcherExact struct {
	value string
}

func (c *ruleMatcherExact) match(value interface{}) bool {
	if v, ok := value.(string); ok {
		return v == c.value
	}
	return false
}
