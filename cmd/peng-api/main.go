package main

import (
	"github.com/labstack/echo"
	"net/http"
)

const version = "0.0.1"

func main() {
	InitSample() // test data

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "success")
	})

	e.GET("/version", func(c echo.Context) error {
		return c.String(http.StatusOK, version)
	})

	e.POST("/join", Join)
	e.GET("/moods", GetMoods)
	e.GET("/data/emoticons", GetEmoticonDatas)

	e.Logger.Fatal(e.Start(":8000"))
}
