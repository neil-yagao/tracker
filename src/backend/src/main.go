package main

import (
	"github.com/astaxie/beego"
	_ "routers"
)

func main() {

	beego.Run()
}

/*import (
	"github.com/astaxie/beego/orm"
	_ "models"
)

func main() {

	orm.RunCommand()
}*/
