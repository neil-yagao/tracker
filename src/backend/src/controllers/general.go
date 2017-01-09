package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
)

type GeneralController struct {
	beego.Controller
}

// passing in param should be pointer
// so that value could be persisted after method exit
func (this *GeneralController) ParseRequestBody(parseTo interface{}) {
	beego.Debug("request body:" + string(this.Ctx.Input.RequestBody))
	err := json.Unmarshal(this.Ctx.Input.RequestBody, parseTo)
	if err != nil {
		beego.Error(err)
		this.Data["json"] = map[string]interface{}{"success": false, "reason": "unable to parse request body"}
		this.ServeJSON()
		this.StopRun()
	}
}

func (this *GeneralController) RecoverFromError() {
	if x := recover(); x != nil {
		beego.Error("run time panic: %v", x)
		this.Data["json"] = map[string]interface{}{"success": false, "reason": fmt.Sprintf("run time panic: %v", x)}
		this.ServeJSON()
		this.StopRun()
	}
}

func (this *GeneralController) GetUserIdentity() string {
	user := this.GetString("user")
	if user == "" {
		this.Data["json"] = map[string]interface{}{"success": false, "reason": "unable to identify user info but required"}
		this.ServeJSON()
		this.StopRun()
	}
	beego.Debug("user identity:" + user)
	return user
}

func (this *GeneralController) ServeJson(data ...interface{}) {
	wrappedResult := map[string]interface{}{"success": true}
	if len(data) > 0 {
		wrappedResult["data"] = data[0]
	}
	this.Data["json"] = wrappedResult
	this.ServeJSON()
	this.StopRun()
}
