package services

import (
	"github.com/venjiang/mons/models"
	"github.com/venjiang/mons/models/enum"
)

func (this *CoreService) GetMenus(size ...int) (*models.PagedPost, error) {
	s := 10
	if len(size) > 0 {
		s = size[0]
	}
	return getPosts(this, enum.PostType_Menu, enum.PostStatus_Published|enum.PostStatus_Approved, 1, s, enum.Query_All)

}
func (this *CoreService) TestMenus(size ...int) (menus []*models.Post) {
	s := 10
	if len(size) > 0 {
		s = size[0]
	}
	r, _ := getPosts(this, enum.PostType_Menu, enum.PostStatus_Published|enum.PostStatus_Approved, 1, s, enum.Query_All)
	menus = r.Posts
	return
}

func (this *CoreService) ExistsMenu() bool {
	rows, _ := this.DbMap.SelectInt(`SELECT count(0) FROM post where type=$1`, int(enum.PostType_Menu))
	return rows > 0
}
