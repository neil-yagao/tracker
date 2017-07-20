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
