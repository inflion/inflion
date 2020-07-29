// Copyright 2020 The Inflion Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package flow

import (
	"encoding/json"
	"fmt"
	"github.com/PagerDuty/go-pagerduty"
	"github.com/inflion/inflion/internal/ops/flow/action"
	"github.com/inflion/inflion/internal/ops/flow/configstore"
	"github.com/inflion/inflion/internal/paws"
	"log"
	"strings"
	"time"
)

type ActionResult struct {
	Action     Action
	Outputs    map[string]string
	ExitStatus bool
}

type ActionExecutor interface {
	Run(context ExecutionContext, action Action) (ActionResult, error)
}

type NullActionExecutor struct {
}

func (n NullActionExecutor) Run(_ ExecutionContext, _ Action) (ActionResult, error) {
	return ActionResult{}, nil
}

type MatcherActionExecutor struct {
}

func (m MatcherActionExecutor) Run(_ ExecutionContext, action Action) (ActionResult, error) {
	log.Println("execute action: " + action.Type)
	return ActionResult{
		Action: action,
		Outputs: map[string]string{
			"result": "false",
		},
		ExitStatus: false,
	}, nil
}

type ConfigActionExecutor struct {
}

func (i ConfigActionExecutor) Run(c ExecutionContext, action Action) (ActionResult, error) {
	log.Println("execute action: " + action.Type)
	log.Printf("action params: %+v", action.Params)

	cs := configstore.EtcdConfigStore{} // TODO move somewhere

	project := c.GetValueByPath(NewPath("system.project")).(string)

	configs, err := cs.List(configstore.ConfigListRequest{
		Project: project,
		Key:     action.Params["key-prefix"],
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
		Action:     action,
		Outputs:    outputs,
		ExitStatus: true,
	}, nil
}

type InstanceDataActionExecutor struct {
}

func (i InstanceDataActionExecutor) Run(ec ExecutionContext, action Action) (ActionResult, error) {
	log.Println("execute action: " + action.Type)
	log.Printf("action params: %+v", action.Params)

	a := paws.AwsAccount{
		AccountId:  ec.GetValueByPath(NewPath("config.account_id")).(string),
		RoleName:   ec.GetValueByPath(NewPath("config.assume_role")).(string),
		ExternalId: ec.GetValueByPath(NewPath("config.external_id")).(string),
	}
	p, err := paws.New(a, "ap-northeast-1")
	if err != nil {
		return ActionResult{}, err
	}

	fc := paws.FilterCondition{
		All:      false,
		TagName:  action.Params["tag"],
		TagValue: action.Params["tag-value"],
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
		Action: action,
		Outputs: map[string]string{
			"instance_ids": strings.Join(instanceIds, ","),
		},
		ExitStatus: true,
	}, nil
}

type InstanceActionExecutor struct {
}

func (i InstanceActionExecutor) Run(ec ExecutionContext, action Action) (ActionResult, error) {
	log.Println("execute action: " + action.Type)
	log.Printf("action params: %+v", action.Params)

	actionType := action.Params["action"]

	a := paws.AwsAccount{
		AccountId:  ec.GetValueByPath(NewPath("config.account_id")).(string),
		RoleName:   ec.GetValueByPath(NewPath("config.assume_role")).(string),
		ExternalId: ec.GetValueByPath(NewPath("config.external_id")).(string),
	}
	p, err := paws.New(a, "ap-northeast-1")
	if err != nil {
		return ActionResult{}, err
	}

	instanceIds := paws.InstanceIds{}

	t := ec.GetValueByPath(NewPath(action.Params["targets"])).(string)

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
		Action: action,
		Outputs: map[string]string{
			"affected": strings.Join(affected, ","),
		},
		ExitStatus: true,
	}, nil
}

type NotificationActionExecutor struct {
}

