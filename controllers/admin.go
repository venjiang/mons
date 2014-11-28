package controllers

import (
	. "github.com/venjiang/mons/core"
	"github.com/venjiang/mons/services"
)

func AdminHome(c *Context, svc *services.CoreService) {
	initData(c)
	c.HTML(200, "home", "layout")
}
func AdminSetting(c *Context, svc *services.CoreService) {
	initData(c)
	c.HTML(200, "nav", "layout")
}

func initData(c *Context) {
	c.SetTemplate("admin")
	c.Data["Title"] = "Dashboard"
}
