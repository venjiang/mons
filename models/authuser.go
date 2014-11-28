package models

import (
	. "github.com/venjiang/mons/core"
	"time"
)

/***************************************************************/
// sessionuser interface implementation
func (u *User) IsAuthenticated() bool {
	return u.authenticated
}

func (u *User) Login() {
	// 更新其它信息,如最后登录时间等
	u.LastLoginTime = time.Now()
	dbmap := OpenDb()
	dbmap.Update(u)
	u.authenticated = true
}

func (u *User) Logout() {
	u.authenticated = false
}

func (u *User) UniqueId() interface{} {
	return u.Id
}

func (u *User) GetById(id interface{}) error {
	dbmap := OpenDb()
	err := dbmap.SelectOne(u, `select * from "user" where id=$1`, id)
	return err
}

// session user martini middleware
func NewAuth() Authenticator {
	return &User{Username: "游客", IsAdmin: false}
}

/***************************************************************/
