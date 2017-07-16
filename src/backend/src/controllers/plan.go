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