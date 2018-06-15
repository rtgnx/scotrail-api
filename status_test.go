package main

import (
	"testing"
)

var html_issue = ` <h2><span class="issues minor">Minor Disruption</span> Edinburgh - Aberdeen</h2><h3>1 services affected</h3> <ul class="services"><li> <a href="#service-13228074341A81G10216120618-73" id="service-link-13228074341A81G10216120618-73" class="collapsed" data-toggle="collapse" aria-expanded="false" aria-controls="service-13228074341A81G10216120618-73"> 17:36 Edinburgh to Aberdeen due 20:07 </a> <div class="collapse calling-stations" id="service-13228074341A81G10216120618-73"> <div class="service">17:36 Edinburgh to Aberdeen due 20:07 has been previously delayed, has been further delayed between Kirkcaldy and Dundee and is now 24 minutes late.<br/>This is due to a fault on a train in front of this one. <h4>Calling at stations:</h4> <div class="calling-station-result">Loading...</div> </div> </div> </li></ul>`

func TestParseRouteDetails(t *testing.T) {

	r := new(Route)

	parseRouteDetails(html_issue, r)

	if len(r.Issue.Severity) < 1 {
		t.Errorf("No reason")
	}

}
