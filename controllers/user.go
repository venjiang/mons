package controllers

import (
	"fmt"
	"github.com/martini-contrib/binding"
	. "github.com/venjiang/mons/core"
	"github.com/venjiang/mons/lib"
	"github.com/venjiang/mons/models"
	"github.com/venjiang/mons/services"
	"log"
	"net/http"
)

func UserRegister(c *Context, errs binding.Errors, form models.UserRegisterForm, svc *services.CoreService) {
	c.Data["Title"] = "用户注册"
	switch c.Request.Method {
	case "POST":
		c.SetErrors(errs)
		// 检查用户名和邮件是否重复
		result := svc.CheckUsername(form.Username)
		log.Printf("check username:%v", result)
		if result {
			c.AddError("username", "用户名已经存在")
		}
		result = svc.CheckEmail(form.Email)
		if result {
			c.AddError("email", "邮箱已被注册了")
		}
		if c.HasError() {
			c.Data["Form"] = form
			c.HTML(200, "/user/register", "layout-simple")
			return
		}
		// 密码哈希
		// password := lib.EncodePassword(form.Password, lib.GenerateSalt())
		// _ = password
		// 创建用户
		err := svc.CreateUser(&models.User{Username: form.Username, Password: form.Password, Email: form.Email, Intro: "测试用户"})
		// err := errors.New("错误测试")
		if err == nil {
			// 登录
			// 返回来源页
			// log.Printf("username:%v", form.Username)
			// c.String(200, form.Username)
			c.Redirect(LoginUrl)
		} else {
			c.Error(500, err)
		}

	default:
		salt := lib.GenerateSalt()
		c.Data["rand"] = salt + fmt.Sprintf("(%d)", len(salt))
		c.HTML(200, "/user/register", "layout-simple")
	}

}

func UserLogin(c *Context, errs binding.Errors, form models.UserLoginForm, svc *services.CoreService, req *http.Request) {
	c.Data["Title"] = "用户登录"
	switch c.Request.Method {
	case "POST":
		c.SetErrors(errs)
		if c.HasError() {
			c.Data["Form"] = form
			c.HTML(200, "/user/login", "layout-simple")
			return
		}
		// 验证用户
		if user, ok := svc.ValidateUser(form.Account, form.Password); ok {
			// err := Authenticate(c.Session, user)
			// // log.Print("AuthenticateSession:", err)
			// if err != nil {
			// 	panic(err)
			// 	return
			// }
			// // 记住我
			// log.Print("form:", form)
			// log.Print("remember_me:", req.FormValue("remember_me"))
			// if form.RememberMe == true {
			// 	c.Session.Options(sessions.Options{MaxAge: 60 * 60 * 24 * 30})
			// }
			// log.Print("url query:", c.Request.URL.Query())
			// if next, ok := c.Query(RedirectParam); ok {
			// 	log.Print("next:", next)
			// 	c.Redirect(next)
			// } else {
			// 	c.Redirect("/")
			// }
			// 重构后的 1
			// err := Authenticate(c.Session, user, form.RememberMe)
			// if err != nil {
			// 	panic(err)
			// 	return
			// }
			// if next, ok := c.Query(RedirectParam); ok {
			// 	log.Print("next:", next)
			// 	c.Redirect(next)
			// } else {
			// 	c.Redirect("/")
			// }
			AuthenticateRedirect(c.Session, user, c.Render, c.Request, form.RememberMe)
			return
		} else {
			c.Data["Form"] = form
			c.AddError("msg", "帐号或密码不正确")
			c.HTML(200, "/user/login", "layout-simple")
			return
		}
	default:
		c.Data["Next"], _ = c.Query(RedirectParam)
		c.HTML(200, "/user/login", "layout-simple")
	}
}
func UserLogout(c *Context, auth Authenticator) {
	Logout(c.Session, auth)
	c.Redirect("/")
}
