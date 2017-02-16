package controllers

import (
	"models"
	"services"
	"strconv"
)

type WorkingSessionController struct {
	GeneralController
}

// @router /session/?:workout [get]
func (this *WorkingSessionController) GetWorkoutSessions() {
	defer this.RecoverFromError()
	workoutId, _ := strconv.Atoi((this.Ctx.Input.Param(":workout")))
	result := services.SessionService.GetWorkoutSession(workoutId)
	this.ServeJson(result)
}

// @router /session/?:workout [post]
func (this *WorkingSessionController) SaveSession() {
	defer this.RecoverFromError()
	achieved := make([]models.WorkingSet, 0)
	this.ParseRequestBody(&achieved)
	workoutId := this.Ctx.Input.Param(":workout")
	userIdentify := this.GetUserIdentity()
	services.SessionService.FinalizeSession(achieved, workoutId, userIdentify)
	this.ServeJson()
}

// @router /session/?:workout/movement [post]
func (this *WorkingSessionController) UpdateSessionMovement() {
	defer this.RecoverFromError()
	updateInfo := new(models.SessionUpdateInfo)
	this.ParseRequestBody(updateInfo)
	services.SessionService.UpdateSessionMovement(updateInfo)
	this.ServeJson()
}
