package controllers

import (
	"models"
	"services"
)

type UserControllers struct {
	GeneralController
}

// @router /login [post]
func (this *UserControllers) UserLogin() {
	userInfo := new(models.LoginInfo)
	this.ParseRequestBody(userInfo)
	services.UserService.HandleUserLogin(*userInfo)
	this.ServeJson()
}