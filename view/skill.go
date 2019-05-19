package view

import (
	"net/http"

	"github.com/labstack/echo"
)

// SkillHandler renders the skill view
func SkillHandler(c echo.Context) error {
	course := c.Param("course")
	return c.Render(http.StatusOK, "skill.htm", map[string]interface{}{
		"course": course,
	})
}
