package services

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"models"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/astaxie/beego"
)

func init() {
	beego.SetLevel(beego.LevelError)
	user := new(models.LoginInfo)
	user.Useridentity = USER_INDENTITY
	user.Username = "testing"
	UserService.HandleUserLogin(*user)
}

var TEST_TEMPLATE_RAW string = "{\"template\":\"testing\",\"name\":\"template name\",\"movements\":[{\"name\":\"Bench Press\",\"weight\":\"85\",\"repeats\":\"6\",\"sets\":\"3\"},{\"name\":\"Barbell Pull\",\"weight\":\"65\",\"repeats\":\"6\",\"sets\":\"3\"},{\"name\":\"Sitting Shoulder Press\",\"weight\":\"25\",\"repeats\":\"10\",\"sets\":\"3\"},{\"name\":\"Pull Up\",\"weight\":\"76\",\"repeats\":\"10\",\"sets\":\"3\"},{\"name\":\"Incline Bench Press\",\"weight\":\"21\",\"repeats\":\"10\",\"sets\":\"3\"},{\"name\":\"Barbell Curl\",\"weight\":\"35\",\"repeats\":\"10\",\"sets\":\"3\"},{\"name\":\"Dip\",\"weight\":\"76\",\"repeats\":\"10\",\"sets\":\"3\"}],\"startAt\":\"2016-11-16\",\"addition\":\"2.5\",\"weekly\":\"Tuesday\",\"targetMuscle\":\"Upper body\",\"description\":\"Here I input some note for myself\"}"

var username = []byte("testing")

var USER_INDENTITY string = fmt.Sprintf("%x", md5.Sum(username))

func TestCreateAndSaveWorkoutTemplateGivenCase(t *testing.T) {
	var result models.WorkoutTemplate
	json.Unmarshal([]byte(TEST_TEMPLATE_RAW), &result)
	workouts := createAndSaveWorkouts(result)
	if len(workouts) > 0 {
		defer cleanUp("workout", workouts)
	}
	if len(workouts) != 4 {
		t.Error("generate incorrect number of workouts from template")
	}
	templateId := findTemplateId("testing")
	if templateId == -1 {
		t.Error("fail to find template")
	} else {
		defer cleanUp("template", templateId)
	}

	assignWorkoutsToTemplate("testing", workouts)
	workoutsFromTemplate := findTemplateWorkouts("testing")
	if len(workoutsFromTemplate) > 0 {
		defer cleanUp("workout", workoutsFromTemplate)
	}
	//assert all matching
	if assertWorkoutsMatching(workouts, workoutsFromTemplate) {
		t.Log("Test case create workout with template pass.\n")
	} else {
		t.Error("Test case create workout with template failed.\n")
	}

}

func TestCreateAndSaveWorkoutNoTemplateCase(t *testing.T) {
	var result models.WorkoutTemplate
	json.Unmarshal([]byte(TEST_TEMPLATE_RAW), &result)
	result.Template = ""
	workouts := createAndSaveWorkouts(result)
	if len(workouts) > 0 {
		defer cleanUp("workout", workouts)
	}
	assignWorkoutsToUser(USER_INDENTITY, workouts)
	workoutsFromUser := findUserWorkouts(USER_INDENTITY)
	//assert all matching
	if len(workoutsFromUser) > 0 {
		defer cleanUp("workout", workoutsFromUser)
	}
	if assertWorkoutsMatching(workouts, workoutsFromUser) {
		t.Log("Test case create workout with no template pass.\n")
	} else {
		t.Error("Test case create workout with no template failed.\n")
	}

}

func assertWorkoutsMatching(workouts []*models.Workout, workoutForCompare []*models.Workout) bool {
	matching := 0
	for _, workout := range workouts {
		if workout.Id == 0 {
			return false
		}
		for _, wtc := range workoutForCompare {
			if wtc.Id == workout.Id {
				matching++
				break
			}
		}
	}
	return len(workouts) == matching
}

func cleanUp(table string, objects interface{}) {

	switch reflect.TypeOf(objects).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(objects)

		for i := 0; i < s.Len(); i++ {
			values := s.Index(i).Elem()
			match := func(name string) bool {
				return strings.ToLower(name) == "id"
			}
			id := values.FieldByNameFunc(match).Int()
			beego.Debug("delete from talbe ", table, " id ", id)

			models.BasicCRUD.Delete(table, id)
		}
	case reflect.Int64:
		models.BasicCRUD.Delete(table, objects.(int64))
	}

}

func TestTimeParse(t *testing.T) {
	timePoint, err := time.Parse("2006-01-02", "2016-11-21")
	if err != nil {
		t.Error(err)
	}
	t.Log(timePoint.String()[0:10])
}
