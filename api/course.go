package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/labstack/echo"
	"github.com/spech66/easyskilltracker/helper"
)

// CourseMeta contains csv/json data
type CourseMeta struct {
	File        string
	Name        string `json:"name"`
	Description string `json:"description"`
	Author      string `json:"author"`
	URL         string `json:"url"`
	Version     string `json:"version"`
	Icon        string `json:"icon"`
	//Groups      []CourseGroup `json:"groups"` // Ignore on this level
}

// GetCourse returns all courses
func GetCourse() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		path := c.Get("data").(string)

		files, err := ioutil.ReadDir(path)
		if err != nil {
			panic(err)
		}

		var courses []CourseMeta
		for _, f := range files {
			if !f.IsDir() && strings.HasSuffix(f.Name(), ".json") {
				filePath := filepath.Join(path, f.Name())
				fmt.Println("Course data from", filePath)

				jsonFile, err := os.Open(filePath)
				if err != nil {
					panic(err)
				}
				defer jsonFile.Close()
				jsonData, _ := ioutil.ReadAll(jsonFile)

				var course CourseMeta
				if err := json.Unmarshal(jsonData, &course); err != nil {
					panic(err)
				}
				fmt.Println(course)

				course.File = strings.Replace(f.Name(), ".json", "", -1)
				courses = append(courses, course)
			}
		}

		jsonData := helper.DataToJSON(&courses)
		return c.JSONBlob(http.StatusOK, jsonData)
	}
}

// PostCourse creates new course
func PostCourse() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		path := c.Get("data").(string)
		courseName := c.FormValue("name")

		// Create clean string - https://golangcode.com/how-to-remove-all-non-alphanumerical-characters-from-a-string/
		reg, err := regexp.Compile("[^a-zA-Z0-9]+")
		courseFile := reg.ReplaceAllString(courseName, "_")

		fmt.Println("Clean name ", courseFile)
		if courseFile == "" || !helper.CheckFilename(courseFile) {
			return c.JSONBlob(http.StatusBadRequest, []byte(`[]`))
		}

		filePath := filepath.Join(path, courseFile+".json")
		fmt.Println("Course data at", filePath)

		var course CourseMeta
		course.Name = courseName
		fmt.Println(course)

		newCourseData, _ := json.MarshalIndent(course, "", "    ")
		err = ioutil.WriteFile(filePath, newCourseData, 0644)
		if err != nil {
			return c.JSONBlob(http.StatusBadRequest, []byte(``))
		}
		return c.JSONBlob(http.StatusOK, newCourseData)
	}
}
