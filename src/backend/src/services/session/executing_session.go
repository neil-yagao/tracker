package session

import (
	"models"
	"time"
)

const IN_SESSION string = "executing"
const DONE_SESSION string = "achieved"

func DoneOneMovement(id int64, sets []*models.Exercise) {
	var sw = new(models.SessionWorkout)
	sw.Id = id
	o.Read(sw)
	for _, e := range sets {
		e.BelongWorkout = sw
		o.Insert(e)
		sw.Exercises = append(sw.Exercises, e)
	}
	sw.Status = "2"
	o.Update(sw)
}

func StartSession(us *models.UserSession) {
	us.ExecutionDate = time.Now()
	us.Status = IN_SESSION
	o.Update(us)
}

func AchievedSession(id int64) {
	us := new(models.UserSession)
	us.Id = id
	o.Read(us)
	us.Status = DONE_SESSION
	o.Update(us)
}
