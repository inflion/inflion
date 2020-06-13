package broker

import (
	"context"
	"github.com/inflion/inflion/internal/ops/flow"
	"github.com/inflion/inflion/internal/ops/monitor"
	"github.com/inflion/inflion/internal/ops/rule"
	"github.com/inflion/inflion/internal/store"
	"log"
)

type eventProcessor interface {
	process(event monitor.MonitoringEvent) error
}

type defaultEventProcessor struct {
	matcher rule.EventMatcher
	querier store.Querier
}

func newDefaultEventProcessor(querier store.Querier, matcher rule.EventMatcher) eventProcessor {
	return defaultEventProcessor{
		matcher: matcher,
		querier: querier,
	}
}

func (d defaultEventProcessor) process(event monitor.MonitoringEvent) error {
	ctx := context.Background()

	matchedRules, err := d.matcher.GetRulesMatchesTo(event)
	if err != nil {
		log.Println(err)
		return err
	}

	for _, rule := range matchedRules {

		// TODO |############################################|
		// TODO |    Use flow.store instead of querier.      |
		// TODO |############################################|

		flows, err := d.querier.GetFlowByName(ctx,
			store.GetFlowByNameParams{
				ProjectID: event.ProjectId,
				FlowName:  rule.Target,
			},
		)

		if err != nil {
			log.Println(err)
			return err
		}

		for _, f := range flows {
			f := flow.NewOpsFlow(ByteRecipeReader{body: f.Body})
			err := f.Run()
			if err != nil {
				log.Println(err)
			}
		}
	}

	return nil
}

type ByteRecipeReader struct {
	body []byte
}

func (b ByteRecipeReader) Read() (flow.Recipe, error) {
	recipe, err := flow.Unmarshal(b.body)
	if err != nil {
		return flow.Recipe{}, err
	}
	return recipe, nil
}
