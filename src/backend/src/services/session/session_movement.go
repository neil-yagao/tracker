package session

import (
	"models"
)

func GetUserSessionMovement(user int64, movement int64) []*models.SessionWorkout {
	passedMovement := make([]*models.SessionWorkout, 0)
	qs := o.QueryTable("session_workout").Filter("BelongSession__AssignTo__Id", user).Filter("MappedMovement__Movement__Id", movement)
	qs.All(&passedMovement)
	for _, sm := range passedMovement {
		o.LoadRelated(sm, "Exercises")
		o.LoadRelated(sm, "BelongSession")
	}
	return passedMovement
}
