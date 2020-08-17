package processor

import "github.com/inflion/inflion/internal/ops/flow"

type ByteFlowReader struct {
	body []byte
}

func (b ByteFlowReader) Read() (flow.Flow, error) {
	recipe, err := flow.Unmarshal(b.body)
	if err != nil {
		return flow.Flow{}, err
	}
	return recipe, nil
}
