package controllers

import (
	"models"
	"services/session"
)

type WorkingSessionController struct {
	GeneralController
}

// @router /session/?:username
func (this *WorkingSessionController) FindUserSession() {
	defer this.RecoverFromError()
	username := this.Ctx.Input.Param(":username")
	user := new(models.UserInfo)
	user.Username = username
	o.Read(user, "Username")
	this.ServeJson(session.FindUserSessions(user))
}
