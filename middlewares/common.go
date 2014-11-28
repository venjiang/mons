package middlewares

import (
	"github.com/go-martini/martini"
	. "github.com/venjiang/mons/core"
	"github.com/venjiang/mons/models/enum"
	"github.com/venjiang/mons/services"
	"log"
)

func Common() martini.Handler {
	return func(c *Context, svc *services.CoreService, log *log.Logger) {

		title := c.Get("Title")
		if title == nil {
			c.Set("Title", "MONS")
		}

		// 获取站点信息
		// c.Mons["site"] = func(args ...interface{}) Data {
		// 	site, err := svc.GetSite()
		// 	if err != nil {
		// 		panic(err)
		// 	}
		// 	log.Print(site)
		// 	return site
		// }
		// // 获取菜单
		// c.Mons["menu"] = func(args ...interface{}) Data {
		// 	s := 10
		// 	if len(args) > 0 {
		// 		s = args[0].(int)
		// 	}
		// 	if len(args) > 1 {
		// 		log.Print("another params:", args[1])
		// 	}
		// 	m, _ := svc.GetMenus(s)
		// 	return m
		// }
		c.Mons["query_all"] = enum.Query_All
		c.Mons["site"] = svc.GetSite
		c.Mons["menu"] = svc.GetMenus
		c.Mons["article"] = svc.GetArticle
		// c.Mons["article2"] = svc.GetArticle2
		c.Mons["articles"] = svc.GetArticles
		// 获取文章
		// 获取文章元数据

	}
}
