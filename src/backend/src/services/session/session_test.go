package session

import (
	"services/plan"
	"testing"
)

func TestAssignPlan(t *testing.T) {
	p := plan.FindPlan(1)
	t.Log(p)
	ApplyPlanToUser(1, p.CreateBy, "2017-07-19")
	t.Log(FindUserSessions(p.CreateBy))
}
