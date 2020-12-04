package context

import (
	"github.com/inflion/inflion/flow/event"
	"log"
	"testing"
)

func TestExecutionContext_GetFiledByPath(t *testing.T) {
	type fields struct {
		Fields map[string]interface{}
		Event  event.InflionEvent
	}
	type args struct {
		path string
	}
	nested := map[string]interface{}{
		"a": "b",
		"c": map[string]interface{}{
			"d": "e",
			"f": "g",
		},
		"z": 1.4567,
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "",
			fields: fields{
				Fields: nested,
				Event:  event.InflionEvent{},
			},
			args: args{
				path: "c.d",
			},
			want: "e",
		},
		{
			name: "",
			fields: fields{
				Fields: nested,
				Event:  event.InflionEvent{},
			},
			args: args{
				path: "c",
			},
			want: "",
		},
		{
			name: "",
			fields: fields{
				Fields: nested,
				Event:  event.InflionEvent{},
			},
			args: args{
				path: "z",
			},
			want: "1.4567",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			c := ExecutionContext{
				fields: tt.fields.Fields,
				event:  tt.fields.Event,
			}

			if got := c.GetFiledByPath(tt.args.path); got != tt.want {
				log.Print(got)
			}
		})
	}
}
