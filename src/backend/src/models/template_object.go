package models

type SessionMovement struct {
	Id         int64     `json:"id" orm:"auto;pk"`
	Movement   *Movement `json:"movement" orm:"rel(fk)"`
	Sets       float64   `json:"sets" orm:"digits(8);decimals(2)"`
	Reps       float64   `json:"reps" orm:"digits(8);decimals(2)"`
	Sequence   int64     `json:"sequence" orm:"digits(8)"`
	NeedWarmup int8      `json:"needWarmup" orm:"digits(4)"`
	Session    *Session  `orm:"rel(fk)"`
}

type Session struct {
	Id           int64              `json:"id" orm:"auto;pk"`
	Name         string             `json:"name" orm:"size(128)"`
	TargetMuscle string             `json:"targetMuscle" orm:"size(128)"`
	Workouts     []*SessionMovement `json:"workouts" orm:"reverse(many)"`
	Repeat       int64              `json:"repeat" orm:"digits(8)"`
	Weekly       string             `json:"weekly"  orm:"size(32)"`
	Plan         *Plan              `orm:"rel(fk)"`
}

type Plan struct {
	Id       int64      `json:"id" orm:"auto;pk"`
	Name     string     `json:"name" orm:"size(128)"`
	Sessions []*Session `json:"sessions" orm:"reverse(many)"`
	CreateBy *UserInfo  `json:"createby" orm:"rel(fk)"`
}
