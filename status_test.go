package main

import (
	"testing"
)

var route_names = []string{
	"route-1-strathclyde_s-glasgow_cathcart_neilston_newton",
	"route-2-strathclyde_s-glasgow_kilbride_barrhead_kilmarnock",
	"route-3-strathclyde_s-glasgow_paisley",
	"route-4-strathclyde_n-glasgow_milngavie_dalmuir_balloch_helensburgh",
	"route-5-strathclyde_n-glasgow_whifflet_motherwell_milngavie",
	"route-6-strathclyde_n-larkhall_glasgow",
	"route-7-strathclyde_n-helensburgh_milngavie_edinburgh_bathgate",
	"route-8-central-edinburgh_glasgow_falkirk",
	"route-9-central-glasgow_edinburgh_dunblane_stirling_alloa",
	"route-10-central-glasgow_shotts",
	"route-11-central-glasgow_lanark",
	"route-12-central-glasgow_falkirk_cumbernauld",
	"route-13-central-motherwell_cumbernauld",
	"route-14-central-glasgow_anniesland",
	"route-15-express-glasgow_aberdeen",
	"route-16-express-edinburgh_aberdeen",
	"route-17-express-glasgow_edinburgh_inverness",
	"route-18-highland-aberdeen_inverness",
	"route-19-highland-glasgow_oban_fort_william_mallaig",
	"route-20-highland-inverness_kyle",
	"route-21-highland-inverness_wick_thurso",
	"route-22-east_scotland-tweedbank_newcraighall_edinburgh",
	"route-23-east_scotland-fife_rosyth_dalgetty",
	"route-24-east_scotland-edinburgh_north_berwick_dunbar",
	"route-25-east_scotland-edinburgh_perth",
	"route-26-ayrshire-glasgow_ardrossan_ayr_largs",
	"route-27-ayrshire-glasgow_gourock_wemyss",
	"route-28-ayrshire-kilmarnock_ayr",
	"route-29-south_west-glasgow_stranraer_ayr_paisley",
	"route-31-south_west-glasgow_carlisle_newcastle",
}

var html_issue = ` <h2><span class="issues minor">Minor Disruption</span> Edinburgh - Aberdeen</h2><h3>1 services affected</h3> <ul class="services"><li> <a href="#service-13228074341A81G10216120618-73" id="service-link-13228074341A81G10216120618-73" class="collapsed" data-toggle="collapse" aria-expanded="false" aria-controls="service-13228074341A81G10216120618-73"> 17:36 Edinburgh to Aberdeen due 20:07 </a> <div class="collapse calling-stations" id="service-13228074341A81G10216120618-73"> <div class="service">17:36 Edinburgh to Aberdeen due 20:07 has been previously delayed, has been further delayed between Kirkcaldy and Dundee and is now 24 minutes late.<br/>This is due to a fault on a train in front of this one. <h4>Calling at stations:</h4> <div class="calling-station-result">Loading...</div> </div> </div> </li></ul>`

func TestParseRouteString(t *testing.T) {
	r := new(Route)

	for _, route := range route_names {
		parseRouteString(route, r)

		if r.RouteID == 0 {
			t.Error("Route ID wasn't parsed")
		}

		if len(r.Region) == 0 {
			t.Error("Route region wasn't parsed")
		}

		if len(r.Stations) == 0 {
			t.Error("Route stations weren't parsed")
		}

	}
}

func TestParseRouteDetails(t *testing.T) {

	r := new(Route)

	parseRouteDetails(html_issue, r)

	if len(r.Issue.Reason) < 1 {
		t.Errorf("No reason")
	}

	if len(r.Issue.AffectedServices) < 1 {
		t.Errorf("No affected services")
	}

	if len(r.Issue.Details) < 1 {
		t.Errorf("No details")
	}

}
