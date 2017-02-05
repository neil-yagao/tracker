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
	template := this.GetString("template")
	workoutService := new(services.Workout)
	workouts := workoutService.FindUserWorkouts(user)
	workouts = append(workouts, workoutService.FindTemplateWorkouts(template)...)
	this.ServeJson(workouts)
}

// @router /workouts [put]
func (this *WorkoutController) InsertWorkout() {
	defer this.RecoverFromError()
	template := new(models.WorkoutTemplate)
	this.ParseRequestBody(template)
	user := this.GetUserIdentity()
	services.WorkoutCreator.CreateWorkoutsFromTemplate(*template, user)
	this.ServeJson()

}
