package routers

import (
	"controllers"
	"controllers/master"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Include(&master.OwnerController{})
	beego.SetStaticPath("/", "webapp")
}
