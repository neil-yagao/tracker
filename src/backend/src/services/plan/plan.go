package plan

import (
	"models"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

var o = orm.NewOrm()
var logger = logs.GetLogger()

func CreateNewPlan(instance *models.Plan) {
	o.Insert(instance)
	logger.Println("inserted plan:", instance)
	sessions := instance.Sessions
	for _, session := range sessions {
		o.Insert(session)
		logger.Println("inserted session:", session)
		movements := session.Workouts
		for _, movement := range movements {
			o.Insert(movement)
			logger.Println("inserted movement:", movement)

		}
	}
}
