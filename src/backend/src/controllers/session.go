package controllers

import (
	"models"
	"services/session"
	"strconv"
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
	sessionId := this.Ctx.Input.Param(":sessionId")
	id, err := strconv.ParseInt(sessionId, 10, 64)
	if err != nil {
		panic(err)
	}
	this.ServeJson(session.FindSessionDetail(id))
}

// @router /session/?:id [post]
func (this *WorkingSessionController) SettleSession() {
	defer this.RecoverFromError()
	sessionId := this.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(sessionId, 10, 64)
	if err != nil {
		panic(err)
	}
	session.AchievedSession(id)
	this.ServeJson()
}

// @router /session-movement/?:movementId [post]
func (this *WorkingSessionController) DoneOneMovement() {
	defer this.RecoverFromError()
	exercises := make([]*models.Exercise, 0)
	this.ParseRequestBody(&exercises)
	paramId := this.Ctx.Input.Param(":movementId")
	movementId, err := strconv.ParseInt(paramId, 10, 64)
	if err != nil {
		panic(err)
	}
	session.DoneOneMovement(movementId, exercises)
	this.ServeJson()
}
