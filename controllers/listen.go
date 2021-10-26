package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type ListenController struct {
	beego.Controller
}

func (c *ListenController) Get() {
	c.TplName = "listen.html"
}
