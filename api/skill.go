package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

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
	Order  int64   `json:"order"`
	Skills []Skill `json:"skills"`
}

// Skill contains csv/json data
type Skill struct {
	Name        string   `json:"name"`
	Order       int64    `json:"order"`
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

// UpgradeProgress sets the progress of a skill
func UpgradeProgress() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		path := c.Get("data").(string)
		courseFile := c.Param("course")
		// beware that we drop the err here as this is only a internal server solution!
		group, _ := strconv.ParseInt(c.FormValue("group"), 10, 64)
		skill, _ := strconv.ParseInt(c.FormValue("skill"), 10, 64)

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
		jsonData, _ := ioutil.ReadAll(jsonFile)
		jsonFile.Close()

		var course Course
		if err := json.Unmarshal(jsonData, &course); err != nil {
			panic(err)
		}
		fmt.Println(course)

		for idxG, g := range course.Groups {
			if g.Order == group {
				fmt.Println("Found group", g.Name)
				for idxS, s := range g.Skills {
					if s.Order == skill {
						fmt.Println("Skill group", s.Name)
						fmt.Println(course.Groups[idxG].Skills[idxS].Progress)
						course.Groups[idxG].Skills[idxS].Progress++
						if course.Groups[idxG].Skills[idxS].Progress >= 4 { // 0-3
							course.Groups[idxG].Skills[idxS].Progress = 0
						}
						fmt.Println(course.Groups[idxG].Skills[idxS].Progress)
					}
				}
			}
		}
		fmt.Println(course)

		newCourseData, _ := json.Marshal(course)
		err = ioutil.WriteFile(filePath, newCourseData, 0644)
		if err != nil {
			return c.JSONBlob(http.StatusBadRequest, []byte(``))
		}

		return c.JSON(http.StatusOK, "")
	}
}
