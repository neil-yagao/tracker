package services

import "controllers"

type WorkoutSerivces struct {
}

/**
* now we only generate workout for 4 times
 */
func (this *WorkoutSerivces) CreateWorkoutsFromeTemplate(template controllers.WorkoutTemplate) {
	var containMovements []controllers.MovementTemplate = template.Movements
	for mv := range containMovements {
		workout := new(controllers.Workout)
        workout.Name = template.Name
        workout.PerformDate =
	}
}
