package controllers

import (
	"github.com/astaxie/beego"
)

//SDFSDF
type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "hellw"
	c.Data["Email"] = "@gmail.com"
	c.TplName = "index.tpl"
}
