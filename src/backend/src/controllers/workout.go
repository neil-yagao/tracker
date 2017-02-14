package controllers

import (
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
	workoutService := new(services.Workout)
	this.ServeJson(workoutService.FindUserWorkouts(user))
}

// @router /workouts [put]
func (this *WorkoutController) InsertWorkout() {
	defer this.RecoverFromError()
	template := new(services.WorkoutDefinition)
	this.ParseRequestBody(template)
	template.AssignWorkoutsToTemplate()
	this.ServeJson()
}
