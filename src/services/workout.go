package services

import (
	"fmt"
	"math"
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
var day, _ = time.ParseDuration("24h")

/**
* now we only generate workout for 4 times
 */

func (this *WorkoutSerivces) CreateWorkoutsFromeTemplate(template models.WorkoutTemplate, userIdentity string) {
	var extraWeight float32 = 0
	addition, _ := strconv.ParseFloat(template.Addition, 32)
	workouts := createAndSaveWorkouts(template)
	user := UserService.GetUserIdFromIdentity(userIdentity)
	for _, workout := range workouts {
		UserService.AssignWorkoutToUser(user, workout.Id)
		generateWorkingSets(workout, template, extraWeight)
		extraWeight += float32(addition)
	}
}

const WORKOUT_QUERY string = "select w.id, w.name, w.target, w.perform_date, w.description from workout w, user_workout uw, user u " +
	"where is_finalized = 0 and uw.user = u.id and uw.workout = w.id and u.useridentity = :useridentity order by perform_date ASC"

func (this *WorkoutSerivces) FindUserWorkouts(user string) []*models.Workout {
	rows := models.BasicCRUD.BuildAndQuery(WORKOUT_QUERY, map[string]interface{}{"useridentity": user})
	defer rows.Close()
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
		movementId := getMovementId(*movement)
		if mv.NeedWarmSet > 0 {
			generateWarmingSet(movementId, workout.Id, mv.Repeats,
				mv.Weight+float64(extraWeight))
		}
		//generate working set for movement
		for sequence := 1; sequence <= mv.Sets; sequence++ {
			assignWorkingSet(movementId, workout.Id, mv.Repeats,
				mv.Weight+float64(extraWeight), sequence)
		}
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
	//if assume addition is larger than 10
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
	for workingWeight := startWeight; workingWeight <= targetWeight*0.9; workingWeight += realAddition {

		assignWorkingSet(movementId, workout, targetNumber, transferToUsableWeight(workingWeight), 0)
	}
}

// return a weigth that round down to a integer times of 2.5

func transferToUsableWeight(original float64) float64 {
	integer := math.Floor(original / float64(2.5))
	return integer * 2.5
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

func createAndSaveWorkouts(template models.WorkoutTemplate) []models.Workout {
	var timePoint time.Time
	timePoint, _ = time.Parse("2006-01-02", template.StartAt)
	workouts := make([]models.Workout, 0)
	timePoint = findNextWeekday(template.Weekly, timePoint)
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

func findNextWeekday(weekday string, startPoint time.Time) time.Time {
	//using loop to escape unmatching infinite loop
	var timePoint time.Time
	var loop int = 0
	for timePoint = startPoint; loop < 8; timePoint = timePoint.Add(day) {
		if timePoint.Weekday().String() == weekday {
			beego.Debug("find next weekday:" + timePoint.Weekday().String())
			return timePoint
		}
		loop++
	}
	panic("unable to indentify weekday:" + weekday)

}
