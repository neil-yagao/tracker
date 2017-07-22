package models

import (
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.Debug = true
	orm.RegisterDriver("mysql", orm.DRMySQL)

	config, err := config.NewConfig("ini", "conf/db.conf")
	if err != nil {
		panic(err)
	}
	user := config.String("username")
	password := config.String("password")
	connectString := user + ":" + password + "@tcp(127.0.0.1:3306)/powerlift?charset=utf8"
	orm.RegisterDataBase("default", "mysql", connectString)

	orm.RegisterModel(new(UserInfo), new(Movement), new(AssignedPlan), new(UserSession),
		new(SessionWorkout), new(Exercise), new(SessionMovement), new(Session), new(Plan), new(PhysiqueInfo))
}
