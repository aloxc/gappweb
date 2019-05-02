package main

import (
	"github.com/aloxc/gappweb/config"
	_ "github.com/aloxc/gappweb/router"
	"github.com/astaxie/beego"
	"os"
)

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	port := os.Getenv(config.WEB_SERVER_PORT)
	if port == "" {
		beego.Run()
	} else {
		beego.Run(port)
	}
}
