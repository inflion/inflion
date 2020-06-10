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

type MockRecipeReader struct{}

func (MockRecipeReader) Read() (Recipe, error) {
	stageS1 := Stage{
		Id:   "stage1-id",
		Name: "stage1-name",
		NextStage: NextStage{
			Id: "cond1-id",
		},
		Actions: Actions{
			Actions: []Action{
				{
					Type:   "matcher",
					Params: nil,
				},
			},
		},
	}
	stageS2 := Stage{
		Id:   "stage2-id",
		Name: "stage2-name",
		NextStage: NextStage{
			Id: "stage3-id",
		},
		Actions: Actions{
			Actions: []Action{
				{
					Type: "instance",
					Params: map[string]string{
						"action":             "restart",
						"target_instance_id": "i-12345",
					},
				},
			},
		},
	}
	stageS3 := Stage{
		Id:   "stage3-id",
		Name: "stage3-name",
		NextStage: NextStage{
			Id: "__end__",
		},
		Actions: Actions{},
	}
	stageS4 := Stage{
		Id:   "stage4-id",
		Name: "stage4-name",
		NextStage: NextStage{
			Id: "__end__",
		},
		Actions: Actions{},
	}

	condition1 := Condition{
		Id: "cond1-id",
		Expressions: []Expression{
			{
				Input:     "last.result",
				Operation: "equals",
				Value:     "success",
			},
		},
		IfTrue: NextStage{
			Id: "stage2-id",
		},
		IfFalse: NextStage{
			Id: "stage3-id",
		},
	}

	return Recipe{
		Conditions: Conditions{
			Conditions: []Condition{
				condition1,
			},
		},
		Stages: Stages{
			Stages: []Stage{
				stageS1, stageS2, stageS3, stageS4,
			},
		},
	}, nil
}
