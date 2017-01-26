package controllers

import "models"

type TemplateController struct {
	GeneralController
}

//@router /templates [get]
func (this *TemplateController) GetTemplates() {
	templates := make([]models.Template, 0)
	rows := models.BasicCRUD.Query("select id, name from template")
	for rows.Next() {
		template := new(models.Template)
		rows.Scan(&template.Id, &template.Name)
		templates = append(templates, *template)
	}
	this.ServeJson(templates)
}
