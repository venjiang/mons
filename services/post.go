package services

import (
	"database/sql"
	"github.com/venjiang/mons/models"
	"github.com/venjiang/mons/models/enum"
)

func getPosts(svc *CoreService, postType enum.PostType, status enum.PostStatus, pageIndex int, pageSize int, qt enum.Query) (*models.PagedPost, error) {
	var posts []*models.Post
	var records int
	var pages int
	del := qt == enum.Query_IsDeleted
	st := int(status)
	sqlSelect := `SELECT * FROM post WHERE type=$1 AND (is_deleted=$4 or $5=-1) and (status & $6=$6 or $6=-1) order by "order",created_time desc limit $3 offset ($2-1)*$3`
	_, err := svc.DbMap.Select(&posts, sqlSelect, int(postType), pageIndex, pageSize, del, int(qt), st)
	sqlCount := `SELECT count(0) FROM post where type=$1 and (is_deleted=$2 or $3=-1) and (status & $4=$4 or $4=-1)`
	i64, _ := svc.DbMap.SelectInt(sqlCount, int(postType), del, int(qt), st)
	records = int(i64)
	pages = calcPages(pageSize, records)

	return &models.PagedPost{Posts: posts, Records: records, Pages: pages}, err
}

func getPost(svc *CoreService, postType enum.PostType, id int, qt enum.Query) (*models.Post, error) {
	post := models.Post{}
	del := qt == enum.Query_IsDeleted
	sqlSelect := `SELECT * FROM post WHERE type=$1 AND id=$2 AND (is_deleted=$3 or $4=-1)`
	err := svc.DbMap.SelectOne(&post, sqlSelect, int(postType), id, del, int(qt))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &post, err
}

func calcPages(pageSize int, records int) int {
	return records/pageSize + 1
}

func (this *CoreService) createPost(post *models.Post) error {
	return this.DbMap.Insert(post)
}
