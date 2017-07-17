package session

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"models"
	"time"
)

const NEW_SESSION string = "assigned"

var o = orm.NewOrm()
var logger = logs.GetLogger("session")

const WEEK_HOUR = 168

var day, _ = time.ParseDuration("24h")

func ApplyPlanToUser(p *models.Plan, user *models.UserInfo, startDate string) {
	var assignedPlan = new(models.AssignedPlan)
	assignedPlan.AssignTime = time.Now()
	assignedPlan.AssignTo = user
	assignedPlan.AssignedPlan = p
	assignedPlan.ExecutingStatus = NEW_SESSION
	logs.Debug("assign plan %v to user %v", p, user)
	o.Insert(assignedPlan)
	createUserSessonBasedOnPlan(p, user, startDate)
}

func createUserSessonBasedOnPlan(p *models.Plan, user *models.UserInfo, startDate string) []*models.UserSession {
	startTimePoint, _ := time.Parse("2006-01-02", startDate)
	//find earlist weekly repeat
	var userSessions []*models.UserSession = make([]*models.UserSession, 0)
	for _, session := range p.Sessions {
		var weekly = session.Weekly
		var nextWorkoutDate = findNextWeekday(weekly, startTimePoint)
		var timeSpan, _ = time.ParseDuration(fmt.Sprint(WEEK_HOUR) + "h")
		var i int64 = 0
		for ; i < session.Repeat; i++ {
			var userSession = new(models.UserSession)
			userSession.AssignTo = user
			userSession.ExpectingDate = nextWorkoutDate
			userSession.Status = NEW_SESSION
			userSession.OriginSession = session
			o.Insert(userSession)
			userSession.Workouts = createSessionWorkout(session, userSession)
			nextWorkoutDate = nextWorkoutDate.Add(timeSpan)
			userSessions = append(userSessions, userSession)
		}
	}
	return userSessions
}

func createSessionWorkout(s *models.Session, us *models.UserSession) []*models.SessionWorkout {
	var workouts []*models.SessionWorkout = make([]*models.SessionWorkout, 0)
	logger.Println(s.Workouts)
	for _, sm := range s.Workouts {
		workout := new(models.SessionWorkout)
		workout.MappedMovement = sm
		workout.BelongSession = us
		workout.Status = "0"
		o.Insert(workout)
		workouts = append(workouts, workout)
	}
	return workouts
}

func findNextWeekday(weekday string, startPoint time.Time) time.Time {
	//using loop to escape unmatching infinite loop
	var timePoint time.Time
	var loop int = 0
	for timePoint = startPoint; loop < 8; timePoint = timePoint.Add(day) {
		if timePoint.Weekday().String() == weekday {
			return timePoint
		}
		loop++
	}
	panic("unable to indentify weekday:" + weekday)

}

func FindUserSessions(user *models.UserInfo) []*models.UserSession {
	var sessions []*models.UserSession
	qs := o.QueryTable("user_session").Filter("assign_to_id", user.Id)
	_, err := qs.Filter("status", NEW_SESSION).RelatedSel().All(&sessions, "OriginSession", "ExpectingDate")
	/*for _, session := range sessions {
		o.LoadRelated(session, "Workouts")
		for _, workout := range session.Workouts {
			o.LoadRelated(workout, "MappedMovement")
			o.LoadRelated(workout.MappedMovement, "Movement")
		}
	}*/
	if err != nil {
		panic(err)
	}
	return sessions
}

func FindSessionDetail(id int64) *models.UserSession {
	var session *models.UserSession = new(models.UserSession)
	session.Id = id
	o.Read(session)
	loadSessionDetail(session)
	return session
}

func loadSessionDetail(session *models.UserSession) {
	o.LoadRelated(session, "Workouts")
	o.LoadRelated(session, "OriginSession")
	for _, workout := range session.Workouts {
		o.LoadRelated(workout, "MappedMovement")
		o.LoadRelated(workout.MappedMovement, "Movement")
	}
}
