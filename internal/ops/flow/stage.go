// Copyright 2020 The Inflion Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package flow

type Stages struct {
	Stages []NormalStage
}

type NormalStage struct {
	Id        string
	Name      string
	NextStage NextStage
	Actions   []Action
}

func (s Stages) get(index int) NormalStage {
	if index < len(s.Stages) {
		return s.Stages[index]
	}

	return NormalStage{}
}

func (s Stages) replaceById(id string, stage NormalStage) Stages {
	newStages := Stages{s.Stages}
	for i, elm := range s.Stages {
		if elm.Id == id {
			newStages.Stages[i] = stage
		} else {
			newStages.Stages[i] = elm
		}
	}
	return newStages
}

func (s Stages) getFirstStage() NormalStage {
	return s.get(0)
}

func (s Stages) getById(id string) (NormalStage, bool) {
	for _, s := range s.Stages {
		if s.Id == id {
			return s, true
		}
	}
	return NormalStage{}, false
}

func (s NormalStage) getId() string {
	return s.Id
}

func (s NormalStage) isEnd() bool {
	return s.Id == endStageId
}

func (s NormalStage) isEmpty() bool {
	return s.Id == ""
}

func (s NormalStage) isNormalStage() bool {
	return true
}

func (s NormalStage) isConditionStage() bool {
	return false
}

func (s NormalStage) getNextStage() NextStage {
	return s.NextStage
}

func (s NormalStage) wasNextStageResolved() bool {
	if s.NextStage.isEnd() {
		return true
	}
	return s.NextStage.Stage != nil
}
