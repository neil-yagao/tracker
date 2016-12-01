package models

type LoginInfo struct {
	Id           int64  `json:"id"`
	Username     string `json:"username"`
	Useridentity string `json:"useridentity"`
}
