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

	doc.Find("ul").Each(func(i int, s *goquery.Selection) {
		stop := new(Stop)
		createStop(s, stop)
		service = append(service, *stop)
	})

	return
}

func createStop(s *goquery.Selection, stop *Stop) {
	text := strings.Split(s.Text(), "-")

	if len(text) < 2 {
		return
	}

	status := strings.Split(s.Text(), "at")

	if len(text) < 2 {
		return
	}

	t := ParseTime(status[1])

	stop = &Stop{
		Station: text[0], Status: status[0],
		Time: t, Connection: new(Service),
	}

	s.ChildrenFiltered("ul.connection").Each(func(i int, s *goquery.Selection) {
		cstop := new(Stop)
		createStop(s, cstop)
		*stop.Connection = append(*stop.Connection, *cstop)
	})
}
