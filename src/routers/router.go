package routers

import (
	"controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Include(&controllers.WorkoutController{})
	beego.Include(&controllers.MovementController{})
	beego.SetStaticPath("/", "webapp")
}
