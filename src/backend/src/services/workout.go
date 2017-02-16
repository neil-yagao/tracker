package services

import (
	"math"
	"models"

	"github.com/astaxie/beego"
)

const MINIMAL_WEIGHT float64 = 2.5
const DEFAULT_WORKOUT_REPEAT_TIME = 4

// return a weigth that round down to a integer times of 2.5
func transferToUsableWeight(original float64) float64 {
	integer := math.Floor(original / MINIMAL_WEIGHT)
	return integer * MINIMAL_WEIGHT
}

func assignWorkingSet(movementId int64, workout int64, targetNumber int, targetWeight float64, sequence int) {
	workingSet := new(models.WorkingSet)
	workingSet.Movement = movementId
	workingSet.Workout = workout
	workingSet.AcheiveNumber = 0
	workingSet.TargetNumber = targetNumber
	workingSet.TargetWeight = targetWeight
	workingSet.Sequence = int8(sequence)
	models.BasicCRUD.Save("working_set", *workingSet)
}

func handleIncorrectFormattingNumber(err error) {
	if err != nil {
		beego.Error("incorrect input during parsing number")
		panic(err)
	}
}

func getMovementId(movement models.Movement) int64 {
	//insert movement if name does not exsited
	//if exsited then return exsiting movement id
	querySql := models.QueryBuilder.BuildQuery("select id from movement where name = :name",
		map[string]interface{}{"name": movement.Name})
	rows := models.BasicCRUD.Query(querySql)
	defer rows.Close()
	var id int64
	if rows.Next() {

		rows.Scan(&id)
	} else {
		id = models.BasicCRUD.Save("movement", movement)

	}
	return id
}

func createAndSaveWorkouts(template *WorkoutDefinition) []*Workout {
	workouts := make([]*Workout, 0)
	//TODO remove hard code 4 in the future
	//TODO let user input repeat times
	var i float64 = 0.0
	for ; i < DEFAULT_WORKOUT_REPEAT_TIME; i++ {
		workout := new(Workout)
		workout.Name = template.Name
		workout.Description = template.Description
		workout.Target = template.TargetMuscle
		workout.Sequence = int64(i)
		workout.Repeating = template.Weekly
		workout.Id = models.BasicCRUD.Save("workout", *workout)
		workouts = append(workouts, workout)
	}
	return workouts
}
