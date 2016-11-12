package controllers

import (
	"encoding/json"
	"models/connect"

	"github.com/astaxie/beego"
)

type WorkoutController struct {
	beego.Controller
}

type workout struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Target      string `json:"target"`
	PerformDate string `json:"performDate"`
	Description string `json:"description"`
}

// @router /workouts [get]
func (this *WorkoutController) getWorkouts() {
	rows := connect.Querier.Query("select id, name, target, perform_date, description from workout")
	workouts := make([]*workout, 0)
	for rows.Next() {
		one := new(workout)
		rows.Scan(&one.Id, &one.Name, &one.Target, &one.PerformdDate, &one.Description)
		movements = append(workouts, one)
	}
	this.Data["json"] = map[string]interface{}{"data": workouts}
	this.ServeJSON()
}

// @router /workouts [put]
func (this *WorkoutController) insertWorkout() {
	ob := make([]workout, 0)
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
}
