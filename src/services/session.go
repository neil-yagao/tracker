package services

import "models"

type sessionService struct {
}

var SessionService sessionService

const QUERY_WORKOUT_SESSION_SQL = " select ws.sequence, ws.target_number, " +
	"ws.target_weight, ws.acheive_number, m.name from movement m, working_set ws, " +
	"workout w where w.id = :workoutId and ws.workout = w.id and m.id = ws.movement "

type session struct {
	Sequence     string `json:"sequence"`
	TargetNumner string `json:"targetNumber"`
	TargetWeight string `json:"targetWeight"`
	AcheiveNumer string `json:"acheiveNumber"`
	MovementName string `json:"movementName"`
}

func (this sessionService) GetWorkoutSession(workoutId int) []session {
	condition := make(map[string]interface{})
	condition["workoutId"] = workoutId
	rows := models.BasicCRUD.BuildAndQuery(QUERY_WORKOUT_SESSION_SQL, condition)
	sessions := make([]session, 0)
	for rows.Next() {
		instance := session{}
		rows.Scan(&instance.Sequence, &instance.TargetNumner,
			&instance.TargetWeight, &instance.AcheiveNumer, &instance.MovementName)
		sessions = append(sessions, instance)
	}
	return sessions
}
