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

func SaveMovementImage(location string, u, m int64) {
	movement := new(models.MovementVideo)
	movement.MappedMovement = new(models.Movement)
	movement.MappedMovement.Id = m
	movement.Location = location
	movement.UploadedBy = new(models.UserInfo)
	movement.UploadedBy.Id = u
	o.InsertOrUpdate(movement)
}

func GetMovementImage(m int64) *models.MovementVideo {
	movement := new(models.MovementVideo)
	movement.MappedMovement = new(models.Movement)
	movement.MappedMovement.Id = m
	o.Read(movement, "MappedMovement")
	//currently not need
	//o.LoadRelated(movement, "UploadedBy")
	return movement
}
