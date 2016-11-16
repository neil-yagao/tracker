package routers

import (
	"controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Include(&master.OwnerController{})
	beego.SetStaticPath("/", "webapp")
}
