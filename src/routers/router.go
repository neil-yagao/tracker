package routers

import (
	"controllers"

	"github.com/astaxie/beego"
)

func init() {
	//beego.Router("/", &controllers.MainController{})
	//beego.Include(&master.OwnerController{})
	beego.Include(&controllers.WorkoutController{})
	beego.Include(&controllers.MovementController{})
	beego.SetStaticPath("/", "webapp")
}
