package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["controllers:WorkingSessionController"] = append(beego.GlobalControllerRouter["controllers:WorkingSessionController"],
		beego.ControllerComments{
			Method: "FindUserSession",
			Router: `/session/?:username`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
