package services

import (
	"models"
	"time"
)

type WorkoutSerivces struct {
}

/**
* now we only generate workout for 4 times
 */
func (this *WorkoutSerivces) CreateWorkoutsFromeTemplate(template models.WorkoutTemplate) {

	workouts := createAndSaveWorkouts(template)
	for _, workout := range workouts {
		createAndSaveMovementForWorkout(workout, template)
	}
}

func createAndSaveMovementForWorkout(workout models.Workout, template models.WorkoutTemplate) {
	var containMovements []models.MovementTemplate = template.Movements
	for _, mv := range containMovements {
		movement := new(models.Movement)
		movement.Description = mv.Name
		movement.TargetMuscle = template.TargetMuscle
		movement.SecondaryMuscle = ""
		movement.Name = mv.Name

	}

}

func getMovementId(movement models.Movement) int64 {
	//insert movement if name does not exsited
	//if exsited then return exsiting movement id
	querySql := models.QueryBuilder.BuildQuery("select id from movement where name = :name",
		map[string]interface{}{"name": movement.Name})
	rows := models.BasicCRUD.Query(querySql)
	var id int64
	if rows.Next() {

		rows.Scan(&id)
	} else {
		id = models.BasicCRUD.Save("movement", movement)

	}
	return id
}

func createAndSaveWorkouts(template models.WorkoutTemplate) []models.Workout {
	var timePoint time.Time
	timePoint, _ = time.Parse("2016-11-17", template.StartAt)
	workouts := make([]models.Workout, 0)
	for i := 0; i <= 4; i++ {
		workout := new(models.Workout)
		workout.Name = template.Name
		workout.PerformDate = timePoint.String()[0:10]
		workout.Description = template.Description
		workout.Target = template.TargetMuscle
		// add once perweek
		duration, _ := time.ParseDuration("168h")
		timePoint.Add(duration)
		workout.Id = models.BasicCRUD.Save("workout", workout)
		workouts = append(workouts, *workout)
	}
	return workouts
}
