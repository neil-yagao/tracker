package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.Debug = true
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "root:password@tcp(127.0.0.1:3306)/powerlift?charset=utf8")

	orm.RegisterModel(new(UserInfo), new(Movement), new(AssignedPlan), new(UserSession),
		new(SessionWorkout), new(Exercise), new(SessionMovement), new(Session), new(Plan), new(PhysiqueInfo))
}
