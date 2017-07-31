package session

import (
	"models"
)

func GetUserSessionMovement(user int64, movement int64) []*models.SessionWorkout {
	passedMovement := make([]*models.SessionWorkout, 0)
	o.QueryTable("session_workout").Filter("BelongSession__AssignTo__Id", user).Filter("MappedMovement_Id", movement)
	for _, sm := range passedMovement {
		o.LoadRelated(sm, "Exercises")
	}
	return passedMovement
}
