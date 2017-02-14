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
	for _, t := range templatesId {
		var criteria map[string]interface{} = map[string]interface{}{"user": id, "template": t}
		models.BasicCRUD.IgnoreSave("user_template", criteria)
	}
	//create template related workout for user
	workouts := findWorkoutsFromTemplates(templatesId)
	assignWorkoutsToUser(identity, workouts, startDate)
}

func (this *UserTemplate) UnassignTemplate() {
	var id int64 = UserService.GetUserIdFromIdentity(this.UserIdentity)
	for _, t := range this.TemplatesId {
		deleteSql := models.QueryBuilder.BuildQuery("delete from user_template where user = :user and template = :template",
			map[string]interface{}{"user": id, "template": t})
		models.BasicCRUD.Update(deleteSql)
		deleteSql = models.QueryBuilder.BuildQuery("delete from template_workout where template = :template", map[string]interface{}{"template": t})
		models.BasicCRUD.Update(deleteSql)
	}
}
