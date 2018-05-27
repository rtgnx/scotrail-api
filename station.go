package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"regexp"
)

type Station struct {
	Name      string  `json:"name"`
	Postcode  string  `json:"postcode"`
	Code      string  `json:"code"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type StationList map[string]Station

var Stations StationList = make(map[string]Station)

func (s *StationList) Search(name string, limit int) (l []Station) {
	for _, v := range *s {
		if ok, _ := regexp.MatchString(".*"+name+".*", v.Name); ok {
			l = append(l, v)
		}

		if len(l) == limit {
			break
		}
	}
	return
}

func load_stations(path string) {

	d, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatalln(err.Error())
	}

	var stations []Station

	err = json.Unmarshal(d, &stations)

	if err != nil {
		log.Fatalln(err.Error())
	}

	for _, s := range stations {
		Stations[s.Code] = s
	}
}
