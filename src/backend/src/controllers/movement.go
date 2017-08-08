package controllers

import (
	"bytes"
	"encoding/base64"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/logs"
	"image/gif"
	"models"
	"os"
	"services/movement"
)

var BASE_STATIC_DIRECTORY string

func init() {
	config, _ := config.NewConfig("ini", "conf/app.conf")
	BASE_STATIC_DIRECTORY = config.String("static")
}

type MovementController struct {
	GeneralController
}

// @router /movements [get]
func (this *MovementController) GetMovements() {
	this.ServeJson(movement.GetAllMovement())
}

// @router /movement [put]
func (this *MovementController) AddNewMovement() {
	m := new(models.Movement)
	this.ParseRequestBody(m)
	logs.Debug(m)
	movement.InsertMovement(m)
	this.ServeJson()
}

// @router /movement/upload/?:user/?:movement [post]
func (this *MovementController) SaveGif() {
	//defer this.RecoverFromError()
	movement := this.GetString(":movement")
	//GetFile not working hack around
	b := this.Ctx.Request.Form["gif"][0]
	decodeBase64AndSaveToFile(b, movement+".gif")
	this.ServeJson()
}

func decodeBase64AndSaveToFile(base64str, file string) {
	unbased, err := base64.StdEncoding.DecodeString(base64str)
	if err != nil {
		panic(err)
	}

	r := bytes.NewReader(unbased)
	im, err := gif.DecodeAll(r)
	if err != nil {
		panic("Bad gif")
	}
	file = BASE_STATIC_DIRECTORY + file
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		panic("Cannot open file")
	}
	defer f.Close()
	gif.EncodeAll(f, im)
}
