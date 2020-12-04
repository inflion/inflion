package slack_notification

import (
	"encoding/json"
	"github.com/ashwanthkumar/slack-go-webhook"
	"github.com/aws/aws-lambda-go/events"
)

type ConfigEvent struct {
	detail              json.RawMessage
	accountMapper       *AwsAccountMapper
	Title               string                 `json:"messageType"`
	ResourceId          string                 `json:"resourceId"`
	Region              string                 `json:"awsRegion"`
	ConfigRuleName      string                 `json:"configRuleName"`
	ResourceType        string                 `json:"resourceType"`
	NewEvaluationResult map[string]interface{} `json:"newEvaluationResult"`
	Account             string                 `json:"account"`
}

func newConfigEvent(event events.CloudWatchEvent, accountMapper *AwsAccountMapper) (*ConfigEvent, error) {
	e := &ConfigEvent{detail: event.Detail, accountMapper: accountMapper}
	err := json.Unmarshal(event.Detail, e)
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (c *ConfigEvent) title() string {
	return c.Title
}

func (c *ConfigEvent) SetTitle(title string) {
	c.Title = title
}

func (c *ConfigEvent) SetRawDetail(detail json.RawMessage) {
	c.detail = detail
}

func (c *ConfigEvent) statusColor() string {
	switch c.getComplianceType() {
	case "COMPLIANT":
		return "#C3FFB9"
	case "NON_COMPLIANT":
		return "#FF0000"
	}
	return "#CCCCCC"
}

func (c *ConfigEvent) getComplianceType() string {
	var complianceType string
	if t, ok := c.NewEvaluationResult["complianceType"].(string); !ok {
		complianceType = "unknown"
	} else {
		complianceType = t
	}

	return complianceType
}

func (c *ConfigEvent) authorName() string {
	return "Config"
}

func (c *ConfigEvent) authorLink() string {
	return "https://console.aws.amazon.com/config/home/"
}

func (c *ConfigEvent) fields() []*slack.Field {
	return []*slack.Field{
		{Title: "Account", Value: c.accountMapper.awsAccount()},
		{Title: "Status", Value: c.getComplianceType()},
		{Title: "RuleName", Value: c.ConfigRuleName},
		{Title: "ResourceId", Value: c.ResourceId},
		{Title: "ResourceType", Value: c.ResourceType},
	}
}

func (c *ConfigEvent) Detail() json.RawMessage {
	return c.detail
}

func (c *ConfigEvent) Ignore(_ string) bool {
	return false
}

func (c *ConfigEvent) addMention(attachment slack.Attachment, _ map[string]string) slack.Attachment {
	return attachment
}
