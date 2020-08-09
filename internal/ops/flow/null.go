package flow

import "github.com/inflion/inflion/internal/ops/flow/context"

type NullActionExecutor struct{}

func (n NullActionExecutor) Run(_ context.ExecutionContext) (ActionResult, error) {
	return ActionResult{}, nil
}
