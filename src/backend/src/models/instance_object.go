package models

import "time"

type AssignedPlan struct {
	Id              int64     `orm:"auto;pk"`
	AssignTo        *UserInfo `orm:"rel(fk)"`
	AssignedPlan    *Plan     `orm:"rel(fk)"`
	AssignTime      time.Time `orm:"auto_now_add;type(date)"`
	ExecutingStatus string    `orm:"null;size(16)"`
}

type UserSession struct {
	Id            int64             `orm:"auto;pk"`
	AssignTo      *UserInfo         `orm:"rel(fk)"`
	ExpectingDate time.Time         `orm:"type(date)"`
	ExecutionDate time.Time         `orm:"null;type(date)"`
	Status        string            `orm:"size(16)"`
	Workouts      []*SessionWorkout `orm:"reverse(many)"`
}

type SessionWorkout struct {
	Id             int64            `orm:"auto;pk"`
	BelongSession  *UserSession     `orm:"rel(fk)"`
	MappedMovement *SessionMovement `orm:"rel(fk)"`
	Exercises      []*Exercise      `orm:"reverse(many)"`
}

type Exercise struct {
	Id            int64           `orm:"auto;pk"`
	BelongWorkout *SessionWorkout `orm:"rel(fk)"`
	Weight        float64         `orm:"digits(8);decimals(2)"`
	Reps          int8            `orm:"digits(8)"`
	Sequence      int8            `orm:"digits(8)"`
}
