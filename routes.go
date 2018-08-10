package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Routes map [name] Route
var Routes map[string]Route

func loadRoutes(path string) {
	b, _ := ioutil.ReadFile(path)

	err := json.Unmarshal(b, &Routes)

	if err != nil {
		log.Fatalln(err.Error())
	}
}
