package controllers

import (
	"models"
	"services/plan"
)

type PlanController struct {
	GeneralController
}

//@router /plan [post]
func (this *PlanController) CreateNewPlan() {
	defer this.RecoverFromError()
	var newPlan = new(models.Plan)
	this.ParseRequestBody(newPlan)
	plan.CreateNewPlan(newPlan)
	this.ServeJson()
}

//@router /plans [get]
func (this *PlanController) GetAllPlans() {
	defer this.RecoverFromError()
	this.ServeJson(plan.GetAllPlan())
}

//@router /plan/?:plan [get]
func (this *PlanController) GetPlan() {
	defer this.RecoverFromError()
	planName := this.Ctx.Input.Param(":plan")
	this.ServeJson(plan.FindPlan(planName))
}
