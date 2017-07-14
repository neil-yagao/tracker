package main

import (
	_ "routers"

	"github.com/astaxie/beego"
	/*"github.com/astaxie/beego/orm"
	_ "models"*/)

func main() {
	//orm.RunCommand()
	beego.Run()
}
