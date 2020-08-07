package flow

type NullActionExecutor struct{}

func (n NullActionExecutor) Run(_ ExecutionContext) (ActionResult, error) {
	return ActionResult{}, nil
}
