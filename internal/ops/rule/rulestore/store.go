// Copyright 2020 The Inflion Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package rulestore

import (
	"github.com/google/uuid"
	"github.com/inflion/inflion/internal/ops/rule"
)

type Store interface {
	GetRules(projectId int64) ([]rule.Rule, error)
	Create(rule RuleJson) (uuid.UUID, error)
	Get(rule RuleJson) (RuleJson, error)
	Update(rule RuleJson) error
	Delete(rule RuleJson) error
}
