package routers

import (
	"controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Include(&controllers.WorkoutController{})
	beego.Include(&controllers.MovementController{})
	beego.SetStaticPath("/", "webapp")
	beego.SetStaticPath("/dist", "webapp/dist")
	beego.SetStaticPath("/node_modules", "webapp/node_modules")
}
