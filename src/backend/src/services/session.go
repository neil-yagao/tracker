package services

import (
	"models"

	"github.com/astaxie/beego"
)

type sessionService struct {
}

var SessionService sessionService

const QUERY_WORKOUT_SESSION_SQL = " select ws.id, ws.sequence, ws.target_number, " +
	"ws.target_weight, ws.acheive_weight, ws.acheive_number, m.name, m.dividable from movement m, working_set ws, " +
	"workout w where w.id = :workoutId and ws.workout = w.id and m.id = ws.movement "

func (this *sessionService) GetWorkoutSession(workoutId int) []models.DividableWorkingSet {
	condition := make(map[string]interface{})
	condition["workoutId"] = workoutId
	rows := models.BasicCRUD.BuildAndQuery(QUERY_WORKOUT_SESSION_SQL, condition)
	defer rows.Close()
	sessions := make([]models.DividableWorkingSet, 0)
	for rows.Next() {
		instance := models.DividableWorkingSet{}
		rows.Scan(&instance.Id, &instance.Sequence, &instance.TargetNumber,
			&instance.TargetWeight, &instance.AcheiveWeight, &instance.AcheiveNumber, &instance.MovementName, &instance.Dividable)
		sessions = append(sessions, instance)
	}
	return sessions
}

func (this *sessionService) FinalizeSession(workingsets []models.WorkingSet, workoutId string) {
	completedSets := make([]map[string]interface{}, 0)
	for _, set := range workingsets {
		beego.Debug("debug:", set.AcheiveNumber, set.Id)
		var completeSet = map[string]interface{}{"acheive_number": set.AcheiveNumber, "id": set.Id}
		completeSet["acheive_weight"] = set.AcheiveWeight
		completedSets = append(completedSets, completeSet)
	}
	models.BasicCRUD.BuildAndUpdate("update working_set set acheive_number = :acheive_number , acheive_weight = :acheive_weight where id = :id", completedSets)
	models.BasicCRUD.Update("update workout set is_finalized = 1  where id = " + workoutId)
}

func (this *sessionService) UpdateSessionMovement(updateInfo *models.SessionUpdateInfo) {
	params := make([]map[string]interface{}, 0)
	for _, id := range updateInfo.WorkSet {
		params = append(params, map[string]interface{}{"movement": updateInfo.Movement, "id": id})
	}
	models.BasicCRUD.BuildAndUpdate("update working_set set movement = :movement where id = :id", params)
}
