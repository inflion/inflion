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
	"github.com/google/uuid"
)

type Store interface {
	GetRules(project string) ([]Rule, error)
	Create(rule RuleJson) (uuid.UUID, error)
	Get(rule RuleJson) (RuleJson, error)
	Update(rule RuleJson) error
	Delete(rule RuleJson) error
}

func NewRuleStore() Store {
	return EtcdStore{}
}
