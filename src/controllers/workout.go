package controllers

import (
	"encoding/json"
	"log"
	"models"
	"services"

	"github.com/astaxie/beego"
)

type WorkoutController struct {
	beego.Controller
}

// @router /workouts [get]
func (this *WorkoutController) GetWorkouts() {
	log.Println("query for workouts")
	rows := models.BasicCRUD.Query("select id, name, target, perform_date, description from workout where is_finalized = 0 order by perform_date ASC")
	workouts := make([]*models.Workout, 0)
	for rows.Next() {
		one := new(models.Workout)
		rows.Scan(&one.Id, &one.Name, &one.Target, &one.PerformDate, &one.Description)
		workouts = append(workouts, one)
	}
	this.Data["json"] = map[string]interface{}{"data": workouts}
	this.ServeJSON()
}

// @router /workouts [put]
func (this *WorkoutController) InsertWorkout() {
	var template models.WorkoutTemplate
	// response.Count = 123
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &template)
	if err != nil {
		beego.Error(err)
	}
	//experiment using goroutine
	beego.Debug("received template:", template)
	go services.WorkoutCreator.CreateWorkoutsFromeTemplate(template)
	this.Data["json"] = map[string]interface{}{"data": "", "success": true}
	this.ServeJSON()

}
