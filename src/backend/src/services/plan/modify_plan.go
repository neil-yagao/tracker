package plan

import (
	"models"
)

func AssignSessionToPlan(pId int64, s *models.Session) {
	var p = new(models.Plan)
	p.Id = pId
	o.Read(p)
	s.Plan = p
	o.Insert(s)
	p.Sessions = append(p.Sessions, s)
}

func RemoveSessionFromPlan(pId int64, sId int64) {
	//Delete all the Movement from session
	var p = new(models.Plan)
	p.Id = pId
	o.Read(p)
	o.LoadRelated(p, "Sessions")
	var s = new(models.Session)
	s.Id = sId
	o.Read(s)
	o.LoadRelated(s, "Workouts")
	for _, sm := range s.Workouts {
		o.Delete(sm)
	}
	o.Delete(s)
	for i, session := range p.Sessions {
		if s.Id == session.Id {
			p.Sessions = append(p.Sessions[:i], p.Sessions[i+1:]...)
			return
		}
	}
}

func AddMovementToSession(sId int64, sm *models.SessionMovement) {
	var s = new(models.Session)
	s.Id = sId
	o.Read(s)
	sm.Session = s
	if sm.Reps == 0 {
		sm.Reps = 12
	}
	if sm.Sets == 0 {
		sm.Sets = 4
	}
	o.Insert(sm)
	s.Workouts = append(s.Workouts, sm)
}

func DeleteMovementFromSession(s *models.Session, sm *models.SessionMovement) {
	o.Delete(sm)
	for i, session := range s.Workouts {
		if s.Id == session.Id {
			s.Workouts = append(s.Workouts[:i], s.Workouts[i+1:]...)
			return
		}
	}
}

func FindPlan(id int64) *models.Plan {
	var plan = new(models.Plan)
	o.QueryTable("plan").Filter("id", id).RelatedSel().One(plan)
	o.LoadRelated(plan, "Sessions")
	for _, session := range plan.Sessions {
		o.LoadRelated(session, "Workouts")
		for _, w := range session.Workouts {
			o.LoadRelated(w, "Movement")
		}
	}
	return plan
}
