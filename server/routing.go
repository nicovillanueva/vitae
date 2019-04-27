package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

// APIVersion is a specific API version with all it's handlers
// Call `.apply()` to enable the routes in the server
type APIVersion struct {
	Number    int
	Enabled   bool
	Handlings []routeDefinition
}

type routeDefinition struct {
	Method  string
	Path    string
	Handler func(echo.Context) error
}

func (a *APIVersion) apply(e *echo.Echo) {
	if !a.Enabled {
		return
	}
	g := e.Group(fmt.Sprintf("/v%d", a.Number))
	for _, r := range a.Handlings {
		g.Add(r.Method, r.Path, r.Handler)
	}
	g.Add("GET", "/", func(c echo.Context) error {
		return c.JSON(200, struct {
			RoutesAvailable []*echo.Route `json:"routes_available"`
		}{
			e.Routes(),
		})
	})
}

func setupRoutes(e *echo.Echo) {
	// Root
	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, struct {
			Status string `json:"status"`
			Usage  string `json:"usage"`
			Stats  string `json:"stats"`
		}{
			Status: "okay",
			Usage:  "Send a GET request to each version's root to see available routes. E.g., `GET /v1/`",
			Stats:  "Stats are available on `GET /stats`",
		})
	})

	a := APIVersion{
		Number:  1,
		Enabled: true,
		Handlings: []routeDefinition{
			{"GET", "/about", handlePersonalStatement},
			{"GET", "/download", handleDownloadCV},
			{"GET", "/skills", handleKeySkills},
			{"GET", "/skills/:sid", handleKeySkill},
			{"GET", "/achievements", handleAchievements},
			{"GET", "/achievements/:aid", handleAchievement},
			{"GET", "/work", handleWorkHistory},
			{"GET", "/work/:wid", handleWorkPlace},
			{"GET", "/education", handleEducation},
			{"GET", "/education/:eid", handleEducationPlace},
			{"GET", "/hobbies", handleHobbies},
			{"GET", "/references", handleReferences},
		},
	}
	a.apply(e)
}
