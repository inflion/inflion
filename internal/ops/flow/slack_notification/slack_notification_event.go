package slack_notification

import (
	"encoding/json"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/inflion/inflion/internal/ops/flow/context"
)

func NewSlackNotificationEvent(notificationType string, actionParams map[string]string, ctx context.ExecutionContext) (SlackNotificationEvent, error) {
	byteEvent := ctx.Event().RawBody()

	switch notificationType {
	case "log":
		return NewLogEvent(byteEvent, actionParams, ctx)
	case "aws":
		return newAwsSlackEvent(byteEvent, actionParams)
	}

	return nil, errors.New("unknown event")
}

func newAwsSlackEvent(byteEvent []byte, actionParams map[string]string) (SlackNotificationEvent, error) {
	var event events.CloudWatchEvent
	err := json.Unmarshal(byteEvent, &event)
	if err != nil {
		return nil, errors.New("unknown aws event")
	}

	var accountMapping string
	if ac, ok := actionParams["account_mapping"]; ok {
		accountMapping = ac
	}
	switch event.Source {
	case "aws.trustedadvisor":
		return newTrustedAdvisorEvent(event, newAwsAccountMapper(event, accountMapping))
	case "aws.guardduty":
		return NewGuardDutyEvent(event, newAwsAccountMapper(event, accountMapping))
	case "aws.config":
		return newConfigEvent(event, newAwsAccountMapper(event, accountMapping))
	}

	return nil, errors.New("unknown event type")
}
