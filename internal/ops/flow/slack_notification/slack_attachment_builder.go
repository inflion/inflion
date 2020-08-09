package slack_notification

import (
	"bytes"
	"encoding/json"
	"github.com/ashwanthkumar/slack-go-webhook"
)

type SlackNotificationEvent interface {
	title() string
	statusColor() string
	authorName() string
	authorLink() string
	fields() []*slack.Field
	Detail() json.RawMessage
	Ignore(string) bool
	addMention(slack.Attachment, map[string]string) slack.Attachment
}

type SlackAttachmentBuilder struct {
	event        SlackNotificationEvent
	actionParams map[string]string
}

func NewSlackAttachmentBuilder(event SlackNotificationEvent, actionParams map[string]string) *SlackAttachmentBuilder {
	return &SlackAttachmentBuilder{event: event, actionParams: actionParams}
}

func (b *SlackAttachmentBuilder) BuildAttachments() []slack.Attachment {
	a := slack.Attachment{
		Title:      p(b.event.title()),
		Color:      p(b.event.statusColor()),
		AuthorName: p(b.event.authorName()),
		AuthorLink: p(b.event.authorLink()),
		Fields:     b.event.fields(),
	}

	var attachments []slack.Attachment
	attachments = append(attachments, b.event.addMention(a, b.actionParams))

	attachments = append(attachments, slack.Attachment{
		Title: p("Detail"),
		Color: p(b.event.statusColor()),
		Text:  p(b.IndentDetail()),
	})

	return attachments
}

func (b *SlackAttachmentBuilder) IndentDetail() string {
	var buf bytes.Buffer
	err := json.Indent(&buf, b.event.Detail(), "", "\t")
	if err != nil {
		return string(b.event.Detail())
	}

	return buf.String()
}

// Utility for convert string to string pointer
func p(v string) *string {
	return &v
}
