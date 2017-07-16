package models

import "time"

//Plan instance that assigned to user
type AssignedPlan struct {
	Id              int64     `orm:"auto;pk" json:"id"`
	AssignTo        *UserInfo `orm:"rel(fk)" json:"assignTo"`
	AssignedPlan    *Plan     `orm:"rel(fk)" json:"assignedPlan"`
	AssignTime      time.Time `orm:"auto_now_add;type(date)" json:"assignTime"`
	ExecutingStatus string    `orm:"null;size(16)" json:"executingStatus"`
}

//Session instance that generated after assign plan to user
type UserSession struct {
	Id            int64             `orm:"auto;pk" json:"id"`
	AssignTo      *UserInfo         `orm:"rel(fk)" json:"assignTo"`
	ExpectingDate time.Time         `orm:"type(date)" json:"expectingDate"`
	ExecutionDate time.Time         `orm:"null;type(date)" json:"executionDate"`
	Status        string            `orm:"size(16)" json:"status"`
	Workouts      []*SessionWorkout `orm:"reverse(many)" json:"workouts"`
}

//workout instance for each assigned session
//mapping to each movement
type SessionWorkout struct {
	Id             int64            `orm:"auto;pk" json:"id"`
	BelongSession  *UserSession     `orm:"rel(fk)" json:"belongSession"`
	MappedMovement *SessionMovement `orm:"rel(fk)" json:"mappedMovement"`
	Exercises      []*Exercise      `orm:"reverse(many)" json:"exercises"`
}

//exercise instance for each movement
//this is actually each user has achieved set
type Exercise struct {
	Id            int64           `orm:"auto;pk" json:"id"`
	BelongWorkout *SessionWorkout `orm:"rel(fk)" json:"belongWorkout"`
	Weight        float64         `orm:"digits(8);decimals(2)" json:"weight"`
	Reps          int8            `orm:"digits(8)" json:"reps"`
	Sequence      int8            `orm:"digits(8)" json:"sequence"`
}
