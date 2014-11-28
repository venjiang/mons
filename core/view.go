package core

import (
	"html/template"
	"time"
)

/*** View Helper ***/
// 显示原始字符串
func Raw(args ...interface{}) template.HTML {
	str := args[0].(string)
	return template.HTML(str)
}

// 格式化时间
func FormatUnixTime(args ...interface{}) string {
	t1 := time.Unix(args[0].(int64), 0)
	return t1.Format(time.Stamp)

}
func FormatUnixNano(unixnano int64) string {
	return time.Unix(0, unixnano).Format("2006-01-02 15:04:05")
}
func FormatTime(args ...interface{}) string {
	return args[0].(time.Time).Format("2006-01-02 15:04:05")
}

// 标题测试
func Title(args ...interface{}) string {
	if len(args) > 0 && len(args[0].(string)) > 0 {
		return args[0].(string) + " - " + "CMS"
	}

	return "CMS默认标题"
}

// 判断是否存变量,如果存在,返回变量值
func Has(v interface{}) bool {
	if v != nil {
		return true
	}
	return false
}

// func MdToHtml(args ...interface{}) template.HTML {
// 	return template.HTML(string(blackfriday.MarkdownBasic([]byte(args[0].(string)))))
// }
func Equal(args ...interface{}) bool {
	return args[0] == args[1]
}
func Translate(lang string, format string) string {
	return Translate(lang, format)
}
func Translatef(lang string, format string, args ...interface{}) string {
	return Translatef(lang, format, args...)
}

// "privilege": func(user interface{}, module int) bool {
// 	if user == nil {
// 		return false
// 	}
// 	return CheckPermission(user, module)
// },
// "plus": func(args ...int) int {
// 	var result int
// 	for _, val := range args {
// 		result += val
// 	}
// 	return result
// },
