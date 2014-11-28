package models

type PostTag struct {
	Id     int `db:"id"`
	PostId int `db:"post_id"`
	TagId  int `db:"tag_id"`
}
