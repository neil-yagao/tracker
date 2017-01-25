package models

type LoginInfo struct {
	UserInfo
	Template string `json:"template"`
}

type UserInfo struct {
	Id           int64  `json:"id"`
	Username     string `json:"username"`
	Useridentity string `json:"useridentity"`
}
