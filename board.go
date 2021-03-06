package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// BoardURL link to html version of live board
const BoardURL = `https://www.scotrail.co.uk/nre/live-boards/%s/lazy`

// Entry on live board
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

// Board type
type Board struct {
	Station   Station `json:"station"`
	Services  []Entry `json:"services"`
	timestamp time.Time
}

// Boards map of recently fetched boards
// used for caching requests
var Boards = make(map[string]Board)

func (b *Board) hasStation(id string) (bool, int) {
	for i, v := range b.Services {
		if strings.Compare(v.ID, id) == 0 {
			return true, i
		}
	}

	return false, 0
}

func getBoard(station string) Board {
	b, ok := Boards[station]

	if ok && time.Now().Before(b.timestamp.Add(time.Second*30)) {
		return b
	}

	res, err := http.Get(fmt.Sprintf(BoardURL, station))

	if err != nil {
		log.Fatalln(err.Error())
	}

	if res.StatusCode != 200 {
		log.Fatalln(res.Status)
	}

	defer res.Body.Close()

	b = parseBoard(res.Body)
	b.Station, _ = Stations[station]

	Boards[station] = b
	return b
}

func parseBoard(r io.Reader) (b Board) {
	doc, _ := goquery.NewDocumentFromReader(r)

	b.Services = make([]Entry, 0)

	doc.Find("tr.service").Each(func(i int, s *goquery.Selection) {
		var e = new(Entry)

		id, ok := s.Attr("data-id")

		if ok {
			if ok, idx := b.hasStation(id); ok {
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

	b.timestamp = time.Now()
	return
}
