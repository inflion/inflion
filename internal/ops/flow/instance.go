package flow

import (
	"github.com/inflion/inflion/internal/ops/flow/context"
	"github.com/inflion/inflion/internal/paws"
	"log"
	"strings"
)

type InstanceActionExecutor struct {
	action Action
}

func (i InstanceActionExecutor) Run(ec context.ExecutionContext) (ActionResult, error) {
	log.Println("execute action: " + i.action.Type)
	log.Printf("action params: %+v", i.action.Params)

	actionType := i.action.Params["action"]

	a := paws.AwsAccount{
		AccountId:  ec.GetValueByPath(context.NewPath("config.account_id")).(string),
		RoleName:   ec.GetValueByPath(context.NewPath("config.assume_role")).(string),
		ExternalId: ec.GetValueByPath(context.NewPath("config.external_id")).(string),
	}
	p, err := paws.New(a, "ap-northeast-1")
	if err != nil {
		return ActionResult{}, err
	}

	instanceIds := paws.InstanceIds{}

	t := ec.GetValueByPath(context.NewPath(i.action.Params["targets"])).(string)

	for _, id := range strings.Split(t, ",") {
		instanceIds = append(instanceIds, &id)
	}

	var affectedInstances paws.InstanceIds

	if actionType == "stop" {
		affectedInstances, err = p.StopInstances(instanceIds)
	} else if actionType == "start" {
		affectedInstances, err = p.StartInstances(instanceIds)
	}

	if err != nil {
		return ActionResult{}, err
	}
	var affected []string
	for _, a := range affectedInstances {
		affected = append(affected, *a)
	}

	return ActionResult{
		Action: i.action,
		Outputs: map[string]string{
			"affected": strings.Join(affected, ","),
		},
		ExitStatus: true,
	}, nil
}
