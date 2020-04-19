// Copyright 2020 The Inflion Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package paws

type Tags struct {
	Tags []Tag
}

func (t *Tags) append(tag Tag) {
	t.Tags = append(t.Tags, tag)
}

func (t *Tags) FindValue(tagName string) (string, bool) {
	for _, t := range t.Tags {
		if t.Key == tagName {
			return t.Value, true
		}
	}
	return "", false
}

func (t *Tags) FindValueOrElse(tagName string, defaultValue string) string {
	for _, t := range t.Tags {
		if t.Key == tagName {
			return t.Value
		}
	}
	return defaultValue
}

type Tag struct {
	Key   string
	Value string
}
