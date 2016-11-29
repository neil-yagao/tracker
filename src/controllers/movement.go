package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"log"
	"models"
)

type MovementController struct {
	beego.Controller
}

/*const MOVEMENT_QUERY string = "SELECT 'm.name' as movement, 'w.name' as workout, ws.target_weight as weight, " +
"ws.target_number as default_number " + "from powerlift.movement m, powerlift.working_set ws, " +
"powerlift.workout w where w.id = ws.workout and ws.movement = m.id;"*/

// @router /movements [get]
func (this *MovementController) GetMovements() {
	rows := models.BasicCRUD.Query("select id, target_muscle, secondary_muscle, name, description from movement where ")
	movements := make([]*models.Movement, 0)
	for rows.Next() {
		one := new(models.Movement)
		rows.Scan(&one.Id, &one.TargetMuscle, &one.SecondaryMuscle, &one.Name, &one.Description)
		movements = append(movements, one)
	}
	this.Data["json"] = map[string]interface{}{"data": movements}
	this.ServeJSON()
}

// @router /movement [put]
func (this *MovementController) InsertMovement() {
	newMovement := make([]models.Movement, 0)
	log.Println("request body:", this.Ctx.Input.RequestBody)
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &newMovement)
	if err != nil {
		beego.Error(err)
	}
	models.BasicCRUD.Save("movement", newMovement)
	this.Data["json"] = map[string]interface{}{"success": true}
	this.ServeJSON()
}
