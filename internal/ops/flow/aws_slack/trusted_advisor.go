package aws_slack

import (
	"encoding/json"
	"github.com/ashwanthkumar/slack-go-webhook"
)

type TrustedAdvisorEvent struct {
	CheckName       string            `json:"check-name"`
	CheckItemDetail map[string]string `json:"check-item-detail"`
	Status          string            `json:"status"`
	ResourceId      string            `json:"resource_id"`
	detail          json.RawMessage
}

func (t *TrustedAdvisorEvent) SetRawDetail(detail json.RawMessage) {
	t.detail = detail
}

func (t *TrustedAdvisorEvent) title() string {
	return t.CheckName
}

func (t *TrustedAdvisorEvent) Ignore(params map[string]string) bool {
	ignoreTitleList, ok := params["ignore_title"]
	if !ok {
		return false
	}

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
		{Title: "Status", Value: t.Status},
		{Title: "Resource Id", Value: t.ResourceId},
	}
}

func (t *TrustedAdvisorEvent) Detail() string {
	return string(t.detail)
}

func (t *TrustedAdvisorEvent) addMention(attachment slack.Attachment, _ map[string]string) slack.Attachment {
	return attachment
}
