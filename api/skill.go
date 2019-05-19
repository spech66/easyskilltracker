package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/labstack/echo"
	"github.com/spech66/easyskilltracker/helper"
)

// Course contains csv/json data
type Course struct {
	File        string
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Author      string        `json:"author"`
	URL         string        `json:"url"`
	Version     string        `json:"version"`
	Icon        string        `json:"icon"`
	Groups      []CourseGroup `json:"groups"`
}

// CourseGroup contains csv/json data
type CourseGroup struct {
	Name   string  `json:"name"`
	Order  int     `json:"order"`
	Skills []Skill `json:"skills"`
}

// Skill contains csv/json data
type Skill struct {
	Name        string   `json:"name"`
	Order       int      `json:"order"`
	Description string   `json:"description"`
	Resources   []string `json:"resources"`
	Progress    int      `json:"progress"`
}

// GetSkill returns all skill data
func GetSkill() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		path := c.Get("data").(string)
		courseFile := c.Param("course")
		fmt.Println("Get for course", courseFile)
		if courseFile == "" || !helper.CheckFilename(courseFile) {
			return c.JSONBlob(http.StatusBadRequest, []byte(`[]`))
		}

		filePath := filepath.Join(path, courseFile+".json")
		fmt.Println("Course data from", filePath)

		jsonFile, err := os.Open(filePath)
		if err != nil {
			panic(err)
		}
		defer jsonFile.Close()
		jsonData, _ := ioutil.ReadAll(jsonFile)

		/*var course Course
		if err := json.Unmarshal(jsonData, &course); err != nil {
			panic(err)
		}
		fmt.Println(course)

		jsonDataExport := helper.DataToJSON(&course)
		return c.JSONBlob(http.StatusOK, jsonDataExport)*/

		return c.JSONBlob(http.StatusOK, jsonData)
	}
}
