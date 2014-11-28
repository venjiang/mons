package models

import (
	"github.com/coopernurse/gorp"
	"time"
)

type Post struct {
	Id           int       `db:"id"`
	ParentId     int       `db:"pid"`
	UserId       int       `db:"uid"`
	SiteId       int       `db:"site_id"`
	Title        string    `db:"title"`
	Excerpt      string    `db:"excerpt"`      // 摘要
	Content      string    `db:"content"`      // 内容
	Status       int       `db:"status"`       // 0 草稿,1已批准,2发布
	Password     string    `db:"password"`     // 密码
	UrlName      string    `db:"url_name"`     // url名称
	Url          string    `db:"url"`          // url 链接
	Order        int       `db:"order"`        // 排序
	Type         int       `db:"type"`         // 类型
	ContentType  int       `db:"content_type"` // 内容类型
	Layout       string    `db:"layout"`       // 布局,为空使用站点设置布局
	IsDeleted    bool      `db:"is_deleted"`
	CanComment   bool      `db:"can_comment"`   // 是否允许评论
	CommentCount int       `db:"comment_count"` // 评论统计
	Ip           string    `db:"ip"`
	Security     int       `db:"security"`
	CreatedTime  time.Time `db:"created_time"`
	UpdatedTime  time.Time `db:"updated_time"`
}

func (this *Post) PreInsert(s gorp.SqlExecutor) error {
	this.CreatedTime = time.Now()
	this.UpdatedTime = this.CreatedTime
	return nil
}
func (this *Post) PreUpdate(s gorp.SqlExecutor) error {
	this.UpdatedTime = time.Now()
	return nil
}
