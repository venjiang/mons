package middlewares

import (
	. "github.com/venjiang/mons/core"
	"log"
	"net/http"
)

func Logger(c *Context, log *log.Logger, req *http.Request) {
	log.Println("-----------------------------")
	log.Println("current template:" + c.GetTemplate()) // 始终默认,因为没有设置值
	log.Println("req.URL.Query:")
	// meta.Keywords = "页面关键字"
	// log.Println(meta.Keywords)
	query := req.URL.Query()
	for key := range query {
		log.Println(key + "->" + query.Get(key))
	}

	c.Parent.Next()
	log.Println("-----------------------------")
}
