package models

import (
	"github.com/coopernurse/gorp"
	"time"
)

type Tag struct {
	Id          int       `db:"id"`
	Name        string    `db:"name"`
	Count       int       `db:"count"`
	CreatedTime time.Time `db:"created_time"`
}

func (this *Tag) PreInsert(s gorp.SqlExecutor) error {
	this.CreatedTime = time.Now()
	return nil
}
