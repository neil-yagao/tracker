package controllers

import (
	"models"
	"services/plan"
	"services/session"
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
	this.ServeJson(newPlan)
}

//@router /plans [get]
func (this *PlanController) GetAllPlans() {
	defer this.RecoverFromError()
	this.ServeJson(plan.GetAllPlan())
}

//@router /plan/?:plan [get]
func (this *PlanController) GetPlan() {
	defer this.RecoverFromError()
	planId := this.GetIntParam(":plan")
	this.ServeJson(plan.FindPlan(planId))
}

type AssignPlanCondition struct {
	User *models.UserInfo `json:"user"`
	Date string           `json:"date"`
}

// @router /plan/?:plan [post]
func (this *PlanController) AssignPlan() {
	defer this.RecoverFromError()
	var assignInfo = new(AssignPlanCondition)
	this.ParseRequestBody(assignInfo)
	id := this.GetIntParam(":plan")
	session.ApplyPlanToUser(id, assignInfo.User, assignInfo.Date)
	this.ServeJson()
}

// @router /plan/?:plan/sessions [post]
func (this *PlanController) AddSessionToPlan() {
	defer this.RecoverFromError()
	planId := this.GetIntParam(":plan")
	session := new(models.Session)
	this.ParseRequestBody(session)
	plan.AssignSessionToPlan(planId, session)
	this.ServeJson()
}

// @router /plan/?:plan/session/?:session [delete]
func (this *PlanController) RemoveSessionFromPlan() {
	defer this.RecoverFromError()
	planId := this.GetIntParam(":plan")
	sessionId := this.GetIntParam(":session")
	plan.RemoveSessionFromPlan(planId, sessionId)
	this.ServeJson()
}

// @router /session/?:session/movements [post]
func (this *PlanController) AddMovementToSession() {
	defer this.RecoverFromError()
	sessionId := this.GetIntParam(":session")
	var sessionMovement = new(models.SessionMovement)
	this.ParseRequestBody(sessionMovement)
	plan.AddMovementToSession(sessionId, sessionMovement)
	this.ServeJson()
}
