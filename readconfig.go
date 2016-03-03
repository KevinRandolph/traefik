package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type dict map[string]interface{}

func main() {
	config, _ := readMarathonConfig()
	for k, v := range *config {
		fmt.Printf("\n\ntype k: %T,   type v: %T\n", k, v)
	}
}

func readMarathonConfig() (*dict, error) {
	// test read of json file and turn it into a dict.
	file, _ := ioutil.ReadFile("marathon_config.json")
	marathonConfig := new(dict)
	json.Unmarshal(file, &marathonConfig)

	fmt.Printf("Test config file: %#+v\n", marathonConfig)
	return marathonConfig, nil
}
