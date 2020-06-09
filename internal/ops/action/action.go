// Copyright 2020 The Inflion Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package action

import (
	"encoding/json"
	"github.com/inflion/inflion/internal/ops/monitor"
	"github.com/inflion/inflion/internal/ops/rule"
	"github.com/inflion/inflion/internal/store"
	"log"
)

type ActionsConfig struct {
	Match   string         `json:"match"`
	Actions []ActionConfig `json:"actions"`
}

type ActionConfig struct {
	Type   string            `json:"type"`
	Params map[string]string `json:"params"`
}

type actionRegistry struct {
	actions map[string]action
}

var staticActionRegistry actionRegistry

func init() {
	staticActionRegistry = actionRegistry{
		actions: make(map[string]action),
	}
}

func newActionConfig(rawAction store.Action) (ActionsConfig, error) {
	action := ActionsConfig{}
	err := json.Unmarshal(rawAction.Body, &action)
	if err != nil {
		log.Println(err)
		return ActionsConfig{}, err
	}
	return action, nil
}

type actionExecutor struct {
	config ActionsConfig
	event  monitor.MonitoringEvent
}

func newActionExecutor(rawAction store.Action, event monitor.MonitoringEvent) (actionExecutor, error) {
	actionConfig, err := newActionConfig(rawAction)
	if err != nil {
		return actionExecutor{}, err
	}

	return actionExecutor{config: actionConfig, event: event}, nil
}

func (a *actionExecutor) exec(rule rule.Rule) {
	if a.config.Match == rule.Name {
		for _, actionConfig := range a.config.Actions {
			action := createAction(actionConfig)
			err := action.do(actionConfig, a.event)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func createAction(config ActionConfig) action {
	if action, ok := staticActionRegistry.actions[config.Type]; ok {
		return action
	}
	return emptyAction{}
}

type action interface {
	do(config ActionConfig, event monitor.MonitoringEvent) error
}

type emptyAction struct {
}

func (e emptyAction) do(config ActionConfig, event monitor.MonitoringEvent) error {
	log.Println("execute empty action")
	return nil
}

type slackNotificationAction struct {
	notifier Notifier
}

func (s slackNotificationAction) do(config ActionConfig, event monitor.MonitoringEvent) error {
	s.notifier.notify(event, config.Params)
	return nil
}
