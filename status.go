package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/bradfitz/slice"
	"github.com/tidwall/gjson"
)

const ROUTE_STATUS_URL = `https://www.scotrail.co.uk/ajax/interactive_map/status`

type ServiceStatus struct {
	ID          string    `json:"id"`
	Time        time.Time `json:"time"`
	Due         time.Time `json:"due"`
	Origin      string    `json:"origin"`
	Destination string    `json:"destination"`
	Status      string    `json:"status"`
}

type Issue struct {
	Severity string          `json:"severity"`
	Details  []ServiceStatus `json:"details"`
}

type Route struct {
	RouteID  int      `json:"route_id"`
	Region   string   `json:"region"`
	Type     string   `json:"type"`
	Map      string   `json:"map"`
	Status   string   `json:"status"`
	Stations []string `json:"stations"`
	Issue    Issue    `json:"issue"`
}

func GetRouteStatuses() (routes []Route) {
	r, err := http.Get(ROUTE_STATUS_URL)

	if err != nil {
		return
	}

	defer r.Body.Close()

	routes = parseRoutes(r.Body)

	slice.Sort(routes[:], func(i, j int) bool {
		return routes[i].RouteID < routes[j].RouteID
	})

	return
}

func parseRoutes(r io.Reader) (routes []Route) {

	b, _ := ioutil.ReadAll(r)
	data := gjson.Get(string(b), "routes").Map()

	for k, v := range data {

		r, ok := Routes[k]

		if !ok {
			log.Fatalf("Route: %s was not found in routes.json", k)
		}

		r.Map, r.Status = v.Get("map").String(), v.Get("status").String()

		if v.Get("html").Exists() {
			parseRouteDetails(v.Get("html").String(), &r)
		}

		routes = append(routes, r)
	}

	return
}

func parseRouteDetails(html string, r *Route) {

	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(html))
	issue := new(Issue)

	re, err := regexp.Compile(`^([0-9]{2}):([0-9]{2})\s([a-zA-Z\s]{1,})\sto\s([a-zA-Z\s]{1,})\sdue\s([0-9]{2}):([0-9]{2})\s(will be|has been|was)\s(cancelled|reinstated|started)`)

	if err != nil {
		log.Fatalln(err.Error())
	}

	doc.Find("span.issues").Each(func(i int, s *goquery.Selection) {
		issue.Severity = s.Text()

		doc.Find("ul.services").Children().Each(func(i int, s *goquery.Selection) {
			id := s.Children().First().AttrOr("href", "#")

			s.Find("div" + id).Children().Each(func(i int, s *goquery.Selection) {
				matches := re.FindAllStringSubmatch(s.Text(), -1)

				if len(matches) == 0 {
					return
				}

				ss := ServiceStatus{
					ID: id[1:], Origin: matches[0][3], Destination: matches[0][4],
					Status: matches[0][7] + " " + matches[0][8],
				}

				issue.Details = append(issue.Details, ss)

			})
		})
	})

	r.Issue = *issue
}
