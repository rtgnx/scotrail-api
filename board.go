package main

import (
	"io"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const BOARD_URL = `https://www.scotrail.co.uk/nre/live-boards/`

type Entry struct {
	ID          string    `json:"id"`
	Platform    int       `json:"platform"`
	Destination string    `json:"destination"`
	Departs     time.Time `json:"departs"`
	Arrives     time.Time `json:"arrives"`
	Expected    string    `json:"expected"`
	Origin      string    `json:"origin"`
	Operator    string    `json:"operator"`
}

type Board []Entry

func getBoard(station string) Board {
	res, err := http.Get(BOARD_URL + station)
	if err != nil {
		log.Fatalln(err.Error())
	}

	if res.StatusCode != 200 {
		log.Fatalln(res.Status)
	}

	defer res.Body.Close()

	return ParseBoard(res.Body)
}

func ParseBoard(r io.Reader) (b Board) {
	doc, _ := goquery.NewDocumentFromReader(r)

	doc.Find("tr.service").Each(func(i int, s *goquery.Selection) {
		var e Entry

		if id, ok := s.Attr("data-id"); ok {
			e.ID = id
		}

		s.ChildrenFiltered("td").Each(func(i int, s *goquery.Selection) {
			switch s.AttrOr("data-label", "") {
			case "Expected", "Origin", "Destination", "Operator":
				label := s.AttrOr("data-label", "")
				reflect.ValueOf(&e).Elem().FieldByName(label).SetString(s.Text())
			case "Platform":
				e.Platform = platform(s.Text())
			case "Arrives", "Departs":
				t := ParseTime(s.Text())
				label := s.AttrOr("data-label", "")
				reflect.ValueOf(&e).Elem().FieldByName(label).Set(reflect.ValueOf(t))
			}
		})

		b = append(b, e)
	})
	return
}

func platform(p string) int {
	i, _ := strconv.ParseInt(p, 10, 32)
	return int(i)
}
