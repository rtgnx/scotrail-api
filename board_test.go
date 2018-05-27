package main

import (
	"strings"
	"testing"
)

const TEST_BOARD = `
	<div id="nre-live-boards" data-crs="KNS">
		<h4>Departure board</h4>
			<table class="table responsive">
				<thead>
					<tr>
						<th>Platform</th>
						<th>Destination</th>
						<th>Departs</th>
						<th>Expected</th>
						<th>Origin</th>
						<th>Operator</th>
					</tr>
				</thead>
				<tbody>
									<tr class="service" data-id="8SlZWSDid2CaZh1W0UJYWA%3D%3D">
						<td data-label="Platform">1</td>
						<td data-label="Destination">Glasgow Central</td>
						<td data-label="Departs">19:34</td>
						<td data-label="Expected">On time</td>
						<td data-label="Origin">Barrhead</td>
						<td data-label="Operator">ScotRail</td>
					</tr>
									<tr class="service" data-id="F8AJ%23Z7tYsWG0lSJcEWzRg%3D%3D">
						<td data-label="Platform">2</td>
						<td data-label="Destination">Kilmarnock</td>
						<td data-label="Departs">19:45</td>
						<td data-label="Expected">On time</td>
						<td data-label="Origin">Glasgow Central</td>
						<td data-label="Operator">ScotRail</td>
					</tr>
									<tr class="service" data-id="%23wAabExJDdHHupQ%2BUU1JgA%3D%3D">
						<td data-label="Platform">1</td>
						<td data-label="Destination">Glasgow Central</td>
						<td data-label="Departs">19:59</td>
						<td data-label="Expected">On time</td>
						<td data-label="Origin">Kilmarnock</td>
						<td data-label="Operator">ScotRail</td>
					</tr>
									<tr class="service" data-id="hE2LKFuzXSUTiq9caPpEXA%3D%3D">
						<td data-label="Platform">-</td>
						<td data-label="Destination">Kilmarnock</td>
						<td data-label="Departs">20:45</td>
						<td data-label="Expected">On time</td>
						<td data-label="Origin">Glasgow Central</td>
						<td data-label="Operator">ScotRail</td>
					</tr>
									<tr class="service" data-id="kiGlAM4dMS9jLfKuIT4TIQ%3D%3D">
						<td data-label="Platform">-</td>
						<td data-label="Destination">Glasgow Central</td>
						<td data-label="Departs">20:59</td>
						<td data-label="Expected">On time</td>
						<td data-label="Origin">Kilmarnock</td>
						<td data-label="Operator">ScotRail</td>
					</tr>
								</tbody>
			</table>
		<h4>Arrival board</h4>
			<table class="table responsive">
				<thead>
					<tr>
						<th>Platform</th>
						<th>Origin</th>
						<th>Arrives</th>
						<th>Expected</th>
						<th>Destination</th>
						<th>Operator</th>
					</tr>
				</thead>
				<tbody>
									<tr class="service" data-id="8SlZWSDid2CaZh1W0UJYWA%3D%3D">
						<td data-label="Platform">1</td>
						<td data-label="Origin">Barrhead</td>
						<td data-label="Arrives">19:34</td>
						<td data-label="Expected">On time</td>
						<td data-label="Destination">Glasgow Central</td>
						<td data-label="Operator">ScotRail</td>
					</tr>
									<tr class="service" data-id="F8AJ%23Z7tYsWG0lSJcEWzRg%3D%3D">
						<td data-label="Platform">2</td>
						<td data-label="Origin">Glasgow Central</td>
						<td data-label="Arrives">19:45</td>
						<td data-label="Expected">On time</td>
						<td data-label="Destination">Kilmarnock</td>
						<td data-label="Operator">ScotRail</td>
					</tr>
									<tr class="service" data-id="%23wAabExJDdHHupQ%2BUU1JgA%3D%3D">
						<td data-label="Platform">1</td>
						<td data-label="Origin">Kilmarnock</td>
						<td data-label="Arrives">19:59</td>
						<td data-label="Expected">On time</td>
						<td data-label="Destination">Glasgow Central</td>
						<td data-label="Operator">ScotRail</td>
					</tr>
									<tr class="service" data-id="hE2LKFuzXSUTiq9caPpEXA%3D%3D">
						<td data-label="Platform">-</td>
						<td data-label="Origin">Glasgow Central</td>
						<td data-label="Arrives">20:45</td>
						<td data-label="Expected">On time</td>
						<td data-label="Destination">Kilmarnock</td>
						<td data-label="Operator">ScotRail</td>
					</tr>
									<tr class="service" data-id="kiGlAM4dMS9jLfKuIT4TIQ%3D%3D">
						<td data-label="Platform">-</td>
						<td data-label="Origin">Kilmarnock</td>
						<td data-label="Arrives">20:59</td>
						<td data-label="Expected">On time</td>
						<td data-label="Destination">Glasgow Central</td>
						<td data-label="Operator">ScotRail</td>
					</tr>
								</tbody>
			</table>
	</div>
`

func TestParseBoard(t *testing.T) {
	board := ParseBoard(strings.NewReader(TEST_BOARD))
	if len(board.Services) < 1 {
		t.Errorf("No data was extracted")
	}
}
