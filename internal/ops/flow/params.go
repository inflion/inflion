package flow

import (
	"log"
)

type ParamsActionExecutor struct {
	action Action
}

func (p ParamsActionExecutor) Run(ExecutionContext) (ActionResult, error) {
	log.Printf("log: %s", p.action.Params)

	outputs := map[string]string{}

	for k, param := range p.action.Params {
		outputs[k] = param
	}

	return ActionResult{
		Action:     p.action,
		Outputs:    outputs,
		ExitStatus: true,
	}, nil
}
