package core

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	// "github.com/martini-contrib/sessionauth"
	"github.com/martini-contrib/sessions"
	"log"
	"net/http"
	// "net/url"
	"path"
	"strconv"
)

const (
	Theme           = "Theme"           // 主题
	Template        = "Template"        // 模板
	ContentText     = "text/plain"      // 文本内容类型
	HasError        = "HasError"        // 表单验证是否有错误
	Errors          = "Errors"          // 表单错误
	AppError        = "AppError"        // 应用错误
	IsAuthenticated = "IsAuthenticated" // 是否验证通过
	Login           = "Login"           // 当前登录用户
)

// Web Context
type Context struct {
	render.Render
	Parent  martini.Context
	Session sessions.Session
	Request *http.Request
	Writer  http.ResponseWriter
	// Cookie    http.Cookie
	Data   map[string]interface{}
	errors map[string]string
	// Login user
	User Authenticator
	//
	Mons map[string]interface{}
	// Funcs is the type of the map defining the mapping from names to functions.
	// Each function must have either a single return value, or two return values of
	// which the second has type error. In that case, if the second (error)
	// return value evaluates to non-nil during execution, execution terminates and
	// Execute returns that error.
	Funcs map[string]interface{}
}

/***************************************************************/
func CreateContext() martini.Handler {
	return func(c martini.Context, r render.Render, session sessions.Session, currentUser Authenticator, req *http.Request, w http.ResponseWriter, l *log.Logger) {
		ctx := &Context{
			Render:  r,
			Session: session,
			Request: req,
			Writer:  w,
			Parent:  c,
			Data:    make(map[string]interface{}),
			errors:  make(map[string]string),
			User:    currentUser,
			Mons:    make(map[string]interface{}),
			Funcs:   make(map[string]interface{}),
		}
		// ctx.setErrors(errs)
		ctx.Data[Theme] = "default"
		ctx.Data[Template] = "default"
		ctx.Data[IsAuthenticated] = currentUser.IsAuthenticated()
		ctx.Data["mons"] = ctx.Mons
		ctx.Data["funcs"] = ctx.Funcs
		// l.Print("martinit-l")
		// log.Print("sys-log")
		ctx.Data[Login] = currentUser

		c.Map(ctx)
	}
}

// 检查表单验证是否存在错误
func (this *Context) HasError() bool {
	return len(this.errors) > 0
}
func (this *Context) SetErrors(errs binding.Errors) {
	if len(errs) > 0 {
		for _, err := range errs {
			this.errors[err.FieldNames[0]] = err.Message
		}

		this.Data[HasError] = this.HasError()
		this.Data[Errors] = this.errors
	}
}
func (this *Context) AddError(fieldName, message string) {
	this.errors[fieldName] = message
	this.Data[HasError] = this.HasError()
	this.Data[Errors] = this.errors
}

/***************************************************************/
func (this *Context) HTML(status int, view string, layout ...string) {
	var opt render.HTMLOptions
	if len(layout) > 0 {
		opt = render.HTMLOptions{Layout: path.Join(this.GetTemplate(), layout[0])}
	}
	this.Render.HTML(status, path.Join(this.GetTemplate(), view), this.Data, opt)
}
func (this *Context) JSON(status int) {
	this.Render.JSON(status, this.Data)
}

func (this *Context) String(status int, content string) {
	this.Writer.Header().Set(render.ContentType, ContentText+"; charset=UTF-8")
	this.Writer.WriteHeader(status)
	this.Writer.Write([]byte(content))
}

func (this *Context) Error(status int, err error) {
	log.Printf("[MONS] Application Error:%v", err)
	this.Data[AppError] = err.Error()
	this.HTML(status, "/error/error-"+strconv.Itoa(status))
}

func (this *Context) File(filepath string) {
	http.ServeFile(this.Writer, this.Request, filepath)
}

/***************************************************************/
// Theme
func (this *Context) GetTheme() string {
	return this.Get(Theme).(string)
}
func (this *Context) SetTheme(theme string) {
	this.Set(Theme, theme)
}

// Template
func (this *Context) GetTemplate() string {
	return this.Get(Template).(string)
}
func (this *Context) SetTemplate(tpl string) {
	this.Set(Template, tpl)
}

// Common
func (this *Context) Get(key string) interface{} {
	return this.Data[key]
}

func (this *Context) Set(key string, val interface{}) {
	this.Data[key] = val
}

func (this *Context) Delete(key string) {
	delete(this.Data, key)
}

func (this *Context) Clear() {
	for key := range this.Data {
		this.Delete(key)
	}
}

/***************************************************************/
// Http Query
// func (this *Context) Query() url.Values {
// 	return this.Request.URL.Query()
// }
func (this *Context) Query(key string) (val string, ok bool) {
	v := this.Request.URL.Query().Get(key)
	if len(v) > 0 {
		return v, true
	}
	return "", false
}

/***************************************************************/
func (this *Context) IsAuthenticated() bool {
	return this.User.IsAuthenticated()
}

// Authenticate User
// type AuthenticateFun func(user sessionauth.User) bool

// func (this *Context) Authenticate(user sessionauth.User, f AuthenticateFun) {
// 	if f(user) {
// 		this.Data["IsAuhenticated"] = true
// 		this.Data["CurrentUser"] = user
// 		return
// 	}
// 	this.Data["IsAuhenticated"] = false
// }

/***************************************************************/
