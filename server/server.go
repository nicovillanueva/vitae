package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	"github.com/nicovillanueva/api-vitae/server/types"
)

var logger echo.Logger
var cvData types.CVData

// Start serves the server
func Start(d types.CVData) {
	cvData = d

	e := echo.New()

	s := newStats()
	e.Use(middleware.Recover())
	e.Use(s.midProcess)
	e.Use(midContactHeader)

	e.Add("GET", "/stats", s.HandleStats)
	setupRoutes(e)

	logger = e.Logger
	if l, ok := e.Logger.(*log.Logger); ok {
		l.SetHeader("${time_rfc3339} ${level}")
	}

	e.Logger.Fatal(e.Start(":80"))
}
