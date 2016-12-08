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

<<<<<<< HEAD
=======
	beego.GlobalControllerRouter["controllers:WorkingSessionController"] = append(beego.GlobalControllerRouter["controllers:WorkingSessionController"],
		beego.ControllerComments{
			Method: "UpdateSessionMovement",
			Router: `/session/?:workout/movement`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

>>>>>>> 3b1011b7ed0b0ec91e9383bd1761c8be17f0ca0e
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
