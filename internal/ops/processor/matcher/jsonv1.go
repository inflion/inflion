package matcher

import (
	"encoding/json"
)

type MetadataJson struct {
	Metadata struct {
		Format struct {
			Version int `json:"version"`
		} `json:"Format"`
	} `json:"Metadata"`
}

func UnmarshalV1(v1FormattedJson []byte) (*RuleRootJsonV1, error) {
	v1 := RuleRootJsonV1{}
	err := json.Unmarshal(v1FormattedJson, &v1)
	if err != nil {
		return nil, err
	}
	return &v1, nil
}

type RuleRootJsonV1 struct {
	Metadata MetadataJson `json:"metadata"`
	Body     struct {
		Name       string            `json:"name"`
		Target     string            `json:"target"`
		Conditions []ConditionJsonV1 `json:"conditions"`
	} `json:"body"`
}

type ConditionJsonV1 struct {
	TargetAttr string `json:"target_attr"`
	MatchType  string `json:"match_type"`
	MatchValue string `json:"match_value"`
}

func (r RuleRootJsonV1) mustConvertConditions() []Condition {
	var c []Condition
	for _, jc := range r.Body.Conditions {
		c = append(c, Condition{
			TargetAttr: jc.TargetAttr,
			MatchType:  jc.MatchType,
			MatchValue: jc.MatchValue,
		})
	}
	return c
}
