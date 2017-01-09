package controllers

import "models"

type MovementController struct {
	GeneralController
}

/*const MOVEMENT_QUERY string = "SELECT 'm.name' as movement, 'w.name' as workout, ws.target_weight as weight, " +
"ws.target_number as default_number " + "from powerlift.movement m, powerlift.working_set ws, " +
"powerlift.workout w where w.id = ws.workout and ws.movement = m.id;"*/

// @router /movements [get]
func (this *MovementController) GetMovements() {
	defer this.RecoverFromError()
	rows := models.BasicCRUD.Query("select id, target_muscle, secondary_muscle, name, description, dividable from movement")
	defer rows.Close()
	movements := make([]*models.Movement, 0)
	for rows.Next() {
		one := new(models.Movement)
		rows.Scan(&one.Id, &one.TargetMuscle, &one.SecondaryMuscle, &one.Name, &one.Description, &one.Dividable)
		movements = append(movements, one)
	}

	this.ServeJson(movements)
}

// @router /movements [put]
func (this *MovementController) InsertMovement() {
	defer this.RecoverFromError()
	newMovement := new(models.Movement)
	this.ParseRequestBody(newMovement)
	models.BasicCRUD.Save("movement", *newMovement)
	this.ServeJson()
}

// @router /movement/:id [post]
func (this *MovementController) UpdateMovement() {
	defer this.RecoverFromError()
	modifyMovement := new(models.Movement)
	this.ParseRequestBody(modifyMovement)
	models.BasicCRUD.BuildAndUpdateOne("movement", *modifyMovement)
	this.ServeJson()
}
