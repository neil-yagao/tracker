package controllers

import (
	"models"
	"services/session"
)

type WorkingSessionController struct {
	GeneralController
}

// @router /session/?:username [get]
func (this *WorkingSessionController) FindUserSession() {
	defer this.RecoverFromError()
	username := this.Ctx.Input.Param(":username")
	user := new(models.UserInfo)
	user.Username = username
	o.Read(user, "Username")
	this.ServeJson(session.FindUserSessions(user))
}

// @router /session-detail/?:sessionId [get]
func (this *WorkingSessionController) GetSessionDetail() {
	defer this.RecoverFromError()
	id := this.GetIntParam(":sessionId")
	this.ServeJson(session.FindSessionDetail(id))

}

// @router /session/?:id [post]
func (this *WorkingSessionController) SettleSession() {
	defer this.RecoverFromError()
	id := this.GetIntParam(":id")
	session.AchievedSession(id)
	this.ServeJson()
}

// @router /session-movement/?:movementId [post]
func (this *WorkingSessionController) DoneOneMovement() {
	defer this.RecoverFromError()
	exercises := make([]*models.Exercise, 0)
	this.ParseRequestBody(&exercises)
	id := this.GetIntParam(":movementId")
	session.DoneOneMovement(id, exercises)
	this.ServeJson()
}
