package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["controllers:MovementController"] = append(beego.GlobalControllerRouter["controllers:MovementController"],
		beego.ControllerComments{
			Method: "GetMovements",
			Router: `/movements`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["controllers:MovementController"] = append(beego.GlobalControllerRouter["controllers:MovementController"],
		beego.ControllerComments{
			Method: "AddNewMovement",
			Router: `/movement`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["controllers:NutritionController"] = append(beego.GlobalControllerRouter["controllers:NutritionController"],
		beego.ControllerComments{
			Method: "SaveUserPhysiqueInfo",
			Router: `/physique/:user`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["controllers:NutritionController"] = append(beego.GlobalControllerRouter["controllers:NutritionController"],
		beego.ControllerComments{
			Method: "GetUserPhysique",
			Router: `/physique/:user`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["controllers:PlanController"] = append(beego.GlobalControllerRouter["controllers:PlanController"],
		beego.ControllerComments{
			Method: "CreateNewPlan",
			Router: `/plan`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["controllers:PlanController"] = append(beego.GlobalControllerRouter["controllers:PlanController"],
		beego.ControllerComments{
			Method: "GetAllPlans",
			Router: `/plans`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["controllers:PlanController"] = append(beego.GlobalControllerRouter["controllers:PlanController"],
		beego.ControllerComments{
			Method: "GetPlan",
			Router: `/plan/?:plan`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["controllers:PlanController"] = append(beego.GlobalControllerRouter["controllers:PlanController"],
		beego.ControllerComments{
			Method: "AssignPlan",
			Router: `/plan/?:plan`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["controllers:PlanController"] = append(beego.GlobalControllerRouter["controllers:PlanController"],
		beego.ControllerComments{
			Method: "AddSessionToPlan",
			Router: `/plan/?:plan/sessions`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["controllers:PlanController"] = append(beego.GlobalControllerRouter["controllers:PlanController"],
		beego.ControllerComments{
			Method: "RemoveSessionFromPlan",
			Router: `/plan/?:plan/session/?:session`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["controllers:PlanController"] = append(beego.GlobalControllerRouter["controllers:PlanController"],
		beego.ControllerComments{
			Method: "AddMovementToSession",
			Router: `/session/?:session/movements`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["controllers:UserController"] = append(beego.GlobalControllerRouter["controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["controllers:WorkingSessionController"] = append(beego.GlobalControllerRouter["controllers:WorkingSessionController"],
		beego.ControllerComments{
			Method: "FindUserSession",
			Router: `/session/?:username`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["controllers:WorkingSessionController"] = append(beego.GlobalControllerRouter["controllers:WorkingSessionController"],
		beego.ControllerComments{
			Method: "GetSessionDetail",
			Router: `/session-detail/?:sessionId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["controllers:WorkingSessionController"] = append(beego.GlobalControllerRouter["controllers:WorkingSessionController"],
		beego.ControllerComments{
			Method: "SettleSession",
			Router: `/session/?:id`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["controllers:WorkingSessionController"] = append(beego.GlobalControllerRouter["controllers:WorkingSessionController"],
		beego.ControllerComments{
			Method: "DoneOneMovement",
			Router: `/session-movement/?:movementId`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
