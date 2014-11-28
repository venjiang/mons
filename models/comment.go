package models

import (
	"github.com/coopernurse/gorp"
	"time"
)

type Comment struct {
	Id          int       `db:"id"`
	PostId      int       `db:"post_id"`
	ParentId    int       `db:"pid"`
	UserId      int       `db:"uid"`
	Username    string    `db:"username"`
	Website     string    `db:"website"`
	Title       string    `db:"title"`
	Content     string    `db:"content"`
	Ip          string    `db:"ip"`
	IsDeleted   bool      `db:"is_deleted"`
	CreatedTime time.Time `db:"created_time"`
}

func (this *Comment) PreInsert(s gorp.SqlExecutor) error {
	this.CreatedTime = time.Now()
	return nil
}
