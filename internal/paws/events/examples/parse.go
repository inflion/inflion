package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	bytes, err := ioutil.ReadFile("trusted_advisor_cloudwatch_event.json")
	if err != nil {
		panic(err)
	}
	var d map[string]interface{}
	json.Unmarshal(bytes, &d)

	for k1, v1 := range d {
		if v2, ok := v1.(string); ok {
			fmt.Printf("%s: %s\n", k1, v2)
		} else if v3, ok := v1.(map[string]interface{});ok {
			for k, val := range v3 {
				fmt.Printf("%s: %s\n", k, val)
			}
		}
	}
}