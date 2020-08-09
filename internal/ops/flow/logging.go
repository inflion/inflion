package flow

import (
	"github.com/inflion/inflion/internal/ops/flow/context"
	"log"
)

type LoggingActionExecutor struct {
	action Action
}

func (l LoggingActionExecutor) Run(_ context.ExecutionContext) (ActionResult, error) {
	log.Printf("log: %s", l.action.Params)

	return ActionResult{
		Action: l.action,
		Outputs: map[string]string{
			"result": "false",
		},
		ExitStatus: true,
	}, nil
}
