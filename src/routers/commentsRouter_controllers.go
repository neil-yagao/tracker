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

	beego.GlobalControllerRouter["controllers:WorkingSessionController"] = append(beego.GlobalControllerRouter["controllers:WorkingSessionController"],
		beego.ControllerComments{
			Method: "GetWorkoutSessions",
			Router: `/session/?:workout`,
			AllowHTTPMethods: []string{"get"},
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
