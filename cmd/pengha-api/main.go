package main

import (
	"flag"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"yonghochoi.com/depthon-2019/cmd/pengha-api/conf"
)

const version = "0.0.1"

func main() {
	configPath := flag.String("config", "./config.yml", "Input config file path")
	flag.Parse()
	conf.SetConfigFilePath(*configPath)

	InitEmoticons() // test data
	InitMoods()

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "success")
	})

	e.GET("/version", func(c echo.Context) error {
		return c.String(http.StatusOK, version)
	})

	e.POST("/join", Join)
	e.POST("/login", Login)
	e.GET("/api/users", GetUsers)
	e.GET("/api/emoticons", GetEmoticonDatas)
	e.GET("/api/moods", GetMoods)
	e.POST("/api/moods", CreateMood)

	e.Logger.Fatal(e.Start(":8000"))
}
