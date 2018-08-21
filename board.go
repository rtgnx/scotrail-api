package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"reflect"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const BOARD_URL = `https://www.scotrail.co.uk/nre/live-boards/%s/lazy`

type Entry struct {
	ID          string `json:"id"`
	Platform    string `json:"platform"`
	Destination string `json:"destination"`
	Departs     string `json:"departs"`
	Arrives     string `json:"arrives"`
	Expected    string `json:"expected"`
	Origin      string `json:"origin"`
	Operator    string `json:"operator"`
}

type Board struct {
	Station  Station `json:"station"`
	Services []Entry `json:"services"`
}

func (b *Board) HasStation(id string) (bool, int) {
	for i, v := range b.Services {
		if strings.Compare(v.ID, id) == 0 {
			return true, i
		}
	}

	return false, 0
}

func getBoard(station string) Board {
	res, err := http.Get(fmt.Sprintf(BOARD_URL, station))
	if err != nil {
		log.Fatalln(err.Error())
	}

	if res.StatusCode != 200 {
		log.Fatalln(res.Status)
	}

	defer res.Body.Close()

	b := ParseBoard(res.Body)
	b.Station, _ = Stations[station]

	return b
}

func ParseBoard(r io.Reader) (b Board) {
	doc, _ := goquery.NewDocumentFromReader(r)

	doc.Find("tr.service").Each(func(i int, s *goquery.Selection) {
		var e *Entry = new(Entry)

		id, ok := s.Attr("data-id")

		if ok {
			if ok, idx := b.HasStation(id); ok {
				e = &(b.Services)[idx]
			}
		}

		s.ChildrenFiltered("td").Each(func(i int, s *goquery.Selection) {
			switch s.AttrOr("data-label", "") {
			case "Expected", "Origin", "Destination", "Operator":
				label := s.AttrOr("data-label", "")
				reflect.ValueOf(e).Elem().FieldByName(label).SetString(s.Text())
			case "Platform":
				e.Platform = s.Text()
			case "Arrives", "Departs":
				t := s.Text()
				label := s.AttrOr("data-label", "")
				reflect.ValueOf(e).Elem().FieldByName(label).Set(reflect.ValueOf(t))
			}
		})

		if len(e.ID) < 1 {
			e.ID = id
			b.Services = append(b.Services, *e)
		}
	})
	return
}
