package flow

import (
	"github.com/inflion/inflion/internal/ops/flow/context"
	"github.com/inflion/inflion/internal/paws"
	"log"
	"strings"
)

type InstanceDataActionExecutor struct {
	action Action
}

func (i InstanceDataActionExecutor) Run(ec context.ExecutionContext) (ActionResult, error) {
	log.Println("execute action: " + i.action.Type)
	log.Printf("action params: %+v", i.action.Params)

	a := paws.AwsAccount{
		AccountId:  ec.GetFiledByPath("config.account_id"),
		RoleName:   ec.GetFiledByPath("config.assume_role"),
		ExternalId: ec.GetFiledByPath("config.external_id"),
	}
	p, err := paws.New(a, "ap-northeast-1")
	if err != nil {
		return ActionResult{}, err
	}

	fc := paws.FilterCondition{
		All:      false,
		TagName:  i.action.Params["tag"],
		TagValue: i.action.Params["tag-value"],
	}

	instances, err := p.GetInstances(fc)
	if err != nil {
		return ActionResult{}, err
	}

	var instanceIds []string

	for _, instance := range instances {
		instanceIds = append(instanceIds, instance.InstanceID)
	}

	return ActionResult{
		Action: i.action,
		Outputs: map[string]string{
			"instance_ids": strings.Join(instanceIds, ","),
		},
		ExitStatus: true,
	}, nil
}
