package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["controllers/master:OwnerController"] = append(beego.GlobalControllerRouter["controllers/master:OwnerController"],
		beego.ControllerComments{
			"Owners",
			`/owners`,
			[]string{"get"},
			nil})

}
