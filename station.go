package main

import (
	"io/ioutil"
	"math"
	"net/http"
	"regexp"

	"github.com/tidwall/gjson"
)

const (
	STATION_LIST_URL = `https://www.scotrail.co.uk/cache/trainline_stations/trainline?_=1530115581789`
)

type Coordinate struct {
	lon float64 `json:"longitude"`
	lat float64 `json:"latitude"`
}

type Station struct {
	Name      string  `json:"name"`
	Code      string  `json:"code"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type StationList map[string]Station

var Stations StationList = make(map[string]Station)

func (s *StationList) Nearest(location Coordinate) (dist int, stn Station) {
	for _, st := range *s {
		if st.Latitude == 0.0 || st.Longitude == 0.0 {
			continue
		}

		d := distance(Coordinate{lon: st.Longitude, lat: st.Latitude}, location)

		if dist == 0 || dist > d {
			dist = d
			stn = st
		}
	}

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

func load_stations() {

	res, _ := http.Get(STATION_LIST_URL)
	b, _ := ioutil.ReadAll(res.Body)

	defer res.Body.Close()
	stations := gjson.Get(string(b), "stations").Map()

	for _, v := range stations {
		Stations[v.Get("crs").String()] = Station{
			Name:      v.Get("name").String(),
			Code:      v.Get("crs").String(),
			Longitude: v.Get("lon").Float(),
			Latitude:  v.Get("lat").Float(),
		}
	}
}

// Conversion from degress to radians
func deg2rad(x float64) float64 {
	return x * float64((math.Pi / 180))
}

// Returns distance between two cordinates
func distance(a, b Coordinate) int {
	var R = 6371e3 // metres
	var φ1 = deg2rad(a.lat)
	var φ2 = deg2rad(b.lat)

	var Δφ = deg2rad(b.lat - a.lat)
	var Δλ = deg2rad(b.lon - a.lon)

	var an = math.Sin(Δφ/2) * math.Sin(Δφ/2)
	an += math.Cos(φ1) * math.Cos(φ2) * math.Sin(Δλ/2) * math.Sin(Δλ/2)

	var c = 2 * math.Atan2(math.Sqrt(an), math.Sqrt(1-an))

	var d = R * c

	return int(d)
}
