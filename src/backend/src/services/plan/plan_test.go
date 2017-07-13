package plan

import (
	"models"
	"testing"
)

var testingPlan *models.Plan
var testingSession *models.Session
var testingSessionMovement *models.SessionMovement
var testingMovement *models.Movement
var testingUser *models.UserInfo

func init() {
	testingUser = new(models.UserInfo)
	testingUser.Username = "tester"
	o.Insert(testingUser)
	testingMovement = new(models.Movement)
	testingMovement.TargetMuscle = "胸"
	testingMovement.SecondaryMuscle = "肱三头肌"
	testingMovement.Name = "平板杠铃卧推"
	testingMovement.Description = "some description"
	testingMovement.Dividable = 0
	testingMovement.AddBy = testingUser
	o.Insert(testingMovement)

	testingPlan = new(models.Plan)
	testingPlan.Name = "测试训练计划"
	testingPlan.CreateBy = testingUser

	testingSession = new(models.Session)
	testingSession.Name = "胸部训练"
	testingSession.TargetMuscle = "胸"
	testingSession.Repeat = 4
	testingSession.Weekly = "Monday"
	testingSession.Plan = testingPlan
	testingPlan.Sessions = append(testingPlan.Sessions, testingSession)

	testingSessionMovement = new(models.SessionMovement)
	testingSessionMovement.Movement = testingMovement
	testingSessionMovement.Sets = 4
	testingSessionMovement.Reps = 12
	testingSessionMovement.Sequence = 1
	testingSessionMovement.NeedWarmup = 0
	testingSessionMovement.Session = testingSession
	testingSession.Workouts = append(testingSession.Workouts, testingSessionMovement)
}

func TestSave(t *testing.T) {
	t.Log("testing:", testingPlan)
	CreateNewPlan(testingPlan)
}
