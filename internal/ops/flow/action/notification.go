package action

import (
	"github.com/ashwanthkumar/slack-go-webhook"
	"github.com/inflion/inflion/internal/ops/monitor"
	"log"
)

var staticSlackAttachmentRegistry map[string]SlackAttachment

func init() {
	staticSlackAttachmentRegistry = map[string]SlackAttachment{
		monitor.CPUUtilization.String():   CpuUtilizationSlackAttachment{},
		monitor.OpenPortDetected.String(): SecurityGroupSlackAttachment{},
	}
}

type SlackNotifier struct {
}

func (s *SlackNotifier) Notify(params map[string]string) []error {
	a := slack.Attachment{}
	a.AddField(slack.Field{Title: "Message", Value: params["message"]})

	if webhookUrl, ok := params["webhook_url"]; ok {
		if webhookChannel, ok := params["channel"]; ok {
			s.send(a, webhookChannel, webhookUrl)
		}
	}

	return nil
}

func (s *SlackNotifier) send(attachment slack.Attachment, channel string, webhookUrl string) {
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
	sgId := event.Body["SecurityGroupId"]

	a := slack.Attachment{}
	a.AddField(slack.Field{Title: "Message", Value: event.Body["Message"].(string)})
	a.AddField(slack.Field{Title: "SecurityGroup", Value: sgId.(string)})
	log.Println(event.Body["OpenPorts"])
	if v, ok := event.Body["OpenPorts"].(string); ok {
		a.AddField(slack.Field{Title: "Open Port", Value: v})
	}
	a.AddAction(slack.Action{Type: "button", Text: "Close Port", Url: "http://localhost:3000/project/1/action/close", Style: "primary"})
	return a
}

type CpuUtilizationSlackAttachment struct {
}

func (c CpuUtilizationSlackAttachment) attachment(event monitor.MonitoringEvent) slack.Attachment {
	id := event.Body["InstanceId"]

	a := slack.Attachment{}
	a.AddField(slack.Field{Title: "Warning", Value: id.(string)}).AddField(slack.Field{Title: "ExitStatus", Value: "Completed"})
	a.AddAction(slack.Action{Type: "button", Text: "I got it", Url: "http://localhost:3000/reboot", Style: "primary"})
	return a
}

func createSlackAttachment(event monitor.MonitoringEvent) SlackAttachment {
	if slackAttachment, ok := staticSlackAttachmentRegistry[event.Body["Type"].(string)]; ok {
		return slackAttachment
	}

	return emptySlackAttachment{}
}
