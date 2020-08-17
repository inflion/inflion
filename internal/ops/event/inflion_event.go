package event

import (
	"encoding/json"
)

type InflionEvent struct {
	project string
	body    map[string]interface{}
	rawBody json.RawMessage
}

func NewInflionEvent(project string, rawBody json.RawMessage) (*InflionEvent, error) {
	var body map[string]interface{}
	err := json.Unmarshal(rawBody, &body)
	if err != nil {
		return nil, err
	}
	return &InflionEvent{project: project, body: body, rawBody: rawBody}, nil
}

func (e InflionEvent) RawBody() json.RawMessage {
	return e.rawBody
}

func (e InflionEvent) Body() map[string]interface{} {
	return e.body
}

func (e InflionEvent) Project() string {
	return e.project
}

func (e InflionEvent) GetValue(key string) (interface{}, bool) {
	if value, ok := e.body[key]; ok {
		return value, true
	}
	return nil, false
}
