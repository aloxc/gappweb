package router

import (
	"github.com/aloxc/gappweb/controller"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controller.IndexController{})

}
