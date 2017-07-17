package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["controllers:MovementController"] = append(beego.GlobalControllerRouter["controllers:MovementController"],
		beego.ControllerComments{
			Method: "GetMovements",
			Router: `/movements`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["controllers:PlanController"] = append(beego.GlobalControllerRouter["controllers:PlanController"],
		beego.ControllerComments{
			Method: "CreateNewPlan",
			Router: `/plan`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["controllers:PlanController"] = append(beego.GlobalControllerRouter["controllers:PlanController"],
		beego.ControllerComments{
			Method: "GetAllPlans",
			Router: `/plans`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["controllers:PlanController"] = append(beego.GlobalControllerRouter["controllers:PlanController"],
		beego.ControllerComments{
			Method: "GetPlan",
			Router: `/plan/?:plan`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["controllers:WorkingSessionController"] = append(beego.GlobalControllerRouter["controllers:WorkingSessionController"],
		beego.ControllerComments{
			Method: "FindUserSession",
			Router: `/session/?:username`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
