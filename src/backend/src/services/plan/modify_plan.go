package plan

import (
	"models"
)

func AssignSessionToPlan(p *models.Plan, s *models.Session) {
	s.Plan = p
	o.Insert(s)
	p.Sessions = append(p.Sessions, s)
}

func RemoveSessionFromPlan(p *models.Plan, s *models.Session) {
	//Delete all the Movement from session
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

func AddMovementToSession(s *models.Session, sm *models.SessionMovement) {
	sm.Session = s
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

func FindPlan(name string) *models.Plan {
	var plan = new(models.Plan)
	o.QueryTable("plan").Filter("name", name).RelatedSel().One(plan)
	o.LoadRelated(plan, "Sessions")
	for _, session := range plan.Sessions {
		o.LoadRelated(session, "Workouts")
		for _, w := range session.Workouts {
			o.LoadRelated(w, "Movement")
		}
	}
	return plan
}
