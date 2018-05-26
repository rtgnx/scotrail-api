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

func Hello(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello, World")
}

func main() {
	addr := ":" + os.Getenv("PORT")
	e := echo.New()

	e.GET("/", Hello)
	e.GET("/live/:stn", GetLiveBoard)

	e.Start(addr)
}
