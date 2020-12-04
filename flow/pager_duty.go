package flow

import (
	"fmt"
	"github.com/PagerDuty/go-pagerduty"
	"github.com/inflion/inflion/flow/context"
	"log"
	"strings"
	"time"
)

type PagerDutyActionExecutor struct {
	action Action
}

func (p PagerDutyActionExecutor) Run(ctx context.ExecutionContext) (ActionResult, error) {
	log.Println("execute action: " + p.action.Type)
	log.Printf("action params: %+v", p.action.Params)

	key, ok := p.action.Params["key"]
	if !ok {
		return ActionResult{
			Action: p.action,
			Outputs: map[string]string{
				"result":  "false",
				"message": "parameter \"key\" not found",
			},
			ExitStatus: false,
		}, nil
	}

	source := "unknown"
	s, ok := ctx.Event().GetValue("source")
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
			Details:   ctx.Event().RawBody(),
		},
	}

	resp, err := pagerduty.ManageEvent(pagerdutyEvent)
	if err != nil {
		return ActionResult{
			Action: p.action,
			Outputs: map[string]string{
				"result":  "false",
				"message": fmt.Sprintf("%+v", err),
			},
			ExitStatus: false,
		}, nil
	}

	return ActionResult{
		Action: p.action,
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
