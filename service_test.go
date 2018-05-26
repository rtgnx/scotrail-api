package main

import (
	"strings"
	"testing"
)

const SERVICE_HTML = `
<ul>
  <li>
    Glasgow Queen Street - Departed at 18:23    </li>
  <li>
    Dalmuir - Arrived on time at 18:41    </li>
  <li>
    Dumbarton Central - Arrived at 18:52    </li>
  <li>
    Helensburgh Upper - Arrived no report at 19:04    </li>
  <li>
    Garelochhead - Arrived at 19:20    </li>
  <li>
    Arrochar & Tarbet - Arrived no report at 19:36    </li>
  <li>
    Ardlui - Arrived no report at 19:50    </li>
  <li>
    Crianlarich - Arrived on time at 20:07      <br /> Train divides here with a portion going to Mallaig    <ul class="connection">
          <li>Crianlarich - Arrived on time at 20:07</li>
          <li>Upper Tyndrum - Arrived no report at 20:29</li>
          <li>Bridge of Orchy - Arrived no report at 20:44</li>
          <li>Rannoch - Arrived on time at 21:05</li>
          <li>Corrour - Arrived no report at 21:20</li>
          <li>Tulloch - Arrived no report at 21:35</li>
          <li>Roy Bridge - Arrived no report at 21:46</li>
          <li>Spean Bridge - Arrived no report at 21:52</li>
          <li>Fort William - Arrived on time at 22:06</li>
          <li>Banavie - Arrived on time at 22:20</li>
          <li>Corpach - Arrived on time at 22:25</li>
          <li>Loch Eil Outward Bound - Arriving on time at 22:30</li>
          <li>Locheilside - Arriving on time at 22:35</li>
          <li>Glenfinnan - Arriving on time at 22:46</li>
          <li>Lochailort - Arriving on time at 23:02</li>
          <li>Beasdale - Arriving on time at 23:11</li>
          <li>Arisaig - Arriving on time at 23:19</li>
          <li>Morar - Arriving on time at 23:28</li>
          <li>Mallaig - Arriving on time at 23:35</li>
        </ul>
    </li>
  <li>
    Tyndrum Lower - Arrived no report at 20:22    </li>
  <li>
    Dalmally - Arrived no report at 20:40    </li>
  <li>
    Loch Awe - Arrived no report at 20:45    </li>
  <li>
    Falls of Cruachan - Arrived no report at 20:50    </li>
  <li>
    Taynuilt - Arrived on time at 20:59    </li>
  <li>
    Connel Ferry - Arrived no report at 21:10    </li>
  <li>
    Oban - Arrived on time at 21:24    </li>
</ul>
`

func TestParseService(t *testing.T) {

	service := parseService(strings.NewReader(SERVICE_HTML))

	if len(service) < 1 {
		t.Errorf("Unable to read service details")
	}

	// TODO: test for connection service
}
