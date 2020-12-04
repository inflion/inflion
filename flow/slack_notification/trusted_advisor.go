package slack_notification

import (
	"encoding/json"
	"github.com/ashwanthkumar/slack-go-webhook"
	"github.com/aws/aws-lambda-go/events"
)

type TrustedAdvisorEvent struct {
	detail          json.RawMessage
	accountMapper   *AwsAccountMapper
	CheckName       string            `json:"check-name"`
	CheckItemDetail map[string]string `json:"check-item-detail"`
	Status          string            `json:"status"`
	ResourceId      string            `json:"resource_id"`
}

func newTrustedAdvisorEvent(event events.CloudWatchEvent, accountMapper *AwsAccountMapper) (*TrustedAdvisorEvent, error) {
	e := &TrustedAdvisorEvent{detail: event.Detail, accountMapper: accountMapper}
	err := json.Unmarshal(event.Detail, e)
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (t *TrustedAdvisorEvent) title() string {
	return t.CheckName
}

func (t *TrustedAdvisorEvent) Ignore(ignoreTitleList string) bool {
	return newIgnoreList(ignoreTitleList).has(t.CheckName)
}

func (t *TrustedAdvisorEvent) authorName() string {
	return "Trusted Advisor"
}

func (t *TrustedAdvisorEvent) authorLink() string {
	return "https://console.aws.amazon.com/trustedadvisor/home?#/dashboard/"
}

func (t *TrustedAdvisorEvent) statusColor() string {
	switch t.Status {
	case "INFO":
		return "#CCCCCC"
	case "OK":
		return "#C3FFB9"
	case "WARN":
		return "#FFFF00"
	case "ERROR":
		return "#FF0000"
	}
	return "#CCCCCC"
}

func (t *TrustedAdvisorEvent) fields() []*slack.Field {
	return []*slack.Field{
		{Title: "Account", Value: t.accountMapper.awsAccount()},
		{Title: "Status", Value: t.Status},
		{Title: "Resource Id", Value: t.ResourceId},
	}
}

func (t *TrustedAdvisorEvent) Detail() json.RawMessage {
	return t.detail
}

func (t *TrustedAdvisorEvent) addMention(attachment slack.Attachment, _ map[string]string) slack.Attachment {
	return attachment
}
