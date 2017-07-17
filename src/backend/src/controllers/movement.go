package controllers

import (
	"services/movement"
)

type MovementController struct {
	GeneralController
}

// @router /movements [get]
func (this *MovementController) GetMovements() {
	this.ServeJson(movement.GetAllMovement())
}
