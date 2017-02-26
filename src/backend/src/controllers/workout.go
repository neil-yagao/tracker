package controllers

import (
	"services"
	"strconv"

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

// @router /workout/:id/movements [get]
func (this *WorkoutController) GetWorkoutMovements() {
	defer this.RecoverFromError()
	workoutId := this.Ctx.Input.Param(":id")

	workout := new(services.WorkoutMovement)
	id, _ := strconv.Atoi(workoutId)
	workout.Workout = int64(id)
	this.ServeJson(workout.GetWorkoutMovements())
}
