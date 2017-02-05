package services

import "models"

type userService struct {
}

var UserService userService

func (this *userService) HandleUserLogin(user models.LoginInfo) {
	//currently just insert user info if not exsited
	//do nothing if user exsited
	if !userExisted(user.Useridentity) {
		models.BasicCRUD.Save("user", user)
	}

}

const USER_EXISTED_QUERY string = "select count(1) from user where useridentity = :useridentity"

func userExisted(userIdentity string) bool {
	rows := models.BasicCRUD.BuildAndQuery(USER_EXISTED_QUERY, map[string]interface{}{"useridentity": userIdentity})
	defer rows.Close()
	if rows.Next() {
		var result int
		rows.Scan(&result)
		return result > 0
	}
	return false
}

func (this *userService) GetUserIdFromIdentity(user string) int64 {
	rows := models.BasicCRUD.Query("select id from user where useridentity = '" + user + "'")
	var id int64
	for rows.Next() {
		rows.Scan(&id)
	}
	return id
}
