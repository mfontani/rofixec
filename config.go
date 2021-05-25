package main

import (
	"encoding/json"
	"io/ioutil"
)

func unmarshalConfig(fileName string) []item {
	byteValue, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	var val []item
	if err := json.Unmarshal(byteValue, &val); err != nil {
		panic(err)
	}
	return val
}
