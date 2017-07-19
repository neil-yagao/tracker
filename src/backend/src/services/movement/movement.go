package movement

import (
	"models"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

var o = orm.NewOrm()
var logger = logs.GetLogger("movement")

func GetAllMovement() []*models.Movement {

	var movements []*models.Movement
	logger.Println("get all movements")
	o.QueryTable("movement").All(&movements)
	return movements
}

func InsertMovement(m *models.Movement) {
	o.InsertOrUpdate(m)
}

func UpdateMovement(m *models.Movement) {
	o.Update(m)
}
