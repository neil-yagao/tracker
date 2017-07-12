package models

type LoginInfo struct {
	UserInfo
	Template string `json:"template"`
}
