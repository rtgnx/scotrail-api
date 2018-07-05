package main

import (
	"strings"
	"testing"
)

func TestSearch(t *testing.T) {
	load_stations()

	results := Stations.Search("Dalmuir", 10)

	if len(results) < 1 {
		t.Error("No results")
	}

	results = Stations.Search("n", 2)

	if len(results) > 2 {
		t.Error("Expected only 2 results")
	}
}

func TestNearest(t *testing.T) {
	load_stations()

	var c = Coordinate{lat: 55.859618, lon: -4.257926}
	d, st := Stations.Nearest(c)

	if strings.Compare(st.Code, "GLC") != 0 {
		t.Errorf("Nearest station is %s within %d meters", st.Name, d)
	}
}
