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

type MockOpsFlow struct{}

func (MockOpsFlow) Read() (Flow, error) {
	stageS1 := NormalStage{
		Id:   "stage1-id",
		Name: "stage1-name",
		NextStage: NextStage{
			Id: "__end__",
		},
		Actions: Actions{
			{
				Type: "notification",
				Params: map[string]string{
					"webhook_url":       "https://hooks.slack.com/services/T0109SMD89H/B010P6MBSJX/xcVSmPiAHLzee0ATth2Gqz",
					"channel":           "notify_test",
					"notification_type": "log",
				},
			},
		},
	}

	stageS2 := NormalStage{
		Id:   "stage2-id",
		Name: "stage2-name",
		NextStage: NextStage{
			Id: "stage3-id",
		},
		Actions: Actions{
			{
				Type: "instance",
				Params: map[string]string{
					"action":             "restart",
					"target_instance_id": "i-12345",
				},
			},
		},
	}
	stageS3 := NormalStage{
		Id:   "stage3-id",
		Name: "stage3-name",
		NextStage: NextStage{
			Id: "__end__",
		},
		Actions: Actions{},
	}
	stageS4 := NormalStage{
		Id:   "stage4-id",
		Name: "stage4-name",
		NextStage: NextStage{
			Id: "__end__",
		},
		Actions: Actions{},
	}

	condition1 := ConditionStage{
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

	return Flow{
		Conditions: Conditions{
			Conditions: []ConditionStage{
				condition1,
			},
		},
		Stages: Stages{
			Stages: []NormalStage{
				stageS1, stageS2, stageS3, stageS4,
			},
		},
	}, nil
}
