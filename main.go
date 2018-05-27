package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func GetLiveBoard(ctx echo.Context) error {
	stn := ctx.Param("stn")

	b := getBoard(stn)

	if len(b) < 1 {
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

func Hello(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello, World")
}

func main() {
	addr := ":" + os.Getenv("PORT")
	e := echo.New()

	load_stations("./stations.json")

	e.GET("/", Hello)
	e.GET("/live/:stn", GetLiveBoard)
	e.GET("/service/:id", GetServiceDetails)
	e.GET("/station/:stn", GetStation)
	e.GET("/search/:name", GetSearchStations)

	e.Start(addr)
}
