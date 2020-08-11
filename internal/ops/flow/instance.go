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
		AccountId:  ec.GetFiledByPath("config.account_id"),
		RoleName:   ec.GetFiledByPath("config.assume_role"),
		ExternalId: ec.GetFiledByPath("config.external_id"),
	}
	p, err := paws.New(a, "ap-northeast-1")
	if err != nil {
		return ActionResult{}, err
	}

	instanceIds := paws.InstanceIds{}

	t := ec.GetFiledByPath("targets")

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
