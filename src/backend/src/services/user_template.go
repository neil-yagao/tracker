package services

import "models"

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
	batchInsert := new(models.BatchSqls)
	for _, t := range templatesId {
		batchInsert.Params = append(batchInsert.Params, map[string]interface{}{"user": id, "template": t})
	}

	batchInsert.BatchInsert("user_template")
	//create template related workout for user
	workouts := findWorkoutsFromTemplates(templatesId)
	assignWorkoutsToUser(identity, workouts, startDate)
	//create working set related to workout
	for _, workout := range workouts {
		generateWorkingSets(*workout, 0)
	}
}

func (this *UserTemplate) UnassignTemplate() {
	var id int64 = UserService.GetUserIdFromIdentity(this.UserIdentity)
	deletingParams := make([]map[string]interface{}, 0)
	for _, t := range this.TemplatesId {
		deletingParams = append(deletingParams, map[string]interface{}{"user": id, "template": t})
	}
	models.BasicCRUD.BuildAndUpdate("delete from user_template where user := user and template := template", deletingParams)
	templateWorkouts := findWorkoutsFromTemplates(this.TemplatesId)
	deletingMap := make([]map[string]interface{}, len(templateWorkouts))
	for i, workout := range templateWorkouts {
		deletingMap[i] = map[string]interface{}{"workoutId": workout.Id}
	}
	models.BasicCRUD.BuildAndUpdate("delete from user_workout where workout = :workoutId", deletingMap)
	models.BasicCRUD.BuildAndUpdate("delete from working_set where workout = :workoutId", deletingMap)
}
