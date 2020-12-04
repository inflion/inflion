package slack_notification

import (
	"github.com/aws/aws-lambda-go/events"
	"log"
	"strings"
)

type AwsAccountMapper struct {
	cloudWatchEvent   events.CloudWatchEvent
	accountMappingStr string
}

func newAwsAccountMapper(cloudWatchEvent events.CloudWatchEvent, accountMappingStr string) *AwsAccountMapper {
	return &AwsAccountMapper{cloudWatchEvent: cloudWatchEvent, accountMappingStr: accountMappingStr}
}

func (m *AwsAccountMapper) awsAccount() string {
	ac, ok := m.convertAccountMapping(m.accountMappingStr)[m.cloudWatchEvent.AccountID]
	if !ok {
		return m.cloudWatchEvent.AccountID
	}

	return ac
}

//convert account mapping.
//from: 1234:name,5678:name2
//to: map[string]string{1234: name, 5678: name2}
func (m *AwsAccountMapper) convertAccountMapping(accountMappingStr string) map[string]string {
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
