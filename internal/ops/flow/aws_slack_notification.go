package flow

import (
	"encoding/json"
	"fmt"
	"github.com/ashwanthkumar/slack-go-webhook"
	"github.com/aws/aws-lambda-go/events"
	"github.com/inflion/inflion/internal/ops/flow/aws_slack"
	"log"
	"strings"
)

type AwsSlackNotificationActionExecutor struct {
	action     Action
	webhookUrl string
	channel    string
}

func (n AwsSlackNotificationActionExecutor) Run(ec ExecutionContext) (ActionResult, error) {
	log.Println("execute action: " + n.action.Type)
	log.Printf("event: %+v", ec.ExecutionFields.Fields["event"].Values)

	if !n.keyExists(n.action.Params) {
		return ActionResult{
			Action: n.action,
			Outputs: map[string]string{
				"result":  "false",
				"message": "required parameter " + strings.Join(n.requiredParamKeys(), ",") + " not found",
			},
			ExitStatus: false,
		}, nil
	}

	n.channel = n.action.Params["channel"]
	n.webhookUrl = n.action.Params["webhook_url"]

	cw := &events.CloudWatchEvent{}
	b, err := json.Marshal(ec.ExecutionFields.Fields["event"].Values)
	err = json.Unmarshal(b, cw)
	if err != nil {
		return ActionResult{
			Action: n.action,
			Outputs: map[string]string{
				"result":  "false",
				"message": "this event is not in cloud watch format",
			},
			ExitStatus: false,
		}, nil
	}

	awsSlackEvent, err := aws_slack.NewAwsSlackEvent(cw.Source)
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

	err = json.Unmarshal(cw.Detail, &awsSlackEvent)
	if err != nil {
		return ActionResult{
			Action: n.action,
			Outputs: map[string]string{
				"result":  "false",
				"message": "unknown event format",
			},
			ExitStatus: false,
		}, nil
	}

	awsSlackEvent.SetRawDetail(ec.ExecutionFields.Fields["raw-event"].Values["json"].(json.RawMessage))

	if awsSlackEvent.Ignore(n.action.Params) {
		return ActionResult{
			Action: n.action,
			Outputs: map[string]string{
				"result":  "true",
				"message": fmt.Sprintf("ignore event:%+v", awsSlackEvent.Detail()),
			},
			ExitStatus: true,
		}, nil
	}

	err = n.send(aws_slack.BuildAttachments(awsSlackEvent, n.awsAccount(ec), n.action.Params))
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

func (n *AwsSlackNotificationActionExecutor) send(attachments []slack.Attachment) error {
	payload := slack.Payload{
		Username:    "Inflion",
		Channel:     "#" + n.channel,
		IconEmoji:   ":lion_face:",
		Attachments: attachments,
	}
	err := slack.Send(n.webhookUrl, "", payload)
	if len(err) > 0 {
		log.Printf("error: %s\n", err)
		return err[0]
	}

	return nil
}

func (n AwsSlackNotificationActionExecutor) keyExists(params map[string]string) bool {
	rp := n.requiredParamKeys()
	for _, p := range rp {
		if _, ok := params[p]; !ok {
			return false
		}
	}

	return true
}

func (n AwsSlackNotificationActionExecutor) requiredParamKeys() []string {
	return []string{
		"webhook_url",
		"channel",
	}
}

func (n AwsSlackNotificationActionExecutor) awsAccount(ec ExecutionContext) string {
	account, ok := ec.ExecutionFields.Fields["event"].Values["account"].(string)
	if !ok {
		return "unknown"
	}

	accountMapping, ok := n.action.Params["account_mapping"]
	if !ok {
		return account
	}

	ac, ok := n.convertAccountMapping(accountMapping)[account]
	if !ok {
		return account
	}

	return ac
}

// convert account mapping.
// from: 1234:name,5678:name2
// to: map[string]string{1234: name, 5678: name2}
func (n AwsSlackNotificationActionExecutor) convertAccountMapping(accountMappingStr string) map[string]string {
	mapping := map[string]string{}
	for _, m := range strings.Split(accountMappingStr, ",") {
		log.Printf("%+v", m)
		tmp := strings.Split(m, ":")
		log.Printf("%+v", tmp)
		if len(tmp) > 1 {
			mapping[tmp[0]] = tmp[1]
		}
	}
	return mapping
}
