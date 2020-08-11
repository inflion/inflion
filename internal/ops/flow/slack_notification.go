package flow

import (
	"errors"
	"fmt"
	"github.com/ashwanthkumar/slack-go-webhook"
	"github.com/inflion/inflion/internal/ops/flow/context"
	"github.com/inflion/inflion/internal/ops/flow/slack_notification"
	"log"
)

type SlackNotificationActionExecutor struct {
	action            Action
	notificationEvent slack_notification.SlackNotificationEvent
	params            slackNotificationParam
}

func NewSlackNotificationActionExecutor(action Action) (*SlackNotificationActionExecutor, error) {
	params, err := newSlackNotificationParam(action.Params)
	if err != nil {
		return nil, err
	}

	return &SlackNotificationActionExecutor{action: action, params: *params}, nil
}

type slackNotificationParam struct {
	webhookUrl       string
	channel          string
	notificationType string
}

func newSlackNotificationParam(params map[string]string) (*slackNotificationParam, error) {
	webhookUrl, ok := params["webhook_url"]
	if !ok {
		return nil, errors.New("action parameter \"webhook_url\" not found")
	}
	channel, ok := params["channel"]
	if !ok {
		return nil, errors.New("action parameter \"channel\" not found")
	}
	notificationType, ok := params["notification_type"]
	if !ok {
		return nil, errors.New("action parameter \"notification_type\" not found")
	}

	return &slackNotificationParam{webhookUrl: webhookUrl, channel: channel, notificationType: notificationType}, nil
}

func (n *SlackNotificationActionExecutor) Run(ec context.ExecutionContext) (ActionResult, error) {
	log.Println("execute action: " + n.action.Type)

	notificationEvent, err := slack_notification.NewSlackNotificationEvent(
		n.params.notificationType,
		n.action.Params,
		ec,
	)

	if err != nil {
		return ActionResult{
			Action: n.action,
			Outputs: map[string]string{
				"result":  "false",
				"message": err.Error(),
			},
			ExitStatus: false,
		}, nil
	}
	n.notificationEvent = notificationEvent

	if adr, ok := n.action.Params["ignore_ip_address"]; ok {
		if n.notificationEvent.Ignore(adr) {
			return ActionResult{
				Action: n.action,
				Outputs: map[string]string{
					"result":  "true",
					"message": fmt.Sprintf("ignore event:%+v", n.notificationEvent.Detail()),
				},
				ExitStatus: true,
			}, nil
		}
	}

	err = n.send(slack_notification.NewSlackAttachmentBuilder(n.notificationEvent, n.action.Params))
	if err != nil {
		return ActionResult{
			Action: n.action,
			Outputs: map[string]string{
				"result":  "false",
				"message": err.Error(),
			},
			ExitStatus: false,
		}, nil
	}

	return ActionResult{
		Action: n.action,
		Outputs: map[string]string{
			"result": "true",
		},
		ExitStatus: true,
	}, nil
}

func (n *SlackNotificationActionExecutor) send(builder *slack_notification.SlackAttachmentBuilder) error {
	log.Print(builder.BuildAttachments())
	payload := slack.Payload{
		Username:    "Inflion",
		Channel:     "#" + n.params.channel,
		IconEmoji:   ":lion_face:",
		Attachments: builder.BuildAttachments(),
	}

	err := slack.Send(n.params.webhookUrl, "", payload)
	if len(err) > 0 {
		log.Printf("error: %s\n", err)
		return err[0]
	}

	return nil
}
