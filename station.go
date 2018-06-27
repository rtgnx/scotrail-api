package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math"
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

// Return map distance(meters) => Station
func (s *StationList) Nearest() (st map[int]Station) {
	return
}

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

type Cordinate struct {
	long, lat float64
}

// Conversion from degress to radians
func deg2rad(x float64) float64 {
	return x * float64((math.Pi / 180))
}

// Returns distance between two cordinates
func distance(a, b Cordinate) int {
	var R = 6371e3 // metres
	var φ1 = deg2rad(a.lat)
	var φ2 = deg2rad(b.lat)
	var Δφ = deg2rad(b.lat - a.lat)
	var Δλ = deg2rad(b.long - a.long)

	var an = math.Sin(Δφ/2)*math.Sin(Δφ/2) + math.Cos(φ1)*math.Cos(φ2)*math.Sin(Δλ/2)*math.Sin(Δλ/2)

	var c = 2 * math.Atan2(math.Sqrt(an), math.Sqrt(1-an))

	var d = R * c
	// Convert to meters and cast to int
	return int(d / 1000)
}
