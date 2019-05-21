package view

import (
	"net/http"

	"github.com/labstack/echo"
)

// CourseHandler renders the course view
func CourseHandler(c echo.Context) error {
	course := c.Param("course")
	return c.Render(http.StatusOK, "course.htm", map[string]interface{}{
		"course": course,
	})
}
