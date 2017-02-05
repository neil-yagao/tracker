package services

import "models"

const QUERY_USER_TEMPLATE string = "select u.username, u.useridentity, t.name, t.id from user u, user_template ut," +
	"template t where u.id = ut.user and ut.template = t.id and u.useridentity = :useridentity"

type UserTemplate struct {
	User         string   `json:"user"`
	UserIdentity string   `json:"userid"`
	Templates    []string `json:"templates"`
	TemplatesId  []int64  `json:"templatesId"`
}

func (this *UserTemplate) GetUserTemplate() *UserTemplate {
	rows := models.BasicCRUD.BuildAndQuery(QUERY_USER_TEMPLATE, map[string]interface{}{"useridentity": this.UserIdentity})
	this.Templates = make([]string, 0)
	this.TemplatesId = make([]int64, 0)
	var username, useridentity string
	for rows.Next() {
		var template string
		var templateid int64
		rows.Scan(&username, &useridentity, &template, &templateid)
		this.User = username
		this.UserIdentity = useridentity
		this.Templates = append(this.Templates, template)
		this.TemplatesId = append(this.TemplatesId, templateid)
	}
	return this
}

func (this *UserTemplate) AssignTemplate() {
	var id int64 = UserService.GetUserIdFromIdentity(this.UserIdentity)
	for _, t := range this.TemplatesId {
		var criteria map[string]interface{} = map[string]interface{}{"user": id, "template": t}
		rows := models.BasicCRUD.BuildAndQuery("select * from user_template where  user = :user and template = :template", criteria)
		if !rows.Next() {
			models.BasicCRUD.Save("user_template", criteria)
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
}
