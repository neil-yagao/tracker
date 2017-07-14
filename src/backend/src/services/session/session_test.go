package session

import (
	"services/plan"
	"testing"
)

func TestAssignPlan(t *testing.T) {
	p := plan.FindPlan("测试训练计划")
	t.Log(p)
	ApplyPlanToUser(p, p.CreateBy, "2017-07-13")
	t.Log(FindUserSessions(p.CreateBy))
}
