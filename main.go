package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func getLiveBoard(ctx echo.Context) error {
	stn := ctx.Param("stn")

	b := getBoard(stn)

	if strings.Compare(b.Station.Code, stn) != 0 {
		return ctx.String(http.StatusNotFound, "Invalid station code")
	}

	return ctx.JSON(http.StatusOK, b)
}

func getServiceDetails(ctx echo.Context) error {
	id := ctx.Param("id")

	service := serviceDetails(id)

	if len(service) < 1 {
		ctx.String(http.StatusNotAcceptable, "Invalid service id")
	}

	return ctx.JSON(http.StatusOK, service)
}

func getStation(ctx echo.Context) error {
	name := ctx.Param("stn")

	if v, ok := Stations[name]; ok {
		return ctx.JSON(http.StatusOK, v)
	}

	return ctx.String(http.StatusNotFound, "No such station")
}

func getSearchStations(ctx echo.Context) error {
	name := ctx.Param("name")
	return ctx.JSON(http.StatusOK, Stations.Search(name, 10))
}

func getStatus(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, getRouteStatuses())
}

func getNearest(ctx echo.Context) error {

	lat, _ := strconv.ParseFloat(ctx.Param("lat"), 64)
	lon, _ := strconv.ParseFloat(ctx.Param("lon"), 64)

	dist, st := Stations.Nearest(Coordinate{lat: lat, lon: lon})
	return ctx.JSON(
		http.StatusOK, map[string]interface{}{"station": st, "distance": dist},
	)
}

func getStations(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, Stations)
}

func index(ctx echo.Context) error {
	d, _ := ioutil.ReadFile("./index.html")
	return ctx.HTMLBlob(http.StatusOK, d)
}

func main() {
	addr := ":" + os.Getenv("PORT")
	e := echo.New()

	e.Use(middleware.CORS())

	loadStations()
	loadRoutes("./routes.json")

	e.GET("/", index)
	e.GET("/status", getStatus)
	e.GET("/live/:stn", getLiveBoard)
	e.GET("/service/:id", getServiceDetails)
	e.GET("/stations", getStations)
	e.GET("/station/:stn", getStation)
	e.GET("/search/:name", getSearchStations)
	e.GET("/nearest/:lat/:lon", getNearest)
	e.Static("/static", "./static")

	e.Logger.Debug(e.Start(addr))
}
