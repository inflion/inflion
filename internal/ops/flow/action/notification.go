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
	"errors"
	"fmt"
	"github.com/ashwanthkumar/slack-go-webhook"
	"log"
	"strings"
)

type AwsSlackNotifier struct {
	AccountMapping map[string]string
}

type MonitoringEvent struct {
	Project string          `json:"project"`
	Body    json.RawMessage `json:"body"`
}

type CloudWatchEvent struct {
	Detail json.RawMessage `json:"detail"`
}

type TrustedAdvisorEvent struct {
	CheckName       string            `json:"check-name"`
	CheckItemDetail map[string]string `json:"check-item-detail"`
	Status          string            `json:"status"`
	ResourceId      string            `json:"resource_id"`
}

func (t *TrustedAdvisorEvent) StatusColor() string {
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

func (t *TrustedAdvisorEvent) StatusIcon() string {
	switch t.Status {
	case "OK":
		return ":ok:"
	case "INFO":
		return ":information_source:"
	case "WARN":
		return ":warning:"
	case "ERROR":
		return ":rotating_light:"
	}
	return ":mag_right:"
}

type GuardDutyEvent struct {
	AccountId   string          `json:"accountId"`
	Region      string          `json:"accountId"`
	Type        string          `json:"type"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Severity    float32         `json:"severity"`
	Service     json.RawMessage `json:"service"`
}

func (g *GuardDutyEvent) SeverityLevel() string {
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

func (g *GuardDutyEvent) StatusColor() string {
	switch g.SeverityLevel() {
	case "LOW":
		return "#C3FFB9"
	case "MIDDLE":
		return "#FFFF00"
	case "HIGH":
		return "#FF0000"
	}
	return "#CCCCCC"
}

func (g *GuardDutyEvent) StatusIcon() string {
	switch g.SeverityLevel() {
	case "LOW":
		return ":information_source:"
	case "MIDDLE":
		return ":warning:"
	case "HIGH":
		return ":rotating_light:"
	}
	return ":mag_right:"
}

func (s *AwsSlackNotifier) Notify(params map[string]string, event map[string]interface{}, rawEvent json.RawMessage) error {
	webhookUrl, ok := params["webhook_url"]
	if !ok {
		return errors.New("webhook_url not specified")
	}
	slackChannel, ok := params["channel"]
	if !ok {
		return errors.New("webhook_channel not specified")
	}

	sender := slackSender{
		webhookUrl: webhookUrl,
		channel:    slackChannel,
	}

	me := MonitoringEvent{}
	err := json.Unmarshal(rawEvent, &me)
	if err != nil {
		return errors.New("invalid json")
	}

	if source, ok := event["source"]; ok {
		cwe := CloudWatchEvent{}
		err = json.Unmarshal(me.Body, &cwe)
		if err != nil {
			return errors.New("invalid json")
		}

		var attachments []slack.Attachment
		if source == "aws.trustedadvisor" {
			attachments, err = s.trustedAdvisor(params, event, cwe, rawEvent)
		} else if source == "aws.guardduty" {
			attachments, err = s.guardDuty(params, event, cwe, rawEvent)
		}
		if err != nil {
			err = sender.send(s.error(err, rawEvent))
			if err != nil {
				log.Println(err)
			}
		}

		if attachments == nil {
			return nil
		}

		err = sender.send(attachments)
		if err != nil {
			log.Println(err)
		}
	}

	return errors.New("no such notification type")
}

func (s *AwsSlackNotifier) error(err error, rawEvent json.RawMessage) []slack.Attachment {
	attachments := []slack.Attachment{
		{
			Fields: []*slack.Field{
				{Title: "Status", Value: "error"},
				{Title: "Error", Value: fmt.Sprintf("%s", err)},
				{Title: "Input", Value: string(rawEvent)},
			},
		},
	}
	return attachments
}

type ignoreList struct {
	values []string
}

func newIgnoreList(commaSeparatedList string) ignoreList {
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

func (i ignoreList) contains(value string) bool {
	for _, v := range i.values {
		if strings.Contains(value, v) {
			return true
		}
	}
	return false
}


func (s *AwsSlackNotifier) trustedAdvisor(params map[string]string, event map[string]interface{}, cwevent CloudWatchEvent, rawEvent json.RawMessage) ([]slack.Attachment, error) {
	var attachments []slack.Attachment

	tae := TrustedAdvisorEvent{}

	err := json.Unmarshal(cwevent.Detail, &tae)
	if err != nil {
		if len(rawEvent) != 0 {
			log.Println("unmarshal error", err)
			return nil, err
		}
		return nil, errors.New("json is empty")
	}

	itemDetailJson, err := json.Marshal(tae.CheckItemDetail)
	if err != nil {
		itemDetailJson = []byte{}
	}

	var accountName string
	if n, ok := s.AccountMapping[event["account"].(string)]; !ok {
		accountName = "unknown"
	} else {
		accountName = n
	}

	if ignoreTitleList, ok := params["ignore_title"]; ok {
		if newIgnoreList(ignoreTitleList).has(tae.CheckName) {
			log.Println("ignore: " + tae.CheckName)
			return nil, nil
		}
	}

	attachments = append(attachments, slack.Attachment{
		Title:      p(tae.CheckName),
		Color:      p(tae.StatusColor()),
		AuthorName: p("Trusted Advisor"),
		AuthorLink: p("https://console.aws.amazon.com/trustedadvisor/home?#/dashboard"),
		Fields: []*slack.Field{
			{Title: "アカウント", Value: accountName},
			{Title: "ステータス", Value: tae.Status},
			{Title: "リソース", Value: tae.ResourceId},
		},
	})
	attachments = append(attachments, slack.Attachment{
		Title: p("詳細"),
		Color: p(tae.StatusColor()),
		Text:  p(string(itemDetailJson)),
	})

	return attachments, nil
}

func (s *AwsSlackNotifier) guardDuty(params map[string]string, event map[string]interface{}, cwevent CloudWatchEvent, rawEvent json.RawMessage) ([]slack.Attachment, error) {
	var attachments []slack.Attachment

	gde := GuardDutyEvent{}

	err := json.Unmarshal(cwevent.Detail, &gde)
	if err != nil {
		if len(rawEvent) != 0 {
			log.Println("unmarshal error", err)
			return nil, err
		}
		return nil, errors.New("json is empty")
	}

	serviceJson, err := json.Marshal(gde.Service)
	if err != nil {
		serviceJson = []byte{}
	}

	var accountName string
	if n, ok := s.AccountMapping[event["account"].(string)]; !ok {
		accountName = "unknown"
	} else {
		accountName = n
	}

	a := slack.Attachment{
		Title:      p(gde.Title),
		Color:      p(gde.StatusColor()),
		AuthorName: p("GuardDuty"),
		AuthorLink: p("https://console.aws.amazon.com/guardduty/home?#/dashboard"),
		Fields: []*slack.Field{
			{Title: "アカウント", Value: accountName},
			{Title: "重要度", Value: gde.SeverityLevel()},
			{Title: "タイプ", Value: gde.Type},
			{Title: "説明", Value: gde.Description},
		},
	}

	switch gde.SeverityLevel() {
	case "LOW":
	case "MEDIUM":
		a.Text = p("<!here>")
	case "HIGH":
		if user, ok := params["critical_mention"]; ok {
			a.Text = p(fmt.Sprintf("<!channel> <@%s>", user))
		} else {
			a.Text = p("<!channel>")
		}
	}

	if ignoreAddressList, ok := params["ignore_ip_addresses"]; ok {
		if newIgnoreList(ignoreAddressList).contains(string(serviceJson)) {
			return nil, nil
		}
	}

	attachments = append(attachments, a)
	attachments = append(attachments, slack.Attachment{
		Title: p("詳細"),
		Color: p(gde.StatusColor()),
		Text:  p(string(serviceJson)),
	})

	return attachments, nil
}

type slackSender struct {
	webhookUrl string
	channel    string
}

func (s *slackSender) send(attachments []slack.Attachment) error {
	payload := slack.Payload{
		Username:    "Inflion",
		Channel:     "#" + s.channel,
		IconEmoji:   ":lion_face:",
		Attachments: attachments,
	}

	err := slack.Send(s.webhookUrl, "", payload)
	if len(err) > 0 {
		log.Printf("error: %s\n", err)
		return err[0]
	}
	return nil
}

// Utility for convert string to string pointer
func p(v string) *string {
	return &v
}
