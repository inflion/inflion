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

type Recipe struct {
	Conditions Conditions
	Stages     Stages
}

// At first, stage linking does not exist.
// The reason is that JSON can't hold links to an object.
func (r Recipe) LinkToNextStages() (recipe Recipe, err error) {
	_, recipe, err = r.link(r.Stages.getFirstStage(), r, 0)
	if err != nil {
		return recipe, err
	}

	//if recipe.hasUnresolvedStages(recipe) {
	//	return recipe, errors.New("stage resolving failed")
	//}

	return recipe, nil
}

func (r Recipe) link(node Node, recipe Recipe, depth int) (resolvedNode Node, resolvedRecipe Recipe, err error) {
	log.Printf("resolve: %s, depth: %d", node.getId(), depth)

	if node.isStage() {
		s := node.(Stage)
		if s.NextStage.isEnd() {
			return node, recipe, nil
		}

		nextNode, err := recipe.SearchNodeById(s.NextStage.Id)
		if err != nil {
			return node, recipe, err
		}

		if !nextNode.wasNextStageResolved() {
			log.Printf("next node id is: %s ", nextNode.getId())
			resolvedNode, resolvedRecipe, err = recipe.link(nextNode, recipe, depth + 1)
			if err != nil {
				return resolvedNode, resolvedRecipe, err
			}

			if resolvedNode == nil {
				return resolvedNode, resolvedRecipe, errors.New("resolved node is nil")
			} else {
				log.Printf("next node of stage(%s) was resolved: %s", s.Id, resolvedNode.getId())
			}

			// Assert never occur
			if nextNode.getId() != resolvedNode.getId() {
				log.Fatalf("ASSERTION FAILURE: %s, %s", nextNode.getId(), resolvedNode.getId())
			}

			log.Printf("resolved node id: %s", resolvedNode.getId())

			s.NextStage.Node = resolvedNode
		} else {
			log.Printf("next node already solved at %s", s.Id)
		}

		recipe.Stages = recipe.Stages.replaceById(s.Id, s)
		return s, recipe, nil
	} else if node.isCondition() {
		log.Println("resolve condition")
		c := node.(Condition)

		nextTrueNode, err := recipe.SearchNodeById(c.IfTrue.Id)
		if err != nil {
			return node, recipe, err
		}
		if !nextTrueNode.wasNextStageResolved() {
			log.Printf("cond(%s) ture node is unsolved, link to resolve node id(%s)", c.Id, c.IfTrue.Id)
			resolvedNode, resolvedRecipe, err = recipe.link(nextTrueNode, recipe, depth + 1)
			if err != nil {
				return resolvedNode, resolvedRecipe, err
			}
			c.IfTrue.Node = resolvedNode
		} else {
			c.IfTrue.Node = nextTrueNode
		}

		nextFalseNode, err := recipe.SearchNodeById(c.IfFalse.Id)
		if err != nil {
			return node, recipe, err
		}
		if !nextFalseNode.wasNextStageResolved() {
			log.Printf("cond(%s) false node is unsolved, link to resolve node id(%s)", c.Id, c.IfFalse.Id)
			resolvedNode, resolvedRecipe, err = r.link(nextFalseNode, recipe, depth + 1)
			if err != nil {
				return resolvedNode, resolvedRecipe, err
			}
			c.IfFalse.Node = resolvedNode
		} else {
			c.IfFalse.Node = nextFalseNode
		}

		recipe.Conditions = recipe.Conditions.ReplaceById(c.Id, c)
		log.Println("resolved condition node is " + c.Id)
		return c, recipe, nil
	} else {
		return Stage{}, Recipe{}, errors.New("node is neither stage nor condition")
	}
}

func (r Recipe) SearchNodeById(id string) (Node, error) {
	for _, c := range r.Conditions.Conditions {
		if c.Id == id {
			return c, nil
		}
	}

	for _, s := range r.Stages.Stages {
		if s.isEnd() {
			return Stage{
				Id:        "__end__",
				Name:      "end",
				NextStage: NextStage{},
				Actions:   Actions{},
			}, nil
		}

		if s.Id == id {
			return s, nil
		}
	}

	return nil, errors.New("Node ID: " + id + " does not exists")
}

func (r Recipe) hasUnresolvedStages(recipe Recipe) bool {
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
