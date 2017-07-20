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
	ExecutionDate time.Time         `orm:"null;type(date)" `
	OriginSession *Session          `orm:"rel(fk)" json:"originSession"`
	Status        string            `orm:"size(16)"`
	Workouts      []*SessionWorkout `orm:"reverse(many)"  json:"workouts"`
}

func (u *UserSession) TableUnique() [][]string {
	return [][]string{
		[]string{"ExpectingDate", "OriginSession"}}
}

//workout instance for each assigned session
//mapping to each movement
type SessionWorkout struct {
	Id             int64            `orm:"auto;pk" json:"id"`
	BelongSession  *UserSession     `orm:"rel(fk)"`
	MappedMovement *SessionMovement `orm:"rel(fk)"  json:"mappedMovement"`
	Exercises      []*Exercise      `orm:"reverse(many)"`
	Status         string           `orm:"size(4)" json:"status"`
}

//exercise instance for each movement
//this is actually each user has achieved set
type Exercise struct {
	Id            int64           `orm:"auto;pk" json:"id"`
	BelongWorkout *SessionWorkout `orm:"rel(fk)" json:"belongWorkout"`
	Weight        float64         `orm:"digits(8);decimals(2)" json:"weight"`
	Reps          int8            `orm:"digits(8)" json:"reps"`
}
