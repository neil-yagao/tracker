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

	beego.GlobalControllerRouter["controllers:MovementController"] = append(beego.GlobalControllerRouter["controllers:MovementController"],
		beego.ControllerComments{
			Method: "InsertMovement",
			Router: `/movements`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["controllers:MovementController"] = append(beego.GlobalControllerRouter["controllers:MovementController"],
		beego.ControllerComments{
			Method: "UpdateMovement",
			Router: `/movement/:id`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["controllers:TemplateController"] = append(beego.GlobalControllerRouter["controllers:TemplateController"],
		beego.ControllerComments{
			Method: "GetTemplates",
			Router: `/templates`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["controllers:TemplateController"] = append(beego.GlobalControllerRouter["controllers:TemplateController"],
		beego.ControllerComments{
			Method: "AssignTemplates",
			Router: `/templates/user`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["controllers:TemplateController"] = append(beego.GlobalControllerRouter["controllers:TemplateController"],
		beego.ControllerComments{
			Method: "UnasignTemplates",
			Router: `/templates/user`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["controllers:TemplateController"] = append(beego.GlobalControllerRouter["controllers:TemplateController"],
		beego.ControllerComments{
			Method: "GetUserCurrentTemplates",
			Router: `/templates/:useridentity`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["controllers:UserControllers"] = append(beego.GlobalControllerRouter["controllers:UserControllers"],
		beego.ControllerComments{
			Method: "UserLogin",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["controllers:WorkingSessionController"] = append(beego.GlobalControllerRouter["controllers:WorkingSessionController"],
		beego.ControllerComments{
			Method: "GetWorkoutSessions",
			Router: `/session/?:workout`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["controllers:WorkingSessionController"] = append(beego.GlobalControllerRouter["controllers:WorkingSessionController"],
		beego.ControllerComments{
			Method: "SaveSession",
			Router: `/session/?:workout`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["controllers:WorkingSessionController"] = append(beego.GlobalControllerRouter["controllers:WorkingSessionController"],
		beego.ControllerComments{
			Method: "UpdateSessionMovement",
			Router: `/session/?:workout/movement`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["controllers:WorkoutController"] = append(beego.GlobalControllerRouter["controllers:WorkoutController"],
		beego.ControllerComments{
			Method: "GetWorkouts",
			Router: `/workouts`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["controllers:WorkoutController"] = append(beego.GlobalControllerRouter["controllers:WorkoutController"],
		beego.ControllerComments{
			Method: "InsertWorkout",
			Router: `/workouts`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

}
