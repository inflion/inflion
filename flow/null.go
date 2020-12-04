package flow

import "github.com/inflion/inflion/flow/context"

type NullActionExecutor struct{}

func (n NullActionExecutor) Run(_ context.ExecutionContext) (ActionResult, error) {
	return ActionResult{}, nil
}
