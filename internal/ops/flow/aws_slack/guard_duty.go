package aws_slack

import (
	"encoding/json"
	"fmt"
	"github.com/ashwanthkumar/slack-go-webhook"
	"log"
	"strings"
)

type GuardDutyEvent struct {
	AccountId   string          `json:"accountId"`
	Region      string          `json:"region"`
	Type        string          `json:"type"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Severity    float32         `json:"severity"`
	Service     json.RawMessage `json:"service"`
	detail      json.RawMessage
}

func (g *GuardDutyEvent) SetRawDetail(detail json.RawMessage) {
	g.detail = detail
}

func (g *GuardDutyEvent) title() string {
	return g.Title
}

func (g *GuardDutyEvent) statusColor() string {
	switch g.severityLevel() {
	case "LOW":
		return "#C3FFB9"
	case "MIDDLE":
		return "#FFFF00"
	case "HIGH":
		return "#FF0000"
	}
	return "#CCCCCC"
}

func (g *GuardDutyEvent) severityLevel() string {
	switch severity := g.Severity; {
	case 1.0 < severity && severity < 3.9:
		return "LOW"
	case 4.0 < severity && severity < 6.9:
		return "MIDDLE"
	case 7.0 < severity && severity < 8.9:
		return "HIGH"
	}
	return "UNKNOWN"
}

func (g *GuardDutyEvent) authorName() string {
	return "GuardDuty"
}

func (g *GuardDutyEvent) authorLink() string {
	return "https://console.aws.amazon.com/guardduty/home?#/dashboard"
}

func (g *GuardDutyEvent) fields() []*slack.Field {
	return []*slack.Field{
		{Title: "Severity Level", Value: g.severityLevel()},
		{Title: "Type", Value: g.Type},
		{Title: "Description", Value: g.Description},
	}
}

func (g *GuardDutyEvent) Detail() string {
	return string(g.detail)
}

func (g *GuardDutyEvent) addMention(attachment slack.Attachment, params map[string]string) slack.Attachment {
	log.Print(g.severityLevel())
	switch g.severityLevel() {
	case "LOW":
		return attachment
	case "MEDIUM":
		attachment.Text = p("<!here>")
		return attachment
	case "HIGH":
		if user, ok := params["critical_mention"]; ok {
			attachment.Text = p(fmt.Sprintf("<!channel> <@%s>", user))
			return attachment
		}
		attachment.Text = p("<!channel>")
		return attachment
	}
	return attachment
}

func (g *GuardDutyEvent) Ignore(params map[string]string) bool {
	serviceJson, err := json.Marshal(g.Service)
	if err != nil {
		return false
	}

	ignore, ok := params["ignore_ip_addresses"]
	if !ok {
		return false
	}

	return newIgnoreList(ignore).contain(string(serviceJson))
}

type ignoreList struct {
	values []string
}

func newIgnoreList(commaSeparatedList string) ignoreList {
	if commaSeparatedList == "" {
		return ignoreList{}
	}

	var values []string
	for _, value := range strings.Split(commaSeparatedList, ",") {
		values = append(values, value)
	}
	return ignoreList{
		values: values,
	}
}

func (i ignoreList) has(value string) bool {
	for _, v := range i.values {
		if v == value {
			return true
		}
	}

	return false
}

func (i ignoreList) contain(value string) bool {
	for _, v := range i.values {
		if strings.Contains(value, v) {
			return true
		}
	}
	return false
}
