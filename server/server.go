package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	docs "github.com/nicovillanueva/vitae/server/docs" // generated swagger docs
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/nicovillanueva/vitae/server/types"
)

// Version is the current server version
const Version = "0.2"

var logger echo.Logger
var cvData types.CVData

// Start serves the server
// @title Vitae
// @description My career's documentation
// @host localhost:8080
// @BasePath /api/
func Start(d types.CVData) {
	docs.SwaggerInfo.Version = Version

	cvData = d
	e := echo.New()

	logger = e.Logger
	if l, ok := e.Logger.(*log.Logger); ok {
		l.SetLevel(log.DEBUG)
		l.SetHeader("${time_rfc3339} ${level}")
	}

	s := newStats()
	e.Use(middleware.Recover())
	e.Use(s.midProcess)
	e.Use(midContactHeader)

	// e.Static("/", "static") // Static page is DISABLED
	e.Add("GET", "/api", handleRootToSwag) // not working?
	e.Add("GET", "/api/stats", s.HandleStats)
	e.Add("GET", "/swagger/*", echoSwagger.WrapHandler)
	routes.apply(e)

	e.Logger.Fatal(e.Start(":8080"))
}
