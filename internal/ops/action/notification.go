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
	"context"
	"github.com/ashwanthkumar/slack-go-webhook"
	"github.com/inflion/inflion/internal/ops/monitor"
	"github.com/inflion/inflion/internal/store"
	"log"
	"time"
)

var staticSlackAttachmentRegistry map[string]SlackAttachment

func init() {
	staticSlackAttachmentRegistry = map[string]SlackAttachment{
		monitor.CPUUtilization.String():   CpuUtilizationSlackAttachment{},
		monitor.OpenPortDetected.String(): SecurityGroupSlackAttachment{},
	}
}

type params map[string]string

type Notifier interface {
	notify(event monitor.MonitoringEvent, params params) []error
}

type notification struct {
	querier   store.Querier
	Notifiers []Notifier
}

func newNotification(querier store.Querier) notification {
	return notification{
		Notifiers: []Notifier{&slackNotifier{querier: querier}},
	}
}

func (n *notification) notify(event monitor.MonitoringEvent, params params) []error {
	var errs []error
	for _, n := range n.Notifiers {
		err := n.notify(event, params)
		for _, e := range err {
			if e != nil {
				errs = append(errs, e)
			}
		}
	}
	return errs
}

type throttledNotification struct {
	querier   store.Querier
	notifiers []Notifier
	memento   map[string]time.Time
}

func NewThrottledNotification(q store.Querier, n []Notifier) Notifier {
	return &throttledNotification{
		querier:   q,
		notifiers: n,
		memento:   map[string]time.Time{},
	}
}

func (t *throttledNotification) notify(event monitor.MonitoringEvent, params params) []error {
	var errs []error
	hash := event.Hash([]string{"Value"})
	now := time.Now()

	for _, n := range t.notifiers {
		if history, ok := t.memento[hash]; ok {
			// Event sent 30 min ago.
			if time.Minute*30 < now.Sub(history) {
				log.Println("30分以上前に通知したので際通知")
				err := n.notify(event, params)
				for _, e := range err {
					if e != nil {
						errs = append(errs, e)
					}
				}
				t.memento[hash] = now
			} else {
				log.Println("通知済みなので無視")
			}
		} else {
			err := n.notify(event, params)
			for _, e := range err {
				if e != nil {
					errs = append(errs, e)
				}
			}
			t.memento[hash] = now
		}
	}
	return errs
}

type slackNotifier struct {
	querier store.Querier
}

func (s *slackNotifier) notify(event monitor.MonitoringEvent, params params) []error {
	a := createSlackAttachment(event)

	webhooks, err := s.querier.GetSlackWebHooks(context.Background(), event.ProjectId)
	if err != nil {
		log.Println(err)
		return []error{err}
	}

	for _, webhook := range webhooks {
		if to, ok := params["to"]; ok {
			if webhook.Name == to {
				s.send(a.attachment(event), webhook.Channel, webhook.WebhookUrl)
			}
		}
	}

	return nil
}

func (s *slackNotifier) send(attachment slack.Attachment, channel string, webhookUrl string) {
	payload := slack.Payload{
		Text:        "notify",
		Username:    "inflion",
		Channel:     "#" + channel,
		IconEmoji:   ":lion_face:",
		Attachments: []slack.Attachment{attachment},
	}

	err := slack.Send(webhookUrl, "", payload)
	if len(err) > 0 {
		log.Printf("error: %s\n", err)
	}
}

type SlackAttachment interface {
	attachment(event monitor.MonitoringEvent) slack.Attachment
}

type emptySlackAttachment struct {
}

func (s emptySlackAttachment) attachment(event monitor.MonitoringEvent) slack.Attachment {
	return slack.Attachment{}
}

type SecurityGroupSlackAttachment struct {
}

func (s SecurityGroupSlackAttachment) attachment(event monitor.MonitoringEvent) slack.Attachment {
	sgId := event.Values["SecurityGroupId"]

	a := slack.Attachment{}
	a.AddField(slack.Field{Title: "Message", Value: event.Message})
	a.AddField(slack.Field{Title: "SecurityGroup", Value: sgId.(string)})
	log.Println(event.Values["OpenPorts"])
	if v, ok := event.Values["OpenPorts"].(string); ok {
		a.AddField(slack.Field{Title: "Open Port", Value: v})
	}
	a.AddAction(slack.Action{Type: "button", Text: "Close Port", Url: "http://localhost:3000/project/1/action/close", Style: "primary"})
	return a
}

type CpuUtilizationSlackAttachment struct {
}

func (c CpuUtilizationSlackAttachment) attachment(event monitor.MonitoringEvent) slack.Attachment {
	id := event.Values["InstanceId"]

	a := slack.Attachment{}
	a.AddField(slack.Field{Title: "Warning", Value: id.(string)}).AddField(slack.Field{Title: "ExitStatus", Value: "Completed"})
	a.AddAction(slack.Action{Type: "button", Text: "I got it", Url: "http://localhost:3000/reboot", Style: "primary"})
	return a
}

func createSlackAttachment(event monitor.MonitoringEvent) SlackAttachment {
	if slackAttachment, ok := staticSlackAttachmentRegistry[event.Type]; ok {
		return slackAttachment
	}

	return emptySlackAttachment{}
}
