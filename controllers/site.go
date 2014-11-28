package controllers

import (
	// "errors"
	// "database/sql"
	// "fmt"
	// "github.com/go-martini/martini"
	// "github.com/martini-contrib/render"
	"log"
	// "net/http"
	. "github.com/venjiang/mons/core"
	// "github.com/venjiang/mons/lib"
	// "github.com/venjiang/mons/models"
	// "github.com/venjiang/mons/models/enum"
	"encoding/json"
	"github.com/venjiang/mons/services"
)

// routes
func Index(c *Context, svc *services.CoreService, log *log.Logger) {
	menu, err := svc.GetMenus(10)
	if err != nil {
		log.Print("menu ", err)
		return
	}
	site, _ := svc.GetSite()
	result, _ := json.Marshal(site)
	c.Funcs["json"] = func() string {
		return string(result)
	}
	c.Data["Menu"] = menu
	c.Funcs["fn"] = func(m ...interface{}) Data {
		return Raw(m[0].(string) + "><b>Menu</b>")
	}
	c.Funcs["fn1"] = func(a, b, c string) (string, error) {
		return "fn1:" + a + b + c, nil
	}
	c.HTML(200, "index", "layout")
}
func TestFunc() string {
	return "调用函数"
}
