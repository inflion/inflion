// Copyright 2020 The Inflion Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package rule

import "strings"

type ContainsMatcher struct {
	Value string
}

func (c *ContainsMatcher) match(value interface{}) bool {
	if v, ok := value.(string); ok {
		return strings.Contains(v, c.Value)
	}
	return false
}

type exactMatcher struct {
	value string
}

func (c *exactMatcher) match(value interface{}) bool {
	if v, ok := value.(string); ok {
		return v == c.value
	}
	return false
}
