package session

import (
	"models"
	"time"
)

const IN_SESSION string = "executing"
const DONE_SESSION string = "achieved"

func DoneOneMovement(sw *models.SessionWorkout, sets []*models.Exercise) {
	for _, e := range sets {
		e.BelongWorkout = sw
		o.Insert(e)
		sw.Exercises = append(sw.Exercises, e)
	}
	sw.Status = "1"
	o.Update(sw)
}

func StartSession(us *models.UserSession) {
	us.ExecutionDate = time.Now()
	us.Status = IN_SESSION
	o.Update(us)
}

func AchievedSession(us *models.UserSession) {
	us.Status = DONE_SESSION
	o.Update(us)
}
