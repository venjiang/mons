package services

import (
	"github.com/coopernurse/gorp"
	"github.com/venjiang/mons/models"
	"github.com/venjiang/mons/models/enum"
	"log"
)

type CoreService struct {
	DbMap *gorp.DbMap
}

/***************************************************************/
func InitCoreServie() (*CoreService, error) {
	db, err := models.InitDb()
	return &CoreService{DbMap: db}, err
}

func (this *CoreService) InitData() {
	if this.ExistsSite() == false {
		site := models.Site{
			Name:                "示例站点",
			Description:         "这里是站点的简短描述信息",
			Url:                 "http://www.venjiang.com",
			Theme:               "",
			Layout:              "",
			PostSize:            15,
			TimezoneOffset:      8,
			AdminEmail:          "venjiang@gmail.com",
			DateFormat:          "Y-m-d",
			DatetimeFormat:      "Y-m-d H:i:s",
			DatetimeShortFormat: "Y-m-d H:i",
			MailServerUrl:       "",
			MailServerLogin:     "",
			MailServerPassword:  "",
			MailServerPort:      432,
		}
		err := this.CreateSite(&site)
		panicif(err, "Add default site error!")
	}
	// menu
	if this.ExistsMenu() == false {
		menu := models.Post{
			Title:    "供应链一体化",
			UrlName:  "scm",
			Type:     int(enum.PostType_Menu),
			Status:   int(enum.PostStatus_Published),
			Security: int(enum.PostSecurity_Visibled),
		}
		this.createPost(&menu)
		menu = models.Post{
			Title:    "在线交易",
			UrlName:  "ebiz",
			Type:     int(enum.PostType_Menu),
			Status:   int(enum.PostStatus_Approved | enum.PostStatus_Published),
			Security: int(enum.PostSecurity_Visibled),
		}
		this.createPost(&menu)
		menu = models.Post{
			Title:    "煤炭云资讯",
			UrlName:  "news",
			Type:     int(enum.PostType_Menu),
			Status:   int(enum.PostStatus_Approved | enum.PostStatus_Published),
			Security: int(enum.PostSecurity_Visibled),
		}
		this.createPost(&menu)
		menu = models.Post{
			Title:    "煤炭云分析",
			UrlName:  "bi",
			Type:     int(enum.PostType_Menu),
			Status:   int(enum.PostStatus_Approved | enum.PostStatus_Published),
			Security: int(enum.PostSecurity_Visibled),
		}
		this.createPost(&menu)
	}
	// article
	if this.ExistsArticle() == false {
		post := models.Post{

			Title:    "欢迎使用MONS系统",
			Content:  "MONS 是一个基于 Go 语言的内容管理系统,使用 <b>Martini</b> 框架搭建...",
			Status:   int(enum.PostStatus_Approved | enum.PostStatus_Published | enum.PostStatus_Highlight), // 批准,已发布,高亮
			Type:     int(enum.PostType_Article),
			Security: int(enum.PostSecurity_Visibled),
		}

		poste := this.createPost(&post)
		panicif(poste, "Add article error.")
	}
	// tag := models.Tag{
	// 	Name: "示例",
	// }
	// tagerr := this.CreateTag(&tag)
	// panicif(tagerr, "Add tag error.")

	if this.ExistsAdmin() == false {
		admin := models.User{Username: "admin",
			Email:    "admin@venjiang.com",
			Password: "123456",
			IsAdmin:  true,
		}
		e := this.CreateUser(&admin)
		panicif(e, "Add site administrator error!")
	}

}

func panicif(err error, msg string) {
	if err != nil {
		log.Print(msg, err)
		panic(err)
	}
}

/***************************************************************/
// 测试结构
// func (this *CoreService) TestGetData() (*models.Data, error) {
// 	d := models.Data{}
// 	err := this.DbMap.SelectOne(&d, `select username "Title" from "user" limit 1`)
// 	return &d, err
// }
