package flow

import (
	"github.com/inflion/inflion/flow/configstore"
	"github.com/inflion/inflion/flow/context"
	"log"
	"strings"
)

type ConfigActionExecutor struct {
	action Action
}

func (i ConfigActionExecutor) Run(c context.ExecutionContext) (ActionResult, error) {
	log.Println("execute action: " + i.action.Type)
	log.Printf("action params: %+v", i.action.Params)

	cs := configstore.EtcdConfigStore{} // TODO move somewhere

	project := c.GetFiledByPath("system.project")

	configs, err := cs.List(configstore.ConfigListRequest{
		Project: project,
		Key:     i.action.Params["key-prefix"],
	})
	if err != nil {
		return ActionResult{}, nil
	}

	outputs := map[string]string{}
	for _, c := range configs.Configs {
		log.Printf("%s = %s", c.Key, c.Value)
		key := strings.Replace(c.Key, "/"+project+"/config/aws/", "", -1)
		outputs[key] = c.Value
	}

	return ActionResult{
		Action:     i.action,
		Outputs:    outputs,
		ExitStatus: true,
	}, nil
}
