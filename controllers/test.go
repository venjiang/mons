package controllers

import (
	// "errors"
	"database/sql"
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	. "github.com/venjiang/mons/core"
	"github.com/venjiang/mons/lib"
	"github.com/venjiang/mons/models"
	"github.com/venjiang/mons/models/enum"
	"github.com/venjiang/mons/services"
	"log"
	"net/http"
)

// routes
func Home(c *Context, svc *services.CoreService) {

	// holder := make(map[string]interface{})
	//holder,err := db.Select(models.TestView{}, `SELECT id "Age", username "Title", created_time "CreatedTime" FROM "user" limit 1`)
	holder := models.TestView{}
	err := svc.DbMap.SelectOne(&holder, `SELECT id "Age", username "Title", created_time "CreatedTime" FROM "user" limit 1`)
	if err != nil {
		log.Printf("err:%v", err)
	}
	//log.Printf("holder len:%v", len(holder))
	log.Printf("holder:%v", holder)
	//c.Data["Title"] = (holder[0].(*models.TestView)).Title
	c.Data["Title"] = holder.Title + FormatTime(holder.CreatedTime)
	c.Data["msg"] = `测<b>试</b>HTML,<a href="http://www.readfog.com">readfog.com</a>`
	c.Data["sb_title"] = "这里是侧栏"
	// 枚举测试
	log.Print("enum:", enum.RoleType_Subscriber, enum.RoleType_Subscriber|enum.RoleType_Contributor|enum.RoleType_Editor)
	// panic(AppErr{ServerError, "unf测试"})
	user, err := svc.GetUser("venjiang")
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	if user == nil {
		u := models.User{Username: "venjiang", Email: "venjiang@gmail.com", Password: "123456"}
		if err := svc.CreateUser(&u); err != nil {
			panic(err)
		}
		c.Data["User"] = u
	} else {
		c.Data["User"] = user
	}
	// 测试密码哈希
	pwdlist := make([]map[string]string, 5)
	salt := lib.GenerateSalt()
	pwd := "12345"
	hash := lib.EncodePbkdf2(pwd, salt)
	pwdlist = append(pwdlist, map[string]string{pwd + "/" + salt: salt + hash + "/" + fmt.Sprintf("%v-%v", lib.ValidatePassword(pwd, salt+hash), len(salt+hash))})

	salt = lib.GenerateSalt()
	pwd = "1"
	hash = lib.EncodePbkdf2(pwd, salt)
	pwdlist = append(pwdlist, map[string]string{pwd + "/" + salt: salt + hash + "/" + fmt.Sprintf("%v-%v", lib.ValidatePassword(pwd, salt+hash), len(salt+hash))})

	salt = lib.GenerateSalt()
	pwd = "123"
	hash = lib.EncodePbkdf2(pwd, salt)
	pwdlist = append(pwdlist, map[string]string{pwd + "/" + salt: salt + hash + "/" + fmt.Sprintf("%v", lib.ValidatePassword(pwd, salt+hash))})
	salt = lib.GenerateSalt()
	pwd = "大空@$%^"
	hash = lib.EncodePbkdf2(pwd, salt)
	pwdlist = append(pwdlist, map[string]string{pwd + "/" + salt: salt + hash + "/" + fmt.Sprintf("%v", lib.ValidatePassword(pwd, salt+hash))})
	c.Data["pwdlist"] = pwdlist
	// 测试登录用户
	// c.Data["LoginUser"] = lu.(*models.User)
	// c.Data["IsAuthenticated"] = lu.IsAuthenticated()
	// Context 获取
	// c.Data["LoginUser"] = c.User.(*models.User)

	c.HTML(200, "home", "layout")
}
func Test(r render.Render) {
	// user := models.User{Username: "venjiang", Email: "信息化事业部"}
	// data := models.Data{User: user}
	// context := map[string]interface{}{"data": data}
	// r.HTML(200, lib.GetView("test_inc"), context)
	// opt := render.HTMLOptions{Layout: "test"}
	r.HTML(200, "test/test", "test测试页面")
}
func TestOpt(r render.Render, req *http.Request) {
	// opt := render.HTMLOptions{Layout: lib.GetView("layout")}
	// log.Println(opt)
	user := models.User{Username: "venjiang", Email: "信息化事业部"}
	// data := models.Data{Title: "变量值1", User: user}
	// context := map[string]interface{}{"data": data}
	l := "layout"
	pl := req.URL.Query().Get("l")
	if len(pl) > 0 {
		l = pl
	}
	log.Println(l)
	// r.HTML(200, lib.GetView("test"), context, opt)
	iHTML(r, 200, "test", l, user)
}
func iHTML(r render.Render, status int, view string, layout string, data interface{}) {
	opt := render.HTMLOptions{Layout: lib.GetView(layout)}
	context := map[string]interface{}{"data": data, "Theme": "default"}
	r.HTML(status, lib.GetView(view), context, opt)
}
func Test2(r render.Render, params martini.Params) {
	// var opt render.HTMLOptions
	opt := render.HTMLOptions{Layout: lib.GetView("layout2")}
	view := "test"
	if val, ok := params["view"]; ok {
		if len(val) > 0 {
			view = val
		}
	}
	log.Println("view:" + view)
	r.HTML(200, lib.GetView(view), map[string]interface{}{"Theme": "default", "var": "变量值2"}, opt)
}
func Channel(params martini.Params) string {
	return params["channel"]
}
func JsonTest(c *Context) {
	c.Data = map[string]interface{}{"hello": "world", "name": "内容管理系统", "age": 36}
	c.JSON(200)
}
func Ctx(c *Context) {
	c.Data["ctx"] = "context 测试"
	c.Data["a"] = "aaaaa"
	user := models.User{Username: "venjiang", Email: "信息化事业部"}
	// data := models.Data{Title: "变量值1", User: user}
	c.Data["user"] = user
	// c.Data["data"] = data

	l := "layout"
	layout := c.Request.URL.Query().Get("l")
	if len(layout) > 0 {
		l = layout
	}
	if theme, ok := c.Query("t"); ok {
		c.SetTheme(theme)
	}
	tpl := c.Request.URL.Query().Get("tpl")
	if len(tpl) > 0 {
		c.SetTemplate(tpl)
	}
	c.HTML(200, "context", l)
}
func InitData(c *Context, svc *services.CoreService) {
	svc.InitData()
	c.String(200, "init data. Site Administrator: admin/123456")
}
