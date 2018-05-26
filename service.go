package main

import (
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const SERVICE_ENDPOINT = `https://www.scotrail.co.uk/nre/service-details/`

type Stop struct {
	Station    string
	Status     string
	Connection *Service
	Time       time.Time
}

type Service []Stop

func serviceDetails(service_id string) Service {
	res, err := http.Get(SERVICE_ENDPOINT + service_id)

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

	t := ParseTime(status[1])

	stop := Stop{
		Station: strings.Trim(text[0], " \n"), Status: strings.Trim(status[0], " \n"),
		Time: t, Connection: new(Service),
	}

	s.ChildrenFiltered("ul.connection").Each(func(i int, s *goquery.Selection) {
		*stop.Connection = append(*stop.Connection, createStop(s))
	})

	return stop
}
