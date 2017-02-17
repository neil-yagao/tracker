package services

import (
	"encoding/json"
	"fmt"
	"models"
	"reflect"
	"strings"
	"testing"

	"github.com/astaxie/beego"
)

const USER_WORKOUT_DEFINITION string = "{\"movements\":" +
	"[{\"name\":\"卧推\",\"repeats\":8,\"weight\":85,\"sets\":3}," +
	"{\"name\":\"上斜板哑铃卧推\",\"repeats\":8,\"weight\":25,\"sets\":3}]," +
	"\"name\":\"测试训练胸部\"," +
	"\"template\":\"测试训练计划\"," +
	"\"targetMuscle\":\"胸\"," +
	"\"description\":\"胸部训练\"," +
	"\"weekly\":\"Monday\"}"

const UPDATE_WORKOUT_DEFINITION string = "{\"movements\":" +
	"[{\"name\":\"杠铃划船\",\"repeats\":8,\"weight\":60,\"sets\":3}," +
	"{\"name\":\"反手引体向上\",\"repeats\":8,\"weight\":0,\"sets\":3}]," +
	"\"name\":\"测试背部训练\",\"template\":\"测试训练计划\",\"targetMuscle\":\"背,二头\",\"description\":\"测试训练计划，背部训练\",\"weekly\":\"Tuesday\"}"
	/**
	 * User create new template
	 * and add workout to it
	 */
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

func RecoverFromError() {
	if x := recover(); x != nil {
		beego.Error("run time panic:", x)
	}
}

func createTestingWorkout(templateStr string, t *testing.T) (WorkoutDefinition, *Template) {

	var workoutDefinition WorkoutDefinition
	json.Unmarshal([]byte(templateStr), &workoutDefinition)
	template := findTemplate(workoutDefinition.Template)
	t.Log("find tempalte:", template)
	if template != nil && template.Id > 0 {
		t.Log("find exsiting template for test case, delete")
		models.BasicCRUD.Delete("template", template.Id)
	}
	workoutDefinition.AssignWorkoutsToTemplate()
	template = findTemplate(workoutDefinition.Template)
	if template == nil {
		t.Error("create new template fail!")
		t.Fail()
	}
	return workoutDefinition, template
}

func TestIntegrationUserCreateNewTemplate(t *testing.T) {
	_, template := createTestingWorkout(USER_WORKOUT_DEFINITION, t)

	workouts := findWorkoutsFromTemplates([]int64{template.Id})
	if len(workouts) != 4 {
		t.Error("template do not create proper workouts")
		t.Fail()
	}
	sequenceMatching := make(map[string]bool, 0)
	for _, workout := range workouts {
		if workout.Repeating != "Monday" {
			t.Error("create workout repeat unmatch fail!")
			t.Error("error workout:", workout)
			t.Fail()
		}
		key := workout.Name + fmt.Sprint(workout.Sequence)
		if sequenceMatching[key] {
			t.Error("create workout sequence unmatch fail!")
			t.Error("error workout:", workout)
			t.Fail()
		} else {
			sequenceMatching[key] = true
		}
	}

	cleanUp("template", template.Id)
	cleanUp("workout", workouts)
	models.BasicCRUD.Update("delete from template_workout where template = " + fmt.Sprint(template.Id))
	t.Log("-------------------------------------------TestIntegrationUserCreateNewTemplate Passed-------------------------------")
}

/**
 * User add workout to exsiting template
 *
 */

func TestIntegrationUserUpdateExistingTemplate(t *testing.T) {
	createTestingWorkout(USER_WORKOUT_DEFINITION, t)
	var workoutDefinition WorkoutDefinition
	json.Unmarshal([]byte(UPDATE_WORKOUT_DEFINITION), &workoutDefinition)
	workoutDefinition.AssignWorkoutsToTemplate()
	template := findTemplate(workoutDefinition.Template)
	workouts := findWorkoutsFromTemplates([]int64{template.Id})
	if len(workouts) != 8 {
		t.Error("template do not create proper workouts")
		t.Fail()
	}
	sequenceMatching := make(map[string]bool, 0)
	for _, workout := range workouts {
		key := workout.Name + fmt.Sprint(workout.Sequence)
		if sequenceMatching[key] {
			t.Error("create workout sequence unmatch fail!")
			t.Error("error workout:", workout)
			t.Fail()
		} else {
			sequenceMatching[key] = true
		}
	}
	// cleanUp("template", template.Id)
	// cleanUp("workout", workouts)
	// for _, workout := range workouts {
	// 	models.BasicCRUD.Update("delete from working_set where workout = " + fmt.Sprint(workout.Id))
	// }
	// models.BasicCRUD.Update("delete from template_workout where template = " + fmt.Sprint(template.Id))
	t.Log("-------------------------------------------TestIntegrationUserUpdateExistingTemplate Passed-------------------------------")
}

