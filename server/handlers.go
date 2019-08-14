package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
)

func handleRootToSwag(c echo.Context) error {
	return c.Redirect(302, "/swagger/index.html")
}

// @Summary Personal Statement
// @Description A simple statement depicting my experience, and stuff
// @Accept json
// @Produce plain
// @Success 200 {string} string "A simple summary"
// @Router /about [GET]
func handlePersonalStatement(c echo.Context) error {
	return c.JSON(http.StatusOK, cvData.Statement)
}

// @Summary Key skills
// @Description A list of some key skills I consider noteworthy
// @Accept json
// @Produce json
// @Success 200 {array} string
// @Router /skills [GET]
func handleKeySkills(c echo.Context) error {
	return c.JSON(http.StatusOK, cvData.Skills)
}

// @Summary A single skill
// @Description A single, cool skill
// @Accept json
// @Produce plain
// @Param id path int64 true "Skill ID"
// @Success 200 {string} string
// @Failure 404 {object} server.responseError
// @Router /skills/{id} [GET]
func handleKeySkill(c echo.Context) error {
	i := cast.ToInt(c.Param("id"))
	if i < 0 || i > len(cvData.Skills)-1 {
		return c.JSON(http.StatusNotFound, errSkillNotFound)
	}
	return c.JSON(http.StatusOK, cvData.Skills[i])
}

// @Summary All achievements and noteworthy facts
// @Description A listing of some achievements I've achieved
// @Accept json
// @Produce json
// @Success 200 {array} string "List of achievements"
// @Router /achievements [GET]
func handleAchievements(c echo.Context) error {
	return c.JSON(http.StatusOK, cvData.Achievements)
}

// @Summary A single achievement
// @Description A single, noteworthy achievement
// @Accept json
// @Produce plain
// @Param id path int64 true "Achievement ID"
// @Success 200 {string} string "A single thing"
// @Failure 404 {object} server.responseError
// @Router /achievements/{id} [GET]
func handleAchievement(c echo.Context) error {
	i := cast.ToInt(c.Param("id"))
	if i < 0 || i > len(cvData.Achievements)-1 {
		return c.JSON(http.StatusNotFound, errAchievementNotFound)
	}
	return c.JSON(http.StatusOK, cvData.Achievements[i])
}

// @Summary Most jobs I've had so far
// @Description A listing of the most relevant job positions I've held.
// @Description Early jobs, such as waiter at a small-town café, are omitted
// @Accept json
// @Produce json
// @Success 200 {array} types.WorkSpan
// @Router /work [GET]
func handleWorkHistory(c echo.Context) error {
	return c.JSON(http.StatusOK, cvData.WorkHistory)
}

// @Summary A single workplace
// @Description A single place I've worked in
// @Accept json
// @Produce plain
// @Param id path int64 true "Job ID"
// @Success 200 {object} types.WorkSpan
// @Failure 404 {object} server.responseError
// @Router /work/{id} [GET]
func handleWorkPlace(c echo.Context) error {
	i := cast.ToInt(c.Param("id"))
	if i < 0 || i > len(cvData.WorkHistory)-1 {
		return c.JSON(http.StatusNotFound, errJobNotFound)
	}
	return c.JSON(http.StatusOK, cvData.WorkHistory[i])
}

// @Summary Courses I've attended
// @Description A listing of most courses I've attended. However, given I'm mostly self-taught, it's not really representative of my expertise
// @Accept json
// @Produce json
// @Success 200 {array} types.EducationSpan
// @Router /education [GET]
func handleEducation(c echo.Context) error {
	return c.JSON(http.StatusOK, cvData.Education)
}

// @Summary A single course
// @Description A single course I've attended
// @Accept json
// @Produce plain
// @Param id path int64 true "Course ID"
// @Success 200 {object} types.EducationSpan
// @Failure 404 {object} server.responseError
// @Router /education/{id} [GET]
func handleEducationPlace(c echo.Context) error {
	i := cast.ToInt(c.Param("id"))
	if i < 0 || i > len(cvData.Education)-1 {
		return c.JSON(http.StatusNotFound, errStudiesNotFound)
	}
	return c.JSON(http.StatusOK, cvData.Education[i])
}

// @Summary Hobbies
// @Description Some relevant hobbies I dabble with
// @Accept json
// @Produce json
// @Success 200 {string} string
// @Router /hobbies [GET]
func handleHobbies(c echo.Context) error {
	return c.JSON(http.StatusOK, cvData.Hobbies)
}

// @Summary References
// @Description Some people which whom I've worked with, and can vouch for me
// @Accept json
// @Produce json
// @Success 200 {array} types.Reference
// @Router /references [GET]
func handleReferences(c echo.Context) error {
	return c.JSON(http.StatusOK, cvData.References)
}

// @Summary A single reference
// @Description Someone who can vouch for me
// @Accept json
// @Produce json
// @Param id path int64 true "Reference ID"
// @Success 200 {object} types.Reference
// @Failure 404 {object} server.responseError
// @Router /references/{id} [GET]
func handleReference(c echo.Context) error {
	i := cast.ToInt(c.Param("id"))
	if i < 0 || i > len(cvData.References)-1 {
		return c.JSON(http.StatusNotFound, errRefNotFound)
	}
	return c.JSON(http.StatusOK, cvData.References[i])
}

// @Summary CV in PDF form
// @Description Download or view this CV in PDF form to share with your HR friends
// @Accept json
// @Produce application/pdf
// @Success 200 {file} string
// @Router /pdf [GET]
func handleDownloadCV(c echo.Context) error {
	return c.Blob(http.StatusOK, "application/pdf", cvData.PDFpayload)
}
