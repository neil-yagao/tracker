package routers

import (
	"controllers"

	"github.com/astaxie/beego"
)

func init() {
	//beego.Router("/", &controllers.MainController{})
	//beego.Include(&master.OwnerController{})

	beego.Include(&controllers.WorkingSessionController{})
	beego.Include(&controllers.MovementController{})
	beego.Include(&controllers.PlanController{})

	beego.SetStaticPath("/", "../../webapp")
	beego.SetStaticPath("/mobile", "../../hybrid/www")
	beego.SetStaticPath("/dist", "../../webapp/dist")
	beego.SetStaticPath("/node_modules", "../../webapp/node_modules")
}
