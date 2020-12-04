package flow

import (
	"github.com/inflion/inflion/flow/context"
	"log"
)

type MatcherActionExecutor struct {
	action Action
}

func (m MatcherActionExecutor) Run(_ context.ExecutionContext) (ActionResult, error) {
	log.Println("execute action: " + m.action.Type)
	return ActionResult{
		Action: m.action,
		Outputs: map[string]string{
			"result": "false",
		},
		ExitStatus: false,
	}, nil
}