func (n NotificationActionExecutor) Run(ec ExecutionContext, a Action) (ActionResult, error) {
	log.Println("execute action: " + a.Type)
	event := ec.ExecutionFields.Fields["event"].Values
	rawEvent := ec.ExecutionFields.Fields["raw-event"].Values["json"]
	log.Printf("event: %+v", event)

	if a.Params["type"] == "aws" {
		var notifier action.AwsSlackNotifier
		if accountMapping, ok := a.Params["account_mapping"]; ok {
			log.Println(accountMapping)
			notifier = action.AwsSlackNotifier{
				AccountMapping: n.convertAccountMapping(accountMapping),
			}
		} else {
			notifier = action.AwsSlackNotifier{}
		}
		err := notifier.Notify(a.Params, event, rawEvent.(json.RawMessage))
		if err != nil {
			log.Printf("notify err, %+v", err)
			return ActionResult{
				Action: a,
				Outputs: map[string]string{
					"result":  "false",
					"message": fmt.Sprintf("%+v", err),
				},
				ExitStatus: false,
			}, nil
		}
		return ActionResult{
			Action: a,
			Outputs: map[string]string{
				"result": "true",
			},
			ExitStatus: true,
		}, nil
	} else {
		return ActionResult{
			Action: a,
			Outputs: map[string]string{
				"result":  "false",
				"message": "type not found",
			},
			ExitStatus: false,
		}, nil
	}
}

// convert account mapping.
// from: 1234:name,5678:name2
// to: map[string]string{1234: name, 5678: name2}
func (n NotificationActionExecutor) convertAccountMapping(from string) map[string]string {
	mapping := map[string]string{}
	for _, m := range strings.Split(from, ",") {
		log.Printf("%+v", m)
		tmp := strings.Split(m, ":")
		log.Printf("%+v", tmp)
		if len(tmp) > 1 {
			mapping[tmp[0]] = tmp[1]
		}
	}
	return mapping
}

type LoggingActionExecutor struct {
}

func (l LoggingActionExecutor) Run(e ExecutionContext, action Action) (ActionResult, error) {
	log.Printf("log: %s", action.Params)

	return ActionResult{
		Action: action,
		Outputs: map[string]string{
			"result": "false",
		},
		ExitStatus: true,
	}, nil
}

type ParamsActionExecutor struct {
}

func (p ParamsActionExecutor) Run(e ExecutionContext, a Action) (ActionResult, error) {
	log.Printf("log: %s", a.Params)

	outputs := map[string]string{}

	for k, param := range a.Params {
		outputs[k] = param
	}

	return ActionResult{
		Action:     a,
		Outputs:    outputs,
		ExitStatus: true,
	}, nil
}

type PagerDutyActionExecutor struct{}

func (p PagerDutyActionExecutor) Run(ec ExecutionContext, action Action) (ActionResult, error) {
	log.Println("execute action: " + action.Type)
	log.Printf("action params: %+v", action.Params)

	key, ok := action.Params["key"]
	if !ok {
		return ActionResult{
			Action: action,
			Outputs: map[string]string{
				"result":  "false",
				"message": "parameter \"key\" not found",
			},
			ExitStatus: false,
		}, nil
	}

	source := "unknown"
	s, ok := ec.ExecutionFields.Fields["event"].Values["source"]
	if ok {
		source = fmt.Sprintf("%v", s)
	}

	pagerdutyEvent := pagerduty.V2Event{
		RoutingKey: key,
		Action:     "trigger",
		DedupKey:   "",
		Client:     "inflion",
		Payload: &pagerduty.V2Payload{
			Summary:   "inflion event",
			Source:    source,
			Severity:  "critical",
			Timestamp: time.Now().Format(time.RFC3339),
			Details:   ec.ExecutionFields.Fields["event"],
		},
	}

	resp, err := pagerduty.ManageEvent(pagerdutyEvent)
	if err != nil {
		return ActionResult{
			Action: action,
			Outputs: map[string]string{
				"result":  "false",
				"message": fmt.Sprintf("%+v", err),
			},
			ExitStatus: false,
		}, nil
	}

	return ActionResult{
		Action: action,
		Outputs: map[string]string{
			"result":    "true",
			"status":    resp.Status,
			"dedup_key": resp.DedupKey,
			"message":   resp.Message,
			"errors":    strings.Join(resp.Errors, ","),
		},
		ExitStatus: true,
	}, nil
}
