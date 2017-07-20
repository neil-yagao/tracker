package controllers

import (
	"models"
	"services/nutrition"
)

type NutritionController struct {
	GeneralController
}

// @router /physique/:user [put]
func (this *NutritionController) SaveUserPhysiqueInfo() {
	defer this.RecoverFromError()
	var info = new(models.PhysiqueInfo)
	userId := this.GetIntParam(":user")
	this.ParseRequestBody(info)
	nutrition.SaveUserPhysique(userId, info)
	this.ServeJson()
}
