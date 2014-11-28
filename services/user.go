package services

import (
	"github.com/venjiang/mons/lib"
	"github.com/venjiang/mons/models"
	"log"
	"strings"
)

/***************************************************************/
// 用户
func (this *CoreService) CreateUser(user *models.User) error {
	salt := lib.GenerateSalt()
	hash := lib.EncodePbkdf2(user.Password, salt)
	user.Password = salt + hash
	// if this.ExistsAdmin() {
	// 	user.IsAdmin = false
	// } else {
	// 	user.IsAdmin = true
	// }
	return this.DbMap.Insert(user)
}

func (this *CoreService) GetUserById(id int) (*models.User, error) {
	// user := &models.User{}
	// err := this.DbMap.SelectOne(&user, `select * from "user" where id=$1`, id)
	// return user, err
	u, err := this.DbMap.Get(models.User{}, id)
	if err != nil {
		panic(err)
	}
	return u.(*models.User), err
}
func (this *CoreService) GetUserByUsername(username string) (*models.User, error) {
	user := models.User{}
	err := this.DbMap.SelectOne(&user, `select * from "user" where username=$1`, username)
	if err != nil {
		return nil, err
	}
	return &user, err
}
func (this *CoreService) GetUserByEmail(email string) (*models.User, error) {
	user := models.User{}
	err := this.DbMap.SelectOne(&user, `select * from "user" where email=$1`, email)
	if err != nil {
		return nil, err
	}
	return &user, err
}
func (this *CoreService) GetUser(account string) (*models.User, error) {
	if strings.IndexRune(account, '@') == -1 {
		return this.GetUserByUsername(account)
	}
	return this.GetUserByEmail(account)
}
func (this *CoreService) CheckUsername(username string) bool {
	rows, _ := this.DbMap.SelectInt(`SELECT count(0) FROM "user" where username=$1`, username)
	log.Printf("username rows:", rows)
	return rows > 0
}
func (this *CoreService) CheckEmail(email string) bool {
	rows, _ := this.DbMap.SelectInt(`select count(0) from "user" where email=$1`, email)
	log.Printf("email rows:", rows)
	return rows > 0
}
func (this *CoreService) ExistsAdmin() bool {
	rows, _ := this.DbMap.SelectInt(`SELECT count(0) FROM "user" where username=$1`, "admin")
	return rows > 0
}

// 检查账号是否存在(用户名或邮箱)
func (this *CoreService) HasUser(account string) bool {
	result := false
	if strings.IndexRune(account, '@') == -1 {
		result = this.CheckUsername(account)
	} else {
		result = this.CheckEmail(account)
	}
	return result
}
func (this *CoreService) ValidateUser(account string, password string) (user *models.User, ok bool) {
	// get user
	u, _ := this.GetUser(account)
	if u != nil {
		log.Print(u)
		// get password / validate password
		if ok := lib.ValidatePassword(password, u.Password); ok {
			return u, ok
		}
	}

	return nil, false
}

// func (this *this.DbMap) GetUserList() ([]*User, error) {
// 	users := make([]*User)
// 	_, err := this.DbMap.Select(&users, "SELECT * FROM user")
// 	return threads, err
// }
