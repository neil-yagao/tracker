package routers

import (
	"controllers"
	"github.com/astaxie/beego/config"

	"github.com/astaxie/beego"
)

func init() {
	//beego.Router("/", &controllers.MainController{})
	//beego.Include(&master.OwnerController{})

	beego.Include(&controllers.WorkingSessionController{})
	beego.Include(&controllers.MovementController{})
	beego.Include(&controllers.PlanController{})
	beego.Include(&controllers.NutritionController{})
	beego.Include(&controllers.UserController{})

	config, _ := config.NewConfig("ini", "conf/app.conf")
	baseStaticPath := config.String("static")
	beego.SetStaticPath("/", "../../web-mobile")
	beego.SetStaticPath("/dist", "../../web-mobile/dist")
	beego.SetStaticPath("/node_modules", "../../web-mobile/node_modules")
	beego.SetStaticPath("/static", baseStaticPath)
}
