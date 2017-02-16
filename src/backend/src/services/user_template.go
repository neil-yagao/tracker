package services

import (
	"encoding/json"
	"fmt"
	"models"
)

const QUERY_USER_TEMPLATE string = "select u.username, u.useridentity, t.name, t.id from user u, user_template ut," +
	"template t where u.id = ut.user and ut.template = t.id and u.useridentity = :useridentity"

type UserTemplate struct {
	User         string   `json:"user"`
	UserIdentity string   `json:"userid"`
	Templates    []string `json:"templates"`
	TemplatesId  []int64  `json:"templatesId"`
	StartDate    string   `json:"startDate"`
}

func (this *UserTemplate) GetUserTemplate() *UserTemplate {
	return getUserTemplate(this.UserIdentity)
}

func getUserTemplate(identity string) *UserTemplate {
	userTemplate := new(UserTemplate)
	rows := models.BasicCRUD.BuildAndQuery(QUERY_USER_TEMPLATE, map[string]interface{}{"useridentity": identity})
	userTemplate.Templates = make([]string, 0)
	userTemplate.TemplatesId = make([]int64, 0)
	var username, useridentity string
	for rows.Next() {
		var template string
		var templateid int64
		rows.Scan(&username, &useridentity, &template, &templateid)
		userTemplate.User = username
		userTemplate.UserIdentity = useridentity
		userTemplate.Templates = append(userTemplate.Templates, template)
		userTemplate.TemplatesId = append(userTemplate.TemplatesId, templateid)
	}
	return userTemplate
}

func (this *UserTemplate) AssignTemplate() {
	assignTempalteToUser(this.TemplatesId, this.UserIdentity, this.StartDate)
}

func assignTempalteToUser(templatesId []int64, identity string, startDate string) {
	var id int64 = UserService.GetUserIdFromIdentity(identity)
	for _, t := range templatesId {
		var criteria map[string]interface{} = map[string]interface{}{"user": id, "template": t}
		models.BasicCRUD.IgnoreSave("user_template", criteria)
	}
	//create template related workout for user
	workouts := findWorkoutsFromTemplates(templatesId)
	assignWorkoutsToUser(identity, workouts, startDate)
	//create working set related to workout
	for _, t := range templatesId {
		template := findTemplateById(t)
		var workoutDefinition WorkoutDefinition
		json.Unmarshal([]byte(template.Content), &workoutDefinition)
		for _, workout := range workouts {
			generateWorkingSets(*workout, &workoutDefinition, 0)
		}
	}
}

func (this *UserTemplate) UnassignTemplate() {
	var id int64 = UserService.GetUserIdFromIdentity(this.UserIdentity)
	for _, t := range this.TemplatesId {
		deleteSql := models.QueryBuilder.BuildQuery("delete from user_template where user = :user and template = :template",
			map[string]interface{}{"user": id, "template": t})
		models.BasicCRUD.Update(deleteSql)
	}
	templateWorkouts := findWorkoutsFromTemplates(this.TemplatesId)
	deletingMap := make([]map[string]interface{}, len(templateWorkouts))
	for i, workout := range templateWorkouts {
		deletingMap[i] = map[string]interface{}{"workoutId": workout.Id}
	}
	models.BasicCRUD.BuildAndUpdate("delete from user_workout where workout = :workoutId", deletingMap)
	models.BasicCRUD.BuildAndUpdate("delete from working_set where workout = :workoutId", deletingMap)
}

func generateWorkingSets(workout Workout, template *WorkoutDefinition, extraWeight float64) {
	var containMovements []models.MovementTemplate = template.Movements
	for _, mv := range containMovements {
		movement := new(models.Movement)
		movement.Description = mv.Name
		movement.TargetMuscle = template.TargetMuscle
		movement.SecondaryMuscle = ""
		movement.Name = mv.Name
		movementId := getMovementId(*movement)
		targetWeight := mv.Weight + float64(extraWeight)
		if mv.NeedWarmSet > 0 {
			generateWarmingSet(movementId, workout.Id, mv.Repeats, targetWeight)
		}
		//generate working set for movement
		for sequence := 1; sequence <= mv.Sets; sequence++ {
			assignWorkingSet(movementId, workout.Id, mv.Repeats, targetWeight, sequence)
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
	for workingWeight := startWeight; workingWeight <= targetWeight*0.9; workingWeight += realAddition {
		assignWorkingSet(movementId, workout, targetNumber, transferToUsableWeight(workingWeight), 0)
	}
}
