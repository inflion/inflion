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
	"errors"
)

type ActionLoader interface {
	Load(action Action) (ActionExecutor, error)
}

type AggregateActionLoader struct{}

func NewAggregateActionLoader() *AggregateActionLoader {
	return &AggregateActionLoader{}
}

func (a AggregateActionLoader) Load(action Action) (ActionExecutor, error) {
	loaders := []ActionLoader{
		EmbeddedActionLoader{},
	}

	for _, l := range loaders {
		a, err := l.Load(action)
		if err != nil {
			return a, err
		}

		return a, nil
	}

	return NullActionExecutor{}, errors.New("no action found at type: " + action.Type)
}

type EmbeddedActionLoader struct{}

func (e EmbeddedActionLoader) Load(action Action) (ActionExecutor, error) {
	switch action.Type {
	case "Params":
		return ParamsActionExecutor{action: action}, nil
	case "config":
		return ConfigActionExecutor{action: action}, nil
	case "matcher":
		return MatcherActionExecutor{action: action}, nil
	case "instance":
		return InstanceActionExecutor{action: action}, nil
	case "instance-data":
		return InstanceDataActionExecutor{action: action}, nil
	case "logging":
		return LoggingActionExecutor{action: action}, nil
	case "pagerduty":
		return PagerDutyActionExecutor{action: action}, nil
	case "notification":
		return NewSlackNotificationActionExecutor(action)
	}

	return NullActionExecutor{}, errors.New("no embedded action found at " + action.Type)
}
