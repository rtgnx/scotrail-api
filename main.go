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

func GetLiveBoard(ctx echo.Context) error {
	stn := ctx.Param("stn")

	b := getBoard(stn)

	if strings.Compare(b.Station.Code, stn) != 0 {
		return ctx.String(http.StatusNotFound, "Invalid station code")
	}

	return ctx.JSON(http.StatusOK, b)
}

func GetServiceDetails(ctx echo.Context) error {
	id := ctx.Param("id")

	service := serviceDetails(id)

	if len(service) < 1 {
		ctx.String(http.StatusNotAcceptable, "Invalid service id")
	}

	return ctx.JSON(http.StatusOK, service)
}

func GetStation(ctx echo.Context) error {
	name := ctx.Param("stn")

	if v, ok := Stations[name]; ok {
		return ctx.JSON(http.StatusOK, v)
	}

	return ctx.String(http.StatusNotFound, "No such station")
}

func GetSearchStations(ctx echo.Context) error {
	name := ctx.Param("name")
	return ctx.JSON(http.StatusOK, Stations.Search(name, 10))
}

func GetStatus(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, GetRouteStatuses())
}

func GetNearest(ctx echo.Context) error {

	lat, _ := strconv.ParseFloat(ctx.Param("lat"), 64)
	lon, _ := strconv.ParseFloat(ctx.Param("lon"), 64)

	dist, st := Stations.Nearest(Coordinate{lat: lat, lon: lon})
	return ctx.JSON(
		http.StatusOK, map[string]interface{}{"station": st, "distance": dist},
	)
}

func GetStations(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, Stations)
}

func Hello(ctx echo.Context) error {
	d, _ := ioutil.ReadFile("./index.html")
	return ctx.HTMLBlob(http.StatusOK, d)
}

func main() {
	addr := ":" + os.Getenv("PORT")
	e := echo.New()

	e.Use(middleware.CORS())

	load_stations()
	load_routes("./routes.json")

	e.GET("/", Hello)
	e.GET("/status", GetStatus)
	e.GET("/live/:stn", GetLiveBoard)
	e.GET("/service/:id", GetServiceDetails)
	e.GET("/stations", GetStations)
	e.GET("/station/:stn", GetStation)
	e.GET("/search/:name", GetSearchStations)
	e.GET("/nearest/:lat/:lon", GetNearest)
	e.Static("/static", "./static")

	e.Logger.Debug(e.Start(addr))
}
