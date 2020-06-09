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
	Stages []Stage
}

type Stage struct {
	Id        string
	Name      string
	NextStage NextStage
	Actions   Actions
}

func (s Stages) get(index int) Stage {
	if index < len(s.Stages) {
		return s.Stages[index]
	}

	return Stage{}
}

func (s Stages) replaceById(id string, stage Stage) Stages {
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

func (s Stages) getFirstStage() Stage {
	return s.get(0)
}

func (s Stages) getById(id string) (Stage, bool) {
	for _, s := range s.Stages {
		if s.Id == id {
			return s, true
		}
	}
	return Stage{}, false
}

func (s Stage) getId() string {
	return s.Id
}

func (s Stage) isEnd() bool {
	return s.Id == "__end__"
}

func (s Stage) isEmpty() bool {
	return s.Id == ""
}

func (s Stage) isStage() bool {
	return true
}

func (s Stage) isCondition() bool {
	return false
}

func (s Stage) getNextStage() NextStage {
	return s.NextStage
}

func (s Stage) getNextNode() Node {
	return s.NextStage.Node
}

func (s Stage) wasNextStageResolved() bool {
	if s.NextStage.isEnd() {
		return true
	}
	return s.NextStage.Node != nil
}

