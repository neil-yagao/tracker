package controllers

import (
	"encoding/json"
	"models"

	"github.com/astaxie/beego"
)

type WorkoutController struct {
	beego.Controller
}

type workout struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Target      string `json:"target"`
	PerformDate string `json:"perform_date"`
	Description string `json:"description"`
}

type workoutTemplate struct {
	Name      string
	Movements []MovementTemplate
	Weekly    string
	Addition  string
}

// @router /workouts [get]
func (this *WorkoutController) getWorkouts() {
	rows := models.Querier.Query("select id, name, target, perform_date, description from workout")
	workouts := make([]*workout, 0)
	for rows.Next() {
		one := new(workout)
		rows.Scan(&one.Id, &one.Name, &one.Target, &one.PerformDate, &one.Description)
		workouts = append(workouts, one)
	}
	this.Data["json"] = map[string]interface{}{"data": workouts}
	this.ServeJSON()
}

// @router /workouts [put]
func (this *WorkoutController) insertWorkout() {
	var template workoutTemplate
	// response.Count = 123
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &template)
}
