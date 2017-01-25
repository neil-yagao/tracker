package services

import "models"

type userWorkoutService struct {
}

var UserWorkoutService userWorkoutService

func (this *userWorkoutService) FindUserWorkouts(user string) []*models.Workout {
	return findUserWorkouts(user)

}

func findUserWorkouts(user string) []*models.Workout {
	rows := models.BasicCRUD.BuildAndQuery(QUERY_UESER_WORKOUT_QUERY, map[string]interface{}{"useridentity": user})
	defer rows.Close()
	return scanWorkoutsResult(rows)
}

func assignWorkoutsToUser(userIdentity string, workouts []*models.Workout) {
	//templateName :=
	user := UserService.GetUserIdFromIdentity(userIdentity)
	for _, workout := range workouts {
		assignWorkoutToUser(user, workout.Id)
	}
}

type UserWorkout struct {
	User    int64
	Workout int64
}

func assignWorkoutToUser(user int64, workout int64) {
	models.BasicCRUD.Save("user_workout", UserWorkout{user, workout})
}
