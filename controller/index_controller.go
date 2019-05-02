package controller

import (
	"github.com/aloxc/gappweb/order"
	"github.com/aloxc/gappweb/user"
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Get() {
	content := ""
	u := &user.User{
		Id: 1,
	}
	user.GetUser(u)
	o := &order.Order{
		Id: 1,
	}
	order.GetOrder(o)
	content += u.String()
	content += "\n"
	content += o.String()
	c.Ctx.WriteString(content)
}
