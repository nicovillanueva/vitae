package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type response struct {
	Content interface{} `json:"content"`
}

type responseError struct {
	Reason string `json:"reason"`
}

func handlePersonalStatement(c echo.Context) error {
	return c.JSON(http.StatusOK, response{cvData.Statement})
}

func handleKeySkills(c echo.Context) error {
	return c.JSON(http.StatusOK, response{cvData.Skills})
}

func handleKeySkill(c echo.Context) error {
	i, err := strconv.Atoi(c.Param("sid"))
	if err != nil {
		return err
	}
	if i < 0 || i > len(cvData.Skills) {
		return c.JSON(http.StatusNotFound, errSkillNotFound)
	}
	return c.JSON(http.StatusOK, response{cvData.Skills[i]})
}

func handleAchievements(c echo.Context) error {
	return c.JSON(http.StatusOK, response{cvData.Achievements})
}
func handleAchievement(c echo.Context) error {
	i, err := strconv.Atoi(c.Param("aid"))
	if err != nil {
		return err
	}
	if i < 0 || i > len(cvData.Achievements)-1 {
		return c.JSON(http.StatusNotFound, errAchievementNotFound)
	}
	return c.JSON(http.StatusOK, response{cvData.Achievements[i]})
}

func handleWorkHistory(c echo.Context) error {
	return c.JSON(http.StatusOK, response{cvData.WorkHistory})
}

func handleWorkPlace(c echo.Context) error {
	i, err := strconv.Atoi(c.Param("wid"))
	if err != nil {
		return err
	}
	if i < 0 || i > len(cvData.WorkHistory)-1 {
		return c.JSON(http.StatusNotFound, errJobNotFound)
	}
	return c.JSON(http.StatusOK, response{cvData.WorkHistory[i]})
}

func handleEducation(c echo.Context) error {
	return c.JSON(http.StatusOK, response{cvData.Education})
}

func handleEducationPlace(c echo.Context) error {
	i, err := strconv.Atoi(c.Param("eid"))
	if err != nil {
		return err
	}
	if i < 0 || i > len(cvData.Education)-1 {
		return c.JSON(http.StatusNotFound, errStudiesNotFound)
	}
	return c.JSON(http.StatusOK, response{cvData.Education[i]})
}

func handleHobbies(c echo.Context) error {
	return c.JSON(http.StatusOK, response{cvData.Hobbies})
}

func handleReferences(c echo.Context) error {
	return c.JSON(http.StatusOK, response{cvData.References})
}

func handleReference(c echo.Context) error {
	i, err := strconv.Atoi(c.Param("rid"))
	if err != nil {
		return err
	}
	if i < 0 || i > len(cvData.References)-1 {
		return c.JSON(http.StatusNotFound, errRefNotFound)
	}
	return c.JSON(http.StatusOK, response{cvData.References[i]})
}
