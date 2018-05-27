package main

import (
	"io"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"
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

type Board struct {
	Station  Station
	Services []Entry
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
	res, err := http.Get(BOARD_URL + station)
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
				e.Platform = platform(s.Text())
			case "Arrives", "Departs":
				t := ParseTime(s.Text())
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

func platform(p string) int {
	i, _ := strconv.ParseInt(p, 10, 32)
	return int(i)
}
