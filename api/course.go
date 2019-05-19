package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/labstack/echo"
	"github.com/spech66/easyskilltracker/helper"
)

// Course contains csv/json data
type Course struct {
	File        string
	Name        string `json:"name"`
	Description string `json:"description"`
	Author      string `json:"author"`
	URL         string `json:"url"`
	Version     string `json:"version"`
	Icon        string `json:"icon"`
}

// GetCourse returns all courses
func GetCourse() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		path := c.Get("data").(string)

		files, err := ioutil.ReadDir(path)
		if err != nil {
			panic(err)
		}

		var courses []Course
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

				var course Course
				if err := json.Unmarshal(jsonData, &course); err != nil {
					panic(err)
				}
				fmt.Println(course)

				course.File = f.Name()
				courses = append(courses, course)
			}
		}

		jsonData := helper.DataToJSON(&courses)
		return c.JSONBlob(http.StatusOK, jsonData)
	}
}
