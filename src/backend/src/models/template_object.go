package models

//workout template information
//each session is composed by serveral movement and each movement has its own target reps and sets
type SessionMovement struct {
	Id         int64     `json:"id" orm:"auto;pk"`
	Movement   *Movement `json:"movement" orm:"rel(fk)"`
	Sets       int8      `json:"sets" orm:"digits(8)"`
	Reps       int8      `json:"reps" orm:"digits(8)"`
	Sequence   int64     `json:"sequence" orm:"digits(8)"`
	NeedWarmup int8      `json:"needWarmup" orm:"digits(4)"`
	Session    *Session  `orm:"rel(fk)"`
}

func (u *SessionMovement) TableUnique() [][]string {
	return [][]string{
		[]string{"Movement", "Session"},
	}
}

//session template information
//each plan is composed by several weekly(current setting) repeating sessions
type Session struct {
	Id           int64              `json:"id" orm:"auto;pk"`
	Name         string             `json:"name" orm:"size(128)"`
	TargetMuscle string             `json:"targetMuscle" orm:"size(128)"`
	Workouts     []*SessionMovement `json:"workouts" orm:"reverse(many)"`
	Repeat       int64              `json:"repeat" orm:"digits(8)"`
	Weekly       string             `json:"weekly"  orm:"size(32)"`
	Plan         *Plan              `orm:"rel(fk)"`
}

func (u *Session) TableUnique() [][]string {
	return [][]string{
		[]string{"Name", "Plan"},
	}
}

//plan template information
//could be assign to user which will create assigned plan object
type Plan struct {
	Id       int64      `json:"id" orm:"auto;pk"`
	Name     string     `json:"name" orm:"size(128)"`
	Sessions []*Session `json:"sessions" orm:"reverse(many)"`
	CreateBy *UserInfo  `json:"createby" orm:"rel(fk)"`
}

func (u *Plan) TableUnique() [][]string {
	return [][]string{
		[]string{"Name"},
	}
}
