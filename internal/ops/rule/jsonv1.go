package rule

import (
	"encoding/json"
	"fmt"
)

type MetadataJson struct {
	Metadata struct {
		Format struct {
			Version int `json:"version"`
		} `json:"Format"`
	} `json:"Metadata"`
}

func Unmarshal(rawJson []byte) (Rule, error) {
	m := MetadataJson{}
	err := json.Unmarshal(rawJson, &m)
	if err != nil {
		return Rule{}, err
	}

	if m.Metadata.Format.Version == 1 {
		v1, err := UnmarshalV1(rawJson)
		if err != nil {
			return Rule{}, err
		}

		return Rule{
			RuleName:   v1.Body.Name,
			Target:     v1.Body.Target,
			Conditions: Conditions{Conditions: v1.mustConvertConditions()},
		}, nil
	}

	return Rule{}, fmt.Errorf("json version %d not supported", m.Metadata.Format.Version)
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
