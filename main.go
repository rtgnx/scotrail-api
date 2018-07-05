package main

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func GetLiveBoard(ctx echo.Context) error {
	stn := ctx.Param("stn")

	b := getBoard(stn)

	if len(b.Services) < 1 {
		return ctx.String(http.StatusNotAcceptable, "Invalid station name")
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

func PostNearest(ctx echo.Context) error {

	pt := new(Cordinate)

	if err := ctx.Bind(pt); err != nil {
		return ctx.String(http.StatusOK, err.Error())
	}

	_, st := Stations.Nearest(*pt)
	return ctx.JSON(http.StatusOK, st)
}

func Hello(ctx echo.Context) error {
	d, _ := ioutil.ReadFile("./index.html")
	return ctx.HTMLBlob(http.StatusOK, d)
}

func main() {
	addr := ":" + os.Getenv("PORT")
	e := echo.New()

	load_stations()
	load_routes("./routes.json")

	e.GET("/", Hello)
	e.GET("/status", GetStatus)
	e.GET("/live/:stn", GetLiveBoard)
	e.GET("/service/:id", GetServiceDetails)
	e.GET("/station/:stn", GetStation)
	e.GET("/search/:name", GetSearchStations)
	e.POST("/nearest", PostNearest)
	e.Static("/static", "./static")

	e.Logger.Debug(e.Start(addr))
}
