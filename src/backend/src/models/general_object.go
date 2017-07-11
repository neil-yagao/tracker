package models

type Movement struct {
	Id              int64  `json:"id" orm:"auto;pk"`
	TargetMuscle    string `json:"targetMuscle" orm:"size(32)"`
	SecondaryMuscle string `json:"secondarMuscle" orm:"size(32)"`
	Name            string `json:"name" orm:"size(128)"`
	Description     string `json:"description" orm:"size(2048)"`
	Dividable       int8   `json:"dividable" orm:"digits(4)"`
}

type UserInfo struct {
	Id           int64  `json:"id" orm:"auto;pk"`
	Username     string `json:"username" orm:"size(128)"`
	Useridentity string `json:"useridentity"  orm:"size(128)"`
}
