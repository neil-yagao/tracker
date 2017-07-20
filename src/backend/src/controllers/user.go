package controllers

import (
	"models"
)

type UserController struct {
	GeneralController
}

// @router /login [post]
func (this *UserController) Login() {
	user := new(models.UserInfo)
	this.ParseRequestBody(user)
	if user.Id == 0 {
		//not register
		o.Insert(user)
	}
	o.Read(user, "Username")
	this.ServeJson(user)
}
