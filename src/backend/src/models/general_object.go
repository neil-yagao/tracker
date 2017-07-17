package models

//general training movement
type Movement struct {
	Id              int64     `json:"id" orm:"auto;pk"`
	TargetMuscle    string    `json:"targetMuscle" orm:"size(32)"`
	SecondaryMuscle string    `json:"secondaryMuscle" orm:"size(128)"`
	Name            string    `json:"name" orm:"size(128)"`
	Description     string    `json:"description" orm:"type(text)"`
	Dividable       int8      `json:"dividable" orm:"digits(4)"`
	AddBy           *UserInfo `json:"addBy" orm:"rel(fk)"`
}

func (u *Movement) TableUnique() [][]string {
	return [][]string{
		[]string{"Name"},
	}
}

type UserInfo struct {
	Id           int64  `json:"id" orm:"auto;pk"`
	Username     string `json:"username" orm:"size(128)"`
	Useridentity string `json:"useridentity"  orm:"size(128)"`
}

func (u *UserInfo) TableUnique() [][]string {
	return [][]string{
		[]string{"Username"},
	}
}
