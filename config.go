package main

import (
	"encoding/json"
	"io/ioutil"
	"regexp"

	"gopkg.in/yaml.v2"
)

var rxYAML = regexp.MustCompile(`[.]ya?ml\z`)
var rxJSON = regexp.MustCompile(`[.]json\z`)

func unmarshalConfig(fileName string) []item {
	byteValue, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	var val []item
	if rxYAML.MatchString(fileName) {
		if err := yaml.Unmarshal(byteValue, &val); err != nil {
			panic(err)
		}
		return val
	} else if rxJSON.MatchString(fileName) {
		if err := json.Unmarshal(byteValue, &val); err != nil {
			panic(err)
		}
		return val
	} else {
		panic("Config " + fileName + " neither JSON nor YAML")
	}
}
