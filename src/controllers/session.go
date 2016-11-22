package controllers

import (
	"services"
	"strconv"

	"github.com/astaxie/beego"
)

type WorkingSessionController struct {
	beego.Controller
}

// @router /session/?:workout [get]
func (this *WorkingSessionController) GetWorkoutSessions() {
	workoutId, _ := strconv.Atoi((this.Ctx.Input.Param(":workout")))
	result := services.SessionService.GetWorkoutSession(workoutId)
	this.Data["json"] = map[string]interface{}{"data": result}
	this.ServeJSON()
}
