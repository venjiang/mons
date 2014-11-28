package middlewares

import (
	// "fmt"
	// "github.com/codegangsta/inject"
	"github.com/go-martini/martini"
	. "github.com/venjiang/mons/core"
	"log"
	"strconv"
)

const (
	ErrorHtml = `<!DOCTYPE html>
<html>
<head lang="zh-CN">
    <meta charset="UTF-8">
    <title>应用程序错误 - MONS</title>
    <link rel="stylesheet" type="text/css" href="/default/css/main.css">
</head>
<body>
    <div class="page">
        <div class="nav">
            <div class="logo">
                <a href="/"><img src="/default/img/logo.png"></a>
            </div>
            <ul>
                <li><a href="/ctx">应用程序上下文</a></li>
                <li><a href="/t">无布局-模板页test</a></li>
                <li><a href="/t1">布局一-模板页test(opt)</a></li>
                <li><a href="/t2/test1">布局二-模板页test1</a></li>
                <li><a href="/user/register">注册</a></li>
                <li><a href="/t2/test3">布局二-模板页test3</a></li>
                <li><a href="/demo.html">静态页</a></li>
                <li><a href="/json">json页</a></li>
                <li><a href="/hello">频道页</a></li>
            </ul>
            <div class="clear"></div>
        </div>
        <div id="container">
            <h2>应用程序错误</h2>
            %s
        </div>
        <div id="footer">
            <hr/>
            版权所有 &copy; 2014 MONS
        </div>
    </div>
</body>
</html>`
)

func Error() martini.Handler {
	return func(this *Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("[MONS] error: %v", err)
				if Config.MustBool("web", "debug") == false {
					switch e := err.(type) {
					case AppErr:
						switch e.Type {
						case UserNotFound:
							this.String(404, "用户没找到!")
						default:
							this.Data[AppError] = e.Message
							this.HTML(500, "/error/error-"+strconv.Itoa(500))

						}

					default:
						log.Print(e)
						this.Data[AppError] = err
						this.HTML(500, "/error/error-500")
					}
				} else {
					panic(err)
				}

				// Lookup the current responsewriter
				// val := c.Get(inject.InterfaceOf((*http.ResponseWriter)(nil)))
				// res := val.Interface().(http.ResponseWriter)

				// // respond with panic message while in development mode
				// var body []byte
				// res.Header().Set("Content-Type", "text/html;charset=UTF-8")
				// body = []byte(fmt.Sprintf(ErrorHtml, err))

				// res.WriteHeader(http.StatusInternalServerError)
				// if nil != body {
				// 	res.Write(body)
				// }
			}
		}()

		// c.Next()
		this.Parent.Next()
	}
}
