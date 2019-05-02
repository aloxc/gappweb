package main

import (
	_ "github.com/aloxc/gappweb/router"
	"github.com/astaxie/beego"
)

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}
