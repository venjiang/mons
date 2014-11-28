package lib

import (
	"path"
)

// 主题
type Theme struct {
	Style    string
	Script   string
	Template string
}

/*** Theme Helper ***/
// 获取主题
func GetTheme() Theme {
	defaultTheme := Theme{Style: "main.css", Script: "", Template: "default"}
	return defaultTheme
}

// 获取视图
func GetView(view string) string {
	return path.Join("default", view)
}

// func (r *renderer) HTML(status int, name string, binding interface{}, htmlOpt ...HTMLOptions) {
// func HTML(r render.Render, status int, view string, data Data, layout string) {
// 	opt := render.HTMLOptions{Layout: GetView(layout)}
// 	context := map[string]interface{}{"data": data}
// 	r.HTML(status, GetView(view), context, opt)
// }
