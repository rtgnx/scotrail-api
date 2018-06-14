package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var Routes map[string]Route

func load_routes(path string) {
	b, _ := ioutil.ReadFile(path)

	err := json.Unmarshal(b, &Routes)

	if err != nil {
		log.Fatalln(err.Error())
	}
}
