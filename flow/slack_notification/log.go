package slack_notification

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ashwanthkumar/slack-go-webhook"
	"github.com/inflion/inflion/flow/context"
	"github.com/olivere/elastic/v7"
	"log"
)

type LogEvent struct {
	event        elastic.SearchHit
	detail       json.RawMessage
	term         string
	indexPattern string
	fieldName    string
	kibanaHost   string
}

func NewLogEvent(byteEvent []byte, actionParams map[string]string, ctx context.ExecutionContext) (*LogEvent, error) {
	var e elastic.SearchHit
	err := json.Unmarshal(byteEvent, &e)
	if err != nil {
		return nil, errors.New("unknown log event")
	}
	i := ctx.GetFiledByPath("index_pattern")
	log.Print(i)
	if i == "" {
		return nil, errors.New("execution fields \"index_pattern\" not found")
	}
	t := ctx.GetFiledByPath("index_pattern")
	if t == "" {
		return nil, errors.New("execution fields \"term\" not found")
	}
	f := ctx.GetFiledByPath("field_name")
	if f == "" {
		return nil, errors.New("execution fields \"field_name\" not found")
	}
	h, ok := actionParams["kibana_host"]
	if !ok {
		h = ""
	}

	return &LogEvent{
		event:        e,
		detail:       byteEvent,
		term:         t,
		indexPattern: i,
		fieldName:    f,
		kibanaHost:   h,
	}, nil
}

func (c *LogEvent) title() string {
	return "Log event"
}

func (c *LogEvent) statusColor() string {
	return "#CCCCCC"
}

func (c *LogEvent) authorName() string {
	return "LogEvent"
}

func (c *LogEvent) authorLink() string {
	return ""
}

func (c *LogEvent) fields() []*slack.Field {
	f := []*slack.Field{
		{Title: "Term", Value: c.term},
		{Title: "Index", Value: c.event.Index},
		{Title: "ID", Value: c.event.Id},
	}
	if c.kibanaHost != "" {
		f = append(f, &slack.Field{Title: "Kibana", Value: c.getKibanaLink()})
	}
	return f
}

func (c *LogEvent) getKibanaLink() string {
	return fmt.Sprintf("http://%s/app/kibana#/doc/%s/%s/doc?id=%s", c.kibanaHost, c.indexPattern, c.event.Index, c.event.Id)
}

func (c *LogEvent) Detail() json.RawMessage {
	return c.detail
}

func (c *LogEvent) Ignore(_ string) bool {
	return false
}

func (c *LogEvent) addMention(attachment slack.Attachment, _ map[string]string) slack.Attachment {
	return attachment
}
