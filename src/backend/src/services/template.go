package services

import (
	"models"
	"strings"
)

type Template struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

func findTemplate(template string) *Template {
	var result *Template
	if len(strings.TrimSpace(template)) == 0 {
		return nil
	} else {

		rows := models.BasicCRUD.BuildAndQuery("select id, name, content from template where name = :name",
			map[string]interface{}{"name": template})
		defer rows.Close()
		if rows.Next() {
			rows.Scan(&result.Id, &result.Name, &result.Content)
		}
		return result
	}
}

func saveTemplate(template Template) int64 {
	return models.BasicCRUD.Save("template", template)
}
