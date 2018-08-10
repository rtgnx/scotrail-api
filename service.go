package main

import (
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// ServiceURL link to service details
const ServiceURL = `https://www.scotrail.co.uk/nre/service-details/`

// Stop - train stop
type Stop struct {
	Station    string
	Status     string
	Connection *Service
	Time       string
}

// Service type
type Service []Stop

func serviceDetails(serviceID string) Service {
	res, err := http.Get(ServiceURL + serviceID)

	if err != nil {
		log.Fatalln(err.Error())
	}

	defer res.Body.Close()

	return parseService(res.Body)
}

func parseService(r io.Reader) (service Service) {
	doc, _ := goquery.NewDocumentFromReader(r)

	doc.Find("ul").ChildrenFiltered("li").Each(func(i int, s *goquery.Selection) {
		service = append(service, createStop(s))
	})

	return
}

func createStop(s *goquery.Selection) Stop {
	text := strings.Split(s.Text(), "-")
	status := strings.Split(text[1], "at")

	t := status[1]

	stop := Stop{
		Station: strings.Trim(text[0], " \n"), Status: strings.Trim(status[0], " \n"),
		Time: t, Connection: new(Service),
	}

	s.ChildrenFiltered("ul.connection").Each(func(i int, s *goquery.Selection) {
		*stop.Connection = append(*stop.Connection, createStop(s))
	})

	return stop
}
