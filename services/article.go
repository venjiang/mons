package services

import (
	"errors"
	"github.com/venjiang/mons/models"
	"github.com/venjiang/mons/models/enum"
)

// 获取文章
func (this *CoreService) GetArticle(id int, qt ...interface{}) (*models.Post, error) {
	if len(qt) > 0 {
		switch qt[0].(type) {
		case int:
			return getPost(this, enum.PostType_Article, id, enum.Query(qt[0].(int)))
		case enum.Query:
			return getPost(this, enum.PostType_Article, id, qt[0].(enum.Query))
		default:
			return nil, errors.New("unknown argment type")
		}
	}
	return getPost(this, enum.PostType_Article, id, enum.Query_All)
}

// 测试自定义类型参数
// func (this *CoreService) GetArticle(id int, qt int) (*models.Post, error) {
// 	return getPost(this, enum.PostType_Article, id, enum.Query(qt))
// }

// func (this *CoreService) GetArticle(id int, qt enum.Query) (*models.Post, error) {
// 	return getPost(this, enum.PostType_Article, id, qt)
// }
func (this *CoreService) GetArticles(status enum.PostStatus, pageIndex int, pageSize int, qt enum.Query) (*models.PagedPost, error) {
	return getPosts(this, enum.PostType_Article, status, pageIndex, pageSize, qt)
}

// 检查文章是否存在
func (this *CoreService) ExistsArticle() bool {
	rows, _ := this.DbMap.SelectInt(`SELECT count(0) FROM post where type=$1`, int(enum.PostType_Article))
	return rows > 0
}
