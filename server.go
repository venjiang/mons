package main

import (
	"github.com/go-martini/martini"
	_ "github.com/lib/pq"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"github.com/venjiang/mons/controllers"
	. "github.com/venjiang/mons/core"
	"github.com/venjiang/mons/middlewares"
	"github.com/venjiang/mons/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

/***************************************************************/
// func init() {
// 	InitConfig()
// 	log.Println(Config.Db.Driver)
// 	log.Println(Config.Db.Source)
// 	log.Println(Config.Web.HttpPort)
// }

// server
func main() {
	m := martini.Classic()
	// sessions
	store := sessions.NewCookieStore([]byte(Config.MustValue("web", "cookie_secure")))
	// sessionauth.RedirectUrl = "/login"
	// sessionauth.RedirectParam = "next"
	LoginUrl = "/login"
	// middlewares
	// session
	store.Options(sessions.Options{MaxAge: 0})
	m.Use(sessions.Sessions("mons_session", store))
	// m.Use(sessionauth.SessionUser(models.NewUser))
	m.Use(middlewares.Authentication(models.NewAuth))
	// render
	m.Use(render.Renderer(render.Options{
		// Directory:  "templates",
		Layout:     "",
		Extensions: []string{".tmpl", ".html"},
		Charset:    "UTF-8",
		IndentJSON: true,
		Funcs: []template.FuncMap{
			{
				"raw":            Raw,
				"formatTime":     FormatTime,
				"formatUnixTime": FormatUnixTime,
				"formatUnixNano": FormatUnixNano,
				"title":          Title,
				"has":            Has,
			},
		},
	}))
	// 中间件-应用程序上下文
	m.Use(CreateContext())
	// 中间件-核心服务
	m.Use(middlewares.CoreService())
	// 中间件-公用服务
	m.Use(middlewares.Common())
	// 日志服务
	// m.Use(middlewares.Logger)
	// 全局错误处理
	m.Use(middlewares.Error())
	// routes
	// test
	m.Get("/", controllers.Index)
	m.Get("/home", controllers.Home)
	m.Get("/t", controllers.Test)
	m.Get("/t1", controllers.TestOpt)
	m.Get("/t2/:view?", controllers.Test2)
	m.Get("/json", controllers.JsonTest)
	m.Any("/ctx", LoginRequired, controllers.Ctx)
	m.Get("/init", controllers.InitData)
	// auth
	m.Any("/register", binding.Form(models.UserRegisterForm{}), controllers.UserRegister)
	m.Any("/login", binding.Form(models.UserLoginForm{}), controllers.UserLogin)
	m.Any("/logout", controllers.UserLogout)
	// user
	// m.Group("/user", func(r martini.Router) {
	// 	r.Any("/register", binding.Form(models.UserRegisterForm{}), controllers.UserRegister)
	// 	r.Any("/login", controllers.UserRegister)
	// })
	// admin
	m.Group("/admin", func(r martini.Router) {
		r.Get("/", controllers.AdminHome)
		r.Get("/setting", controllers.AdminSetting)
	}, LoginRequired)
	// channel
	m.Get("/:channel", controllers.Channel)

	port := strconv.Itoa(Config.MustInt("web", "http_port"))

	log.Printf("[%s]server is starting...", port)
	// port := strconv.Itoa(Config.Web.HttpPort)
	http.ListenAndServe(":"+port, m)
}