/**
 * User assign templates to self
 *
 */
func TestIntegrationAssignTemplateToUser(t *testing.T) {
	_, template := createTestingWorkout(USER_WORKOUT_DEFINITION, t)
	testUser := new(models.UserInfo)
	testUser.Useridentity = USER_INDENTITY
	testUser.Username = string(username)
	UserService.HandleUserLogin(*testUser)
	userTempalte := new(UserTemplate)
	userTempalte.User = testUser.Username
	userTempalte.UserIdentity = testUser.Useridentity
	userTempalte.Templates = []string{template.Name}
	userTempalte.TemplatesId = []int64{template.Id}
	userTempalte.AssignTemplate()
	templateWorkouts := findWorkoutsFromTemplates([]int64{template.Id})
	userWorkouts := findUserWorkouts(userTempalte.UserIdentity)
	var matching int = 0
	for _, template := range templateWorkouts {
		t.Log("using template workout:", *template)
		for _, user := range userWorkouts {
			t.Log("examine user workout:", *user)
			if template.Id == user.Id {
				matching++
			}
		}
		sessionContent := SessionService.GetWorkoutSession(int(template.Id))
		if len(sessionContent) != 6 {
			t.Error("Unable to generate proper working set during assign")
			t.Fail()
		}
	}

	if matching != len(templateWorkouts) {
		t.Error("Unable to assign template to user")
	}

	cleanUp("template", template.Id)
	cleanUp("workout", templateWorkouts)
	for _, workout := range userWorkouts {
		models.BasicCRUD.Update("delete from working_set where workout = " + fmt.Sprint(workout.Id))
	}
	models.BasicCRUD.Update("delete from template_workout where template = " + fmt.Sprint(template.Id))
	models.BasicCRUD.Update("delete from user_template where template = " + fmt.Sprint(template.Id))
	t.Log("-------------------------------------------TestIntegrationAssignTemplateToUser Passed-------------------------------")
}

/**
 * User unassign templates
 *
 */
func TestIntegrationUnassignTemplateToUser(t *testing.T) {
	_, template := createTestingWorkout(USER_WORKOUT_DEFINITION, t)
	testUser := new(models.UserInfo)
	testUser.Useridentity = USER_INDENTITY
	testUser.Username = string(username)
	UserService.HandleUserLogin(*testUser)
	userTempalte := new(UserTemplate)
	userTempalte.User = testUser.Username
	userTempalte.UserIdentity = testUser.Useridentity
	userTempalte.Templates = []string{template.Name}
	userTempalte.TemplatesId = []int64{template.Id}
	userTempalte.AssignTemplate()
	templateWorkouts := findWorkoutsFromTemplates([]int64{template.Id})
	userWorkouts := findUserWorkouts(userTempalte.UserIdentity)
	assignLength := len(userWorkouts)
	t.Log("find user workouts before unassign", userWorkouts)
	userTempalte.UnassignTemplate()
	userWorkouts = findUserWorkouts(userTempalte.UserIdentity)
	t.Log("find user workouts after unassign", userWorkouts)
	if (assignLength - len(userWorkouts)) != len(templateWorkouts) {
		t.Error("Unable to unassign template to user")
	}

	cleanUp("template", template.Id)
	cleanUp("workout", templateWorkouts)
	for _, workout := range templateWorkouts {
		models.BasicCRUD.Update("delete from working_set where workout = " + fmt.Sprint(workout.Id))
	}
	models.BasicCRUD.Update("delete from template_workout where template = " + fmt.Sprint(template.Id))
	models.BasicCRUD.Update("delete from user_template where template = " + fmt.Sprint(template.Id))
	t.Log("-------------------------------------------TestIntegrationUnassignTemplateToUser Passed-------------------------------")

}

/**
 * User finish one session
 */
