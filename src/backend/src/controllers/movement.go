package controllers

import (
	"github.com/astaxie/beego/logs"
	"models"
	"services/movement"
)

type MovementController struct {
	GeneralController
}

// @router /movements [get]
func (this *MovementController) GetMovements() {
	this.ServeJson(movement.GetAllMovement())
}

// @router /movement [put]
func (this *MovementController) AddNewMovement() {
	m := new(models.Movement)
	this.ParseRequestBody(m)
	logs.Debug(m)
	movement.InsertMovement(m)
	this.ServeJson()
}

// @router /movement/upload/?:user/?:movement
func (this *MovementController) AttachMovementGif() {
	movement := this.GetString(":movement")
	logs.Debug("movement:" + movement)
	this.SaveToFile("gif", "./"+movement+".gif")
	this.ServeJson()
}
