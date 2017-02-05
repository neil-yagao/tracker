package services

import "models"

type userWorkoutService struct {
}

type Workout struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Target      string `json:"target"`
	PerformDate string `json:"perform_date"`
	Description string `json:"description"`
}

var UserWorkoutService userWorkoutService

func (this *Workout) FindUserWorkouts(user string) []*Workout {
	return findUserWorkouts(user)

}

func findUserWorkouts(user string) []*Workout {
	rows := models.BasicCRUD.BuildAndQuery(QUERY_UESER_WORKOUT_QUERY, map[string]interface{}{"useridentity": user})
	defer rows.Close()
	return scanWorkoutsResult(rows)
}

func assignWorkoutsToUser(userIdentity string, workouts []*Workout) {
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
