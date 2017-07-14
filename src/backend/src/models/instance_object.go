package models

import "time"

//Plan instance that assigned to user
type AssignedPlan struct {
	Id              int64     `orm:"auto;pk"`
	AssignTo        *UserInfo `orm:"rel(fk)"`
	AssignedPlan    *Plan     `orm:"rel(fk)"`
	AssignTime      time.Time `orm:"auto_now_add;type(date)"`
	ExecutingStatus string    `orm:"null;size(16)"`
}

//Session instance that generated after assign plan to user
type UserSession struct {
	Id            int64             `orm:"auto;pk"`
	AssignTo      *UserInfo         `orm:"rel(fk)"`
	ExpectingDate time.Time         `orm:"type(date)"`
	ExecutionDate time.Time         `orm:"null;type(date)"`
	OriginSession *Session          `orm:"rel(fk)"`
	Status        string            `orm:"size(16)"`
	Workouts      []*SessionWorkout `orm:"reverse(many)"`
}

//workout instance for each assigned session
//mapping to each movement
type SessionWorkout struct {
	Id             int64            `orm:"auto;pk"`
	BelongSession  *UserSession     `orm:"rel(fk)"`
	MappedMovement *SessionMovement `orm:"rel(fk)"`
	Exercises      []*Exercise      `orm:"reverse(many)"`
	Status         string           `orm:"size(4)"`
}

//exercise instance for each movement
//this is actually each user has achieved set
type Exercise struct {
	Id            int64           `orm:"auto;pk"`
	BelongWorkout *SessionWorkout `orm:"rel(fk)"`
	Weight        float64         `orm:"digits(8);decimals(2)"`
	Reps          int8            `orm:"digits(8)"`
	Sequence      int8            `orm:"digits(8)"`
}
