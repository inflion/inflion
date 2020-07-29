package aws_slack

import (
	"encoding/json"
	"errors"
	"github.com/ashwanthkumar/slack-go-webhook"
)

type AwsSlackEvent interface {
	title() string
	statusColor() string
	authorName() string
	authorLink() string
	fields() []*slack.Field
	Detail() string
	Ignore(map[string]string) bool
	addMention(slack.Attachment, map[string]string) slack.Attachment
	SetRawDetail(json.RawMessage)
}

func NewAwsSlackEvent(source string) (AwsSlackEvent, error) {
	switch source {
	case "aws.trustedadvisor":
		return &TrustedAdvisorEvent{}, nil
	case "aws.guardduty":
		return &GuardDutyEvent{}, nil
	case "aws.config":
		return &ConfigEvent{}, nil
	}

	return nil, errors.New("unknown event type")
}

func BuildAttachments(awsEvent AwsSlackEvent, account string, params map[string]string) []slack.Attachment {
	a := slack.Attachment{
		Title:      p(awsEvent.title()),
		Color:      p(awsEvent.statusColor()),
		AuthorName: p(awsEvent.authorName()),
		AuthorLink: p(awsEvent.authorLink()),
		Fields:     awsEvent.fields(),
	}

	a.Fields = append(a.Fields, &slack.Field{
		Title: "Account",
		Value: account,
	})

	var attachments []slack.Attachment
	attachments = append(attachments, awsEvent.addMention(a, params))

	attachments = append(attachments, slack.Attachment{
		Title: p("Detail"),
		Color: p(awsEvent.statusColor()),
		Text:  p(awsEvent.Detail()),
	})

	return attachments
}

// Utility for convert string to string pointer
func p(v string) *string {
	return &v
}
