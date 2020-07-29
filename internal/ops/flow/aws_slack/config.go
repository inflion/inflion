package aws_slack

import (
	"encoding/json"
	"github.com/ashwanthkumar/slack-go-webhook"
)

type ConfigEvent struct {
	Title               string                 `json:"messageType"`
	ResourceId          string                 `json:"resourceId"`
	Region              string                 `json:"awsRegion"`
	ConfigRuleName      string                 `json:"configRuleName"`
	ResourceType        string                 `json:"resourceType"`
	NewEvaluationResult map[string]interface{} `json:"newEvaluationResult"`
	Account             string                 `json:"account"`
	detail              json.RawMessage
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
		{Title: "Status", Value: c.getComplianceType()},
		{Title: "RuleName", Value: c.ConfigRuleName},
		{Title: "ResourceId", Value: c.ResourceId},
		{Title: "ResourceType", Value: c.ResourceType},
	}
}

func (c *ConfigEvent) Detail() string {
	return string(c.detail)
}

func (c *ConfigEvent) Ignore(_ map[string]string) bool {
	return false
}

func (c *ConfigEvent) addMention(attachment slack.Attachment, _ map[string]string) slack.Attachment {
	return attachment
}
