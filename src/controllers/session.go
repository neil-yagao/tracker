package controllers

import (
	"encoding/json"
	"log"
	"models"
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

// @router /session/?:workout [post]
func (this *WorkingSessionController) SaveSession() {
	achieved := make([]models.WorkingSet, 0)
	log.Println("request body:", this.Ctx.Input.RequestBody)
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &achieved)
	if err != nil {
		beego.Error(err)
	}
	log.Println(achieved)
	workoutId := this.Ctx.Input.Param(":workout")
	services.SessionService.FinalizeSession(achieved, workoutId)
	this.Data["json"] = map[string]interface{}{"success": true}
	this.ServeJSON()
}
