package main

import "testing"

func TestSearch(t *testing.T) {
	load_stations("./stations.json")

	results := Stations.Search("Dalmuir", 10)

	if len(results) < 1 {
		t.Error("No results")
	}

	results = Stations.Search("n", 2)

	if len(results) > 2 {
		t.Error("Expected only 2 results")
	}
}
