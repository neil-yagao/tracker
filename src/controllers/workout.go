package controllers

import (
	"encoding/json"
	"log"
	"models"

	"github.com/astaxie/beego"
)

type WorkoutController struct {
	beego.Controller
}

type Workout struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Target      string `json:"target"`
	PerformDate string `json:"perform_date"`
	Description string `json:"description"`
}

type WorkoutTemplate struct {
	Name         string             `json:"name"`
	Movements    []MovementTemplate `json:"movements"`
	StartAt      string             `json:"startAt"`
	Weekly       string             `json:"weekly"`
	Addition     string             `json:"addition"`
	TargetMuscle string             `json:"targetMuscle"`
	Description  string             `json:"description"`
}

// @router /workouts [get]
func (this *WorkoutController) getWorkouts() {
	rows := models.Querier.Query("select id, name, target, perform_date, description from workout")
	workouts := make([]*Workout, 0)
	for rows.Next() {
		one := new(Workout)
		rows.Scan(&one.Id, &one.Name, &one.Target, &one.PerformDate, &one.Description)
		workouts = append(workouts, one)
	}
	this.Data["json"] = map[string]interface{}{"data": workouts}
	this.ServeJSON()
}

// @router /workouts [put]
func (this *WorkoutController) insertWorkout() {
	var template WorkoutTemplate
	// response.Count = 123
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &template)
	if err != nil {
		log.Fatalln(err)
		panic(err)
	}

}
