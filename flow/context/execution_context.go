package context

import (
	"encoding/json"
	"fmt"
	event2 "github.com/inflion/inflion/flow/event"
	"github.com/jeremywohl/flatten"
)

type ExecutionContext struct {
	fields map[string]interface{}
	event  event2.InflionEvent
}

func NewExecutionContext() ExecutionContext {
	return ExecutionContext{
		fields: map[string]interface{}{},
		event:  event2.InflionEvent{},
	}
}

func NewExecutionContextWithEvent(event *event2.InflionEvent) ExecutionContext {
	return ExecutionContext{
		fields: map[string]interface{}{},
		event:  *event,
	}
}

func (c ExecutionContext) Fields() map[string]interface{} {
	return c.fields
}

func (c ExecutionContext) Event() event2.InflionEvent {
	return c.event
}

func (c ExecutionContext) AddField(key string, field interface{}) {
	c.fields[key] = field
}

func (c ExecutionContext) GetFiledByPath(path string) string {
	bytes, err := json.Marshal(c.fields)
	if err != nil {
		return ""
	}

	var m map[string]interface{}
	err = json.Unmarshal(bytes, &m)
	if err != nil {
		return ""
	}

	f, err := flatten.Flatten(m, "", flatten.DotStyle)
	if err != nil {
		return ""
	}

	v, ok := f[path]
	if !ok {
		return ""
	}

	return fmt.Sprintf("%v", v)
}
