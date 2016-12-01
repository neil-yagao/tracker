package controllers

type UserControllers struct {
	GeneralController
}

// @router /login [post]
func (this *UserControllers) UserLogin() {
	userInfo := new(map[string]interface{})
	this.ParseRequestBody(userInfo)
}
