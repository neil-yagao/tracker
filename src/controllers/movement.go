package controllers

import (
	"models/connect"

	"github.com/astaxie/beego"
)

type MovementController struct {
	beego.Controller
}

type movement struct {
	Id              int
	TargetMuscle    string
	SecondaryMuscle string
	Name            string
	Description     string
}

/*const MOVEMENT_QUERY string = "SELECT 'm.name' as movement, 'w.name' as workout, ws.target_weight as weight, " +
"ws.target_number as default_number" + "from powerlift.movement m, powerlift.working_set ws, " +
"powerlift.workout w where w.id = ws.workout and ws.movement = m.id;"*/

// @router /movements [get]
func (this *MovementController) getMovements() {
	rows := connect.Querier.Query("select id, target_muscle, secondary_muscle, name, description from movement where ")
	movements := make([]*movement, 0)
	for rows.Next() {
		one := new(movement)
		rows.Scan(&one.Id, &one.TargetMuscle, &one.SecondaryMuscle, &one.Equipment, &one.Description)
		movements = append(movements, one)
	}
	this.Data["json"] = map[string]interface{}{"data": movements}
	this.ServeJSON()
}
