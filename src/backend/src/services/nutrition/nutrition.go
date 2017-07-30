package nutrition

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"models"
)

var o = orm.NewOrm()
var logger = logs.GetLogger("movement")

func SaveUserPhysique(userId int64, info *models.PhysiqueInfo) {
	user := new(models.UserInfo)
	user.Id = userId
	info.User = user
	o.Insert(info)
}

func GetUserPhysique(userId int64) *models.PhysiqueInfo {
	var physique = new(models.PhysiqueInfo)
	err := o.QueryTable("physique_info").Filter("user_id", userId).OrderBy("-record_time", "-id").One(physique)
	if err != nil {
		panic(err)
	}
	logs.Debug(physique)
	return physique
}
