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

import (
	"errors"
	"log"
)

type NextStage struct {
	Id    string
	Stage Stage
}

const endStageId = "__end__"

func (n NextStage) isEnd() bool {
	return n.Id == endStageId
}

type Stage interface {
	getId() string
	isEnd() bool
	isNormalStage() bool
	isConditionStage() bool
	wasNextStageResolved() bool
}

type Flow struct {
	Conditions Conditions
	Stages     Stages
}

type Reader interface {
	Read() (Flow, error)
}

// At first, stage linking does not exist.
// The reason is that JSON can't hold links to an object.
func (r Flow) LinkToNextStages() (recipe Flow, err error) {
	_, recipe, err = r.link(r.Stages.getFirstStage(), r, 0)
	if err != nil {
		return recipe, err
	}

	return recipe, nil
}

func (r Flow) link(stage Stage, recipe Flow, depth int) (resolvedStage Stage, resolvedFlow Flow, err error) {
	log.Printf("resolve: %s, depth: %d", stage.getId(), depth)

	if stage.isNormalStage() {
		s := stage.(NormalStage)
		if s.NextStage.isEnd() {
			return stage, recipe, nil
		}

		nextStage, err := recipe.SearchStageById(s.NextStage.Id)
		if err != nil {
			return stage, recipe, err
		}

		if !nextStage.wasNextStageResolved() {
			log.Printf("next stage id is: %s ", nextStage.getId())
			resolvedStage, resolvedFlow, err = recipe.link(nextStage, recipe, depth+1)
			if err != nil {
				return resolvedStage, resolvedFlow, err
			}

			if resolvedStage == nil {
				return resolvedStage, resolvedFlow, errors.New("resolved stage is nil")
			} else {
				log.Printf("next stage of stage(%s) was resolved: %s", s.Id, resolvedStage.getId())
			}

			// Assert never occur
			if nextStage.getId() != resolvedStage.getId() {
				log.Fatalf("ASSERTION FAILURE: %s, %s", nextStage.getId(), resolvedStage.getId())
			}

			log.Printf("resolved stage id: %s", resolvedStage.getId())

			s.NextStage.Stage = resolvedStage
		} else {
			log.Printf("next stage already solved at %s", s.Id)
		}

		recipe.Stages = recipe.Stages.replaceById(s.Id, s)
		return s, recipe, nil
	} else if stage.isConditionStage() {
		log.Println("resolve condition")
		c := stage.(ConditionStage)

		nextTrueStage, err := recipe.SearchStageById(c.IfTrue.Id)
		if err != nil {
			return stage, recipe, err
		}
		if !nextTrueStage.wasNextStageResolved() {
			log.Printf("cond(%s) ture stage is unsolved, link to resolve stage id(%s)", c.Id, c.IfTrue.Id)
			resolvedStage, resolvedFlow, err = recipe.link(nextTrueStage, recipe, depth+1)
			if err != nil {
				return resolvedStage, resolvedFlow, err
			}
			c.IfTrue.Stage = resolvedStage
		} else {
			c.IfTrue.Stage = nextTrueStage
		}

		nextFalseStage, err := recipe.SearchStageById(c.IfFalse.Id)
		if err != nil {
			return stage, recipe, err
		}
		if !nextFalseStage.wasNextStageResolved() {
			log.Printf("cond(%s) false stage is unsolved, link to resolve stage id(%s)", c.Id, c.IfFalse.Id)
			resolvedStage, resolvedFlow, err = r.link(nextFalseStage, recipe, depth+1)
			if err != nil {
				return resolvedStage, resolvedFlow, err
			}
			c.IfFalse.Stage = resolvedStage
		} else {
			c.IfFalse.Stage = nextFalseStage
		}

		recipe.Conditions = recipe.Conditions.ReplaceById(c.Id, c)
		log.Println("resolved condition stage is " + c.Id)
		return c, recipe, nil
	} else {
		return NormalStage{}, Flow{}, errors.New("stage is neither stage nor condition")
	}
}

func (r Flow) SearchStageById(id string) (Stage, error) {
	for _, c := range r.Conditions.Conditions {
		if c.Id == id {
			return c, nil
		}
	}

	for _, s := range r.Stages.Stages {
		if s.isEnd() {
			return NormalStage{
				Id:        endStageId,
				Name:      "end",
				NextStage: NextStage{},
				Actions:   Actions{},
			}, nil
		}

		if s.Id == id {
			return s, nil
		}
	}

	return nil, errors.New("Stage ID: " + id + " does not exists")
}

func (r Flow) hasUnresolvedStages(recipe Flow) bool {
	for _, s := range recipe.Stages.Stages {
		if !s.wasNextStageResolved() {
			log.Printf("unresolved stage id: %s", s.Id)
			return true
		}
	}
	for _, c := range recipe.Conditions.Conditions {
		if !c.wasNextStageResolved() {
			log.Printf("unresolved condition id: %s", c.Id)
			return true
		}
	}
	return false
}
