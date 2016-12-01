package services

import (
	"models"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type WorkoutSerivces struct {
}

var WorkoutCreator WorkoutSerivces
var week, _ = time.ParseDuration("168h")

/**
* now we only generate workout for 4 times
 */

func (this *WorkoutSerivces) CreateWorkoutsFromeTemplate(template models.WorkoutTemplate) {
	var extraWeight float32 = 0
	addition, _ := strconv.ParseFloat(template.Addition, 32)
	workouts := createAndSaveWorkouts(template)
	for _, workout := range workouts {
		generateWorkingSets(workout, template, extraWeight)
		extraWeight += float32(addition)
	}
}

const WORKOUT_QUERY string = "select w.id, w.name, w.target, w.perform_date, w.description from workout w, user_workout uw, user u " +
	"where is_finalized = 0 and uw.user = u.id and uw.workout = w.id and u.useridentity = :useridentity order by perform_date ASC"

func (this *WorkoutSerivces) FindUserWorkouts(user string) []*models.Workout {
	rows := models.BasicCRUD.BuildAndQuery(WORKOUT_QUERY, map[string]interface{}{"useridentity": user})
	workouts := make([]*models.Workout, 0)
	for rows.Next() {
		one := new(models.Workout)
		rows.Scan(&one.Id, &one.Name, &one.Target, &one.PerformDate, &one.Description)
		workouts = append(workouts, one)
	}
	return workouts
}

func generateWorkingSets(workout models.Workout, template models.WorkoutTemplate, extraWeight float32) {
	var containMovements []models.MovementTemplate = template.Movements
	for _, mv := range containMovements {
		movement := new(models.Movement)
		movement.Description = mv.Name
		movement.TargetMuscle = template.TargetMuscle
		movement.SecondaryMuscle = ""
		movement.Name = mv.Name
		movmentId := getMovementId(*movement)
		sets, err := strconv.Atoi(mv.Sets)
		handleIncorrectFormattingNumber(err)
		targetWeight, err := strconv.ParseFloat(mv.Weight, 32)
		for sequence := 1; sequence <= sets; sequence++ {
			workingSet := new(models.WorkingSet)
			workingSet.Movement = movmentId
			workingSet.Workout = workout.Id
			workingSet.AcheiveNumber = 0
			workingSet.TargetNumber, err = strconv.Atoi(mv.Repeats)
			handleIncorrectFormattingNumber(err)
			workingSet.TargetWeight = targetWeight + float64(extraWeight)
			workingSet.Sequence = int8(sequence)
			handleIncorrectFormattingNumber(err)
			models.BasicCRUD.Save("working_set", *workingSet)
		}

	}
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
	timePoint, _ = time.Parse("2006-01-02", template.StartAt)
	workouts := make([]models.Workout, 0)

	//TODO remove hard code 4 in the future
	//TODO let user input repeat times
	for i := 0; i < 4; i++ {
		logs.Info("current time Point:%v", timePoint)
		workout := new(models.Workout)
		workout.Name = template.Name
		workout.PerformDate = timePoint.String()[0:10]
		workout.Description = template.Description
		workout.Target = template.TargetMuscle
		// add once perweek
		timePoint = timePoint.Add(week)
		workout.Id = models.BasicCRUD.Save("workout", *workout)
		workouts = append(workouts, *workout)
	}
	return workouts
}
