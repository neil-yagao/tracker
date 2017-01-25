package services

import (
	"models"
	"strings"
)

type templateService struct{}

var TemplateService templateService

func findTemplateId(template string) int64 {
	if len(strings.TrimSpace(template)) == 0 {
		return -1
	} else {

		rows := models.BasicCRUD.BuildAndQuery("select id from template where name = :name",
			map[string]interface{}{"name": template})
		defer rows.Close()
		var templateId int64 = -1
		if rows.Next() {
			rows.Scan(&templateId)
		}
		return templateId
	}
}

func saveTemplate(template models.Template) int64 {
	return models.BasicCRUD.Save("template", template)
}
