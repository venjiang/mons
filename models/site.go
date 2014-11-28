package models

import (
	"github.com/coopernurse/gorp"
	"time"
)

type Site struct {
	Id             int       `db:"id"`
	Name           string    `db:"name"`
	Description    string    `db:"description"`
	Url            string    `db:"url"`
	Theme          string    `db:"theme"`
	Layout         string    `db:"layout"`
	PostSize       int       `db:"post_size"`
	TimezoneOffset int       `db:"timezone_offset"`
	CreatedTime    time.Time `db:"created_time"`
	UpdatedTime    time.Time `db:"updated_time"`

	CanRegister bool `db:"can_register"`
	CanComment  bool `db:"can_comment"`

	AdminEmail string `db:"admin_email"`

	DateFormat          string `db:"date_format"`
	DatetimeFormat      string `db:"datetime_format"`
	DatetimeShortFormat string `db:"datetime_short_format"`

	MailServerUrl      string `db:"mailserver_url"`
	MailServerLogin    string `db:"mailserver_login"`
	MailServerPassword string `db:"mailserver_password"`
	MailServerPort     int    `db:"mailserver_port"`

	ImageSize int `db:"image_size"`
}

func (this *Site) PreInsert(s gorp.SqlExecutor) error {
	this.CreatedTime = time.Now()
	this.UpdatedTime = this.CreatedTime
	return nil
}
func (this *Site) PreUpdate(s gorp.SqlExecutor) error {
	this.UpdatedTime = time.Now()
	return nil
}
