package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/tidwall/gjson"
)

const ROUTE_STATUS_URL = `https://www.scotrail.co.uk/ajax/interactive_map/status`

type Issue struct {
	Reason           string   `json:"reason"`
	Details          []string `json:"details"`
	AffectedServices []string `json:"affected_services"`
}

type Route struct {
	RouteID  int      `json:"route_id"`
	Region   string   `json:"region"`
	Type     string   `json:"type"`
	Map      string   `json:"map"`
	Status   string   `json:"status"`
	Stations []string `json:"stations"`
	Issue    Issue
}

func GetRouteStatuses() (routes []Route) {
	r, err := http.Get(ROUTE_STATUS_URL)

	if err != nil {
		return
	}

	defer r.Body.Close()

	routes = parseRoutes(r.Body)
	return
}

func parseRoutes(r io.Reader) (routes []Route) {

	b, _ := ioutil.ReadAll(r)
	data := gjson.Get(string(b), "routes").Map()

	json.Unmarshal(b, data)

	for k, v := range data {
		r := new(Route)
		parseRouteString(k, r)

		r.Map, r.Status = v.Get("map").String(), v.Get("status").String()
		log.Printf("%v", r)

		if v.Get("html").Exists() {
			parseRouteDetails(v.Get("html").String(), r)
		}

		routes = append(routes, *r)
	}

	return
}

func parseRouteString(s string, r *Route) {
	re, _ := regexp.Compile(`route-([0-99]{1,2})-([a-zA-Z_]{1,})-([a-zA-Z_]{1,})`)

	res := re.FindAllStringSubmatch(s, -1)

	if len(res) != 1 || len(res[0]) != 4 {
		log.Fatalln("ERROR: Unable to parse route string")
	}

	id, _ := strconv.ParseInt(res[0][1], 10, 32)
	r.RouteID = int(id)
	r.Region = res[0][2]

	var stations = res[0][3]

	// Remove directions suffix
	for _, n := range []string{"_north", "_east", "_south", "_west"} {
		stations = strings.Replace(stations, n, "", -1)
	}

	r.Stations = strings.Split(stations, "_")
}

func parseRouteDetails(html string, r *Route) {

	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(html))
	issue := new(Issue)

	doc.Find("span.issues").Each(func(i int, s *goquery.Selection) {
		issue.Reason = s.Text()
	})

	doc.Find("ul.services").Children().Each(func(i int, s *goquery.Selection) {
		id := s.Children().First().AttrOr("href", "#")
		issue.AffectedServices = append(issue.AffectedServices, id[1:])
		s.Find("div" + id).Children().Each(func(i int, s *goquery.Selection) {
			issue.Details = append(issue.Details, s.Text())
		})
	})

	r.Issue = *issue
}
