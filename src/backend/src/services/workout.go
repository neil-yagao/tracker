package services

import (
	"fmt"
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
		createWorkout(workout, template.Movements)
		workouts = append(workouts, workout)
	}
	return workouts
}

func createWorkout(workout *Workout, movements []models.MovementTemplate) {
	workout.Id = models.BasicCRUD.Save("workout", *workout)
	batchInsert := new(models.BatchSqls)

	for _, mv := range movements {
		movement := new(models.Movement)
		movement.Description = mv.Name
		movement.SecondaryMuscle = ""
		movement.Name = mv.Name
		movementId := getMovementId(*movement)
		batchInsert.Params = append(batchInsert.Params, map[string]interface{}{
			"workout":       workout.Id,
			"movement":      movementId,
			"target_weight": mv.Weight,
			"sets":          mv.Sets,
			"repeats":       mv.Repeats})
	}
	batchInsert.BatchInsert("workout_movement")
}

func generateWorkingSets(workout Workout, extraWeight float64) {
	containMovements := getWorkoutMovements(workout.Id)
	for _, mv := range containMovements {
		targetWeight := mv.Weight + float64(extraWeight)
		if mv.NeedWarmSet > 0 {
			generateWarmingSet(mv.Movement, workout.Id, mv.Repeats, targetWeight)
		}
		//generate working set for movement
		batchInsert := new(models.BatchSqls)
		for sequence := 1; sequence <= mv.Sets; sequence++ {
			workingSet := new(models.WorkingSet)
			workingSet.Movement = mv.Movement
			workingSet.Workout = workout.Id
			workingSet.AcheiveNumber = 0
			workingSet.TargetNumber = mv.Repeats
			workingSet.TargetWeight = targetWeight
			workingSet.Sequence = int8(sequence)
			jsonValue := models.InterfaceToJSONObject(workingSet)
			delete(jsonValue, "movementName")
			batchInsert.Params = append(batchInsert.Params, jsonValue)
		}
		batchInsert.BatchInsert("working_set")
	}
}

func generateWarmingSet(movementId int64, workout int64, targetNumber int, targetWeight float64) {
	rows := models.BasicCRUD.Query("select dividable from movement where id = " + fmt.Sprint(movementId))
	var movementDividable int
	for rows.Next() {
		rows.Scan(&movementDividable)
	}
	var startWeight float64
	// if movement is dividable
	// warm set start from 20KG
	// else warm set start at 25%
	if movementDividable == 1 {
		startWeight = 20
	} else {
		startWeight = targetWeight * 0.2
	}
	var assumeAddition float64 = targetWeight * 0.15
	// if assume addition is larger than 10
	// then using 10 as real addition
	// else if assume addtion is smaller than 10 but larger than 5
	// then using 5 as real addition
	// else using half of the targetWeight
	var realAddition float64
	if assumeAddition >= 10 {
		realAddition = 10
	} else if assumeAddition >= 5 && assumeAddition < 10 {
		realAddition = 5
	} else {
		realAddition = (targetWeight - startWeight) / 2
	}
	batchInsert := new(models.BatchSqls)

	for workingWeight := startWeight; workingWeight <= targetWeight*0.9; workingWeight += realAddition {
		workingSet := new(models.WorkingSet)
		workingSet.Movement = movementId
		workingSet.Workout = workout
		workingSet.AcheiveNumber = 0
		workingSet.TargetNumber = targetNumber
		workingSet.TargetWeight = targetWeight
		workingSet.Sequence = 0
		jsonValue := models.InterfaceToJSONObject(workingSet)
		delete(jsonValue, "movementName")
		batchInsert.Params = append(batchInsert.Params, jsonValue)
	}
	batchInsert.BatchInsert("working_set")
}
