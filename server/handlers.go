package server

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
	"net/http"
)

func handlePersonalStatement(c echo.Context) error {
	return c.JSON(http.StatusOK, cvData.Statement)
}

func handleKeySkills(c echo.Context) error {
	return c.JSON(http.StatusOK, cvData.Skills)
}

func handleKeySkill(c echo.Context) error {
	i := cast.ToInt(c.Param("id"))
	if i < 0 || i > len(cvData.Skills)-1 {
		return c.JSON(http.StatusNotFound, errSkillNotFound)
	}
	return c.JSON(http.StatusOK, cvData.Skills[i])
}

func handleAchievements(c echo.Context) error {
	return c.JSON(http.StatusOK, cvData.Achievements)
}
func handleAchievement(c echo.Context) error {
	i := cast.ToInt(c.Param("id"))
	if i < 0 || i > len(cvData.Achievements)-1 {
		return c.JSON(http.StatusNotFound, errAchievementNotFound)
	}
	return c.JSON(http.StatusOK, cvData.Achievements[i])
}

func handleWorkHistory(c echo.Context) error {
	return c.JSON(http.StatusOK, cvData.WorkHistory)
}

func handleWorkPlace(c echo.Context) error {
	i := cast.ToInt(c.Param("id"))
	if i < 0 || i > len(cvData.WorkHistory)-1 {
		return c.JSON(http.StatusNotFound, errJobNotFound)
	}
	return c.JSON(http.StatusOK, cvData.WorkHistory[i])
}

func handleEducation(c echo.Context) error {
	return c.JSON(http.StatusOK, cvData.Education)
}

func handleEducationPlace(c echo.Context) error {
	i := cast.ToInt(c.Param("id"))
	if i < 0 || i > len(cvData.Education)-1 {
		return c.JSON(http.StatusNotFound, errStudiesNotFound)
	}
	return c.JSON(http.StatusOK, cvData.Education[i])
}

func handleHobbies(c echo.Context) error {
	return c.JSON(http.StatusOK, cvData.Hobbies)
}

func handleReferences(c echo.Context) error {
	return c.JSON(http.StatusOK, cvData.References)
}

func handleReference(c echo.Context) error {
	i := cast.ToInt(c.Param("id"))
	if i < 0 || i > len(cvData.References)-1 {
		return c.JSON(http.StatusNotFound, errRefNotFound)
	}
	return c.JSON(http.StatusOK, cvData.References[i])
}

func handleDownloadCV(c echo.Context) error {
	return c.Blob(http.StatusOK, "application/pdf", cvData.CVinPDF)
}
