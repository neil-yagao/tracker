package controllers

import (
	"models"
	"services"
	"strconv"

	"github.com/astaxie/beego"
)

/**
 * TemplateController
 * Controller for template related
 */
type TemplateController struct {
	GeneralController
}

//@router /templates [get]
func (this *TemplateController) GetTemplates() {
	defer this.RecoverFromError()
	templates := make([]services.Template, 0)
	rows := models.BasicCRUD.Query("select id, name from template")
	for rows.Next() {
		template := new(services.Template)
		rows.Scan(&template.Id, &template.Name)
		templates = append(templates, *template)
	}
	this.ServeJson(templates)
}

//@router /templates/user [post]
func (this *TemplateController) AssignTemplates() {
	defer this.RecoverFromError()
	template := new(services.UserTemplate)
	this.ParseRequestBody(template)
	template.AssignTemplate()
	this.ServeJson()
}

//@router /templates/user [delete]
func (this *TemplateController) UnasignTemplates() {
	defer this.RecoverFromError()
	useridentity := this.GetUserIdentity()
	templateId, err := this.GetInt("templatesid")
	if err != nil {
		beego.Error(err)
		panic(err)
	}
	template := new(services.UserTemplate)
	template.UserIdentity = useridentity
	template.TemplatesId = append(make([]int64, 1), int64(templateId))
	template.UnassignTemplate()
	this.ServeJson()
}

//@router /templates/:useridentity [get]
func (this *TemplateController) GetUserCurrentTemplates() {
	defer this.RecoverFromError()
	userIdentity := this.Ctx.Input.Param(":useridentity")
	template := new(services.UserTemplate)
	template.UserIdentity = userIdentity
	this.ServeJson(template.GetUserTemplate())
}

//@router /templates/：id/workouts [get]
func (this *TemplateController) GetTemplateWorkouts() {
	defer this.RecoverFromError()
	templateId := this.Ctx.Input.Param(":id")
	template := new(services.TemplateWorkout)
	id, _ := strconv.Atoi(templateId)
	template.Template = int64(id)
	this.ServeJson(template.FindWorkouts())
}
