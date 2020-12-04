package flow

type ByteFlowReader struct {
	body []byte
}

func (b ByteFlowReader) Read() (Flow, error) {
	recipe, err := Unmarshal(b.body)
	if err != nil {
		return Flow{}, err
	}
	return recipe, nil
}
