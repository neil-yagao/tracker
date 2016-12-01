package controllers

import (
	"log"
	"models"
	"services"
	"strconv"
)

type WorkingSessionController struct {
	GeneralController
}

// @router /session/?:workout [get]
func (this *WorkingSessionController) GetWorkoutSessions() {
	workoutId, _ := strconv.Atoi((this.Ctx.Input.Param(":workout")))
	result := services.SessionService.GetWorkoutSession(workoutId)
	this.ServeJson(result)
}

// @router /session/?:workout [post]
func (this *WorkingSessionController) SaveSession() {
	achieved := make([]models.WorkingSet, 0)
	this.ParseRequestBody(&achieved)
	log.Println(achieved)
	workoutId := this.Ctx.Input.Param(":workout")
	services.SessionService.FinalizeSession(achieved, workoutId)
	this.ServeJson()
}
