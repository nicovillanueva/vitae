package server

import (
	"github.com/labstack/echo/v4"
)

var routes = apiVersion{
	Prefix:  "/api",
	Enabled: true,
	Handlings: []routeDefinition{
		{"GET", "/", handleRootToSwag},
		{"GET", "/about", handlePersonalStatement},
		{"GET", "/pdf", handleDownloadCV},
		{"GET", "/skills", handleKeySkills},
		{"GET", "/skills/:id", handleKeySkill},
		{"GET", "/achievements", handleAchievements},
		{"GET", "/achievements/:id", handleAchievement},
		{"GET", "/work", handleWorkHistory},
		{"GET", "/work/:id", handleWorkPlace},
		{"GET", "/education", handleEducation},
		{"GET", "/education/:id", handleEducationPlace},
		{"GET", "/hobbies", handleHobbies},
		{"GET", "/references", handleReferences},
		{"GET", "/references/:id", handleReference},
	},
}

type apiVersion struct {
	Prefix    string
	Enabled   bool
	Handlings []routeDefinition
}

type routeDefinition struct {
	Method  string
	Path    string
	Handler func(echo.Context) error
}

func (a *apiVersion) apply(e *echo.Echo) {
	if !a.Enabled {
		return
	}
	if a.Prefix != "" {
		g := e.Group(a.Prefix)
		for _, r := range a.Handlings {
			g.Add(r.Method, r.Path, r.Handler)
		}
		return
	}
	for _, r := range a.Handlings {
		e.Add(r.Method, r.Path, r.Handler)
	}
}
