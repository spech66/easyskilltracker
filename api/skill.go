package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/labstack/echo"
	"github.com/spech66/easyskilltracker/helper"
)

// Skill contains csv/json data
type Skill struct {
	Group       string
	Name        string `json:"name"`
	Description string `json:"description"`
	Resources   string `json:"resources"`
	Progress    string `json:"progress"`
}

// GetSkill returns all skill data
func GetSkill() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		path := c.Get("data").(string)

		files, err := ioutil.ReadDir(path)
		if err != nil {
			panic(err)
		}

		var skills []Skill
		for _, f := range files {
			if !f.IsDir() {
				filePath := filepath.Join(path, f.Name())
				fmt.Println("Skill data from", filePath)
				lines := helper.ReadAllDataFromCSV(filePath)

				firstLine := true
				for _, line := range lines {
					data := Skill{
						Group:       strings.Replace(f.Name(), ".csv", "", -1),
						Name:        line[0],
						Description: line[1],
						Resources:   line[2],
						Progress:    line[3],
					}
					if !firstLine {
						skills = append(skills, data)
					}
					firstLine = false
				}
			}
		}

		jsonData := helper.DataToJSON(&skills)
		return c.JSONBlob(http.StatusOK, jsonData)
	}
}
