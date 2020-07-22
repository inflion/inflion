package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/configservice"
	"log"
)

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	input := configservice.PutConfigRuleInput{ConfigRule: &configservice.ConfigRule{
		ConfigRuleName: aws.String("test-from-gopher"),
		Source: &configservice.Source{
			Owner:            aws.String("AWS"),
			SourceDetails:    nil,
			SourceIdentifier: aws.String("CLOUD_TRAIL_ENABLED"),
		},
	}}

	cs := configservice.New(sess)
	res, err := cs.PutConfigRule(&input)
	if err != nil {
		log.Println(err)
	}

	out, err := cs.PutRemediationConfigurations(
		&configservice.PutRemediationConfigurationsInput{RemediationConfigurations: []*configservice.RemediationConfiguration{
			{
				ConfigRuleName:           aws.String("test-from-gopher"),
				CreatedByService:         nil,
				ExecutionControls:        nil,
				MaximumAutomaticAttempts: nil,
				Parameters: map[string]*configservice.RemediationParameterValue{
					"TopicArn": {
						StaticValue: &configservice.StaticValue{Values: []*string{aws.String("test")}},
					},
					"Message": {
						StaticValue: &configservice.StaticValue{Values: []*string{aws.String("test")}},
					},
					"AutomationAssumeRole": {
						StaticValue: &configservice.StaticValue{Values: []*string{aws.String("test")}},
					},
				},
				ResourceType:        nil,
				RetryAttemptSeconds: nil,
				TargetId:            aws.String("AWS-PublishSNSNotification"),
				TargetType:          aws.String("SSM_DOCUMENT"),
				TargetVersion:       aws.String("1"),
			},
		}},
	)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(out)

	fmt.Printf("%+v", res)
}

