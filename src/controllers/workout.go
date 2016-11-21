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
	rows := models.BasicCRUD.Query("select id, name, target, perform_date, description from workout")
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
		log.Fatalln(err)
		panic(err)
	}
	//experiment using goroutine
	log.Print("received template:")
	log.Println(template)
	go services.WorkoutCreator.CreateWorkoutsFromeTemplate(template)
	this.Data["json"] = map[string]interface{}{"data": "", "success": true}
	this.ServeJSON()

}
