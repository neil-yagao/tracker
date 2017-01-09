package controllers

import (
	"models"
	"services"

	"github.com/astaxie/beego"
)

type WorkoutController struct {
	GeneralController
}

// @router /workouts [get]
func (this *WorkoutController) GetWorkouts() {
	defer this.RecoverFromError()
	beego.Debug("query for workouts")
	user := this.GetUserIdentity()
	workouts := services.WorkoutCreator.FindUserWorkouts(user)
	this.ServeJson(workouts)
}

// @router /workouts [put]
func (this *WorkoutController) InsertWorkout() {
	defer this.RecoverFromError()
	template := new(models.WorkoutTemplate)
	this.ParseRequestBody(template)
	user := this.GetUserIdentity()
	services.WorkoutCreator.CreateWorkoutsFromeTemplate(*template, user)
	this.ServeJson()

}
