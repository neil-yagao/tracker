package services

import (
	"fmt"
	"models"
	"time"

	"github.com/astaxie/beego"
)

type Workout struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Target      string `json:"target"`
	Description string `json:"description"`
	Sequence    int64  `json:"sequence"`
	Repeat      string `json:"repeat"`
}

const WEEK_HOUR = 168

var day, _ = time.ParseDuration("24h")

const QUERY_UESER_WORKOUT_QUERY string = "select w.id, w.name, w.target, uw.perform_date, w.description from workout w, user_workout uw, user u " +
	"where uw.is_finalized = 0 and uw.user = u.id and uw.workout = w.id and u.useridentity = :useridentity order by perform_date ASC"

func (this *Workout) FindUserWorkouts(user string) []*Workout {
	return findUserWorkouts(user)

}

func findUserWorkouts(user string) []*Workout {
	rows := models.BasicCRUD.BuildAndQuery(QUERY_UESER_WORKOUT_QUERY, map[string]interface{}{"useridentity": user})
	defer rows.Close()
	return scanWorkoutsResult(rows)
}

func assignWorkoutsToUser(userIdentity string, workouts []*Workout, startDate string) {
	//templateName :=
	user := UserService.GetUserIdFromIdentity(userIdentity)
	startTimePoint, _ := time.Parse("2006-01-02", startDate)
	for _, workout := range workouts {
		var timeSpan, _ = time.ParseDuration(fmt.Sprint(workout.Sequence*WEEK_HOUR) + "h")
		timePoint := findNextWeekday(workout.Repeat, startTimePoint.Add(timeSpan))
		assignWorkoutToUser(user, workout.Id, timePoint.String()[0:10])
	}
}

type UserWorkout struct {
	User        int64
	Workout     int64
	PerformDate string
	IsFinalized int8
}

func assignWorkoutToUser(user int64, workout int64, performDate string) {
	models.BasicCRUD.Save("user_workout", UserWorkout{user, workout, performDate, 0})
}

func findNextWeekday(weekday string, startPoint time.Time) time.Time {
	//using loop to escape unmatching infinite loop
	var timePoint time.Time
	var loop int = 0
	for timePoint = startPoint; loop < 8; timePoint = timePoint.Add(day) {
		if timePoint.Weekday().String() == weekday {
			beego.Debug("find next weekday:" + timePoint.Weekday().String())
			return timePoint
		}
		loop++
	}
	panic("unable to indentify weekday:" + weekday)

}
