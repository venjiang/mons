package models

import (
	"github.com/coopernurse/gorp"
	"github.com/martini-contrib/binding"
	"net/http"
	"regexp"
	"time"
)

type User struct {
	Id            int
	Username      string
	Password      string
	Email         string
	Intro         string `db:"intro"`
	Website       string `db:"website"`
	IsAdmin       bool   `db:"is_admin"`
	Role          int    `db:"role"`   // 1:'subscriber',2:'contributor',4:'author',8:'editor',16'admin'
	Status        int    `db:"status"` // 0:'wait',1:'approved',2:'disabled'
	CreatedTime   time.Time
	CreatedAt     int64
	UpdatedTime   time.Time
	LastLoginTime time.Time `db:"lastlogin_time"`
	authenticated bool      `db:"-" json:"-"`
}

func (this *User) PreInsert(s gorp.SqlExecutor) error {
	this.Role = 1
	this.Status = 0
	this.CreatedTime = time.Now()
	this.CreatedAt = this.CreatedTime.UnixNano()
	this.UpdatedTime = this.CreatedTime
	this.LastLoginTime = this.CreatedTime
	return nil
}
func (this *User) PreUpdate(s gorp.SqlExecutor) error {
	this.UpdatedTime = time.Now()
	return nil
}

type UserRegisterForm struct {
	Email           string `form:"email" binding:"required"`
	Username        string `form:"username" binding:"required"`
	Password        string `form:"password" binding:"required"`
	ConfirmPassword string `form:"confirm_password" `
}

type UserLoginForm struct {
	Account    string `form:"account" binding:"required"`
	Password   string `form:"password" binding:"required"`
	RememberMe bool   `form:"remember_me"`
}

func (form UserRegisterForm) Validate(errs binding.Errors, req *http.Request) binding.Errors {
	if len(form.Username) < 5 {
		errs = append(errs, binding.Error{
			FieldNames:     []string{"username"},
			Classification: "LengthError",
			Message:        "用户名长度不能小于5个字符",
		})
	}

	if len(form.Password) < 5 {
		errs = append(errs, binding.Error{
			FieldNames:     []string{"password"},
			Classification: "LengthError",
			Message:        "密码长度不能小于5个字符",
		})
	}
	if form.Password != form.ConfirmPassword {
		errs = append(errs, binding.Error{
			FieldNames:     []string{"confirm_password"},
			Classification: "LengthError",
			Message:        "确认密码不匹配",
		})
	}

	re := regexp.MustCompile(".+@.+\\..+")
	matched := re.Match([]byte(form.Email))
	if matched == false {
		errs = append(errs, binding.Error{
			FieldNames:     []string{"email"},
			Classification: "FormatError",
			Message:        "请输入有效的电子邮箱地址",
		})
	}
	return errs
}
