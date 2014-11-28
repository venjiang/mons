mons
====

## Install

- go get github.com/go-martini/martini
- go get github.com/martini-contrib/render
- go get github.com/martini-contrib/binding
- go get github.com/martini-contrib/sessions
- go get github.com/martini-contrib/sessionauth
- go get github.com/coopernurse/gorp
- go get github.com/lib/pq

- go get github.com/Unknwon/goconfig

## Run
- GIN实时编译
	+ go get github.com/codegangsta/gin (%GOPATH%/bin下会自动加入gin)
	+ cd path/mons
	+ gin -a 8080
- GOPM 包管理
	+ go get -u github.com/gpmgo/gopm
	+ 运行 gopm run main.go
	+ 构建 gopm build
	+ 生成包管理文件 gopm gen

## 扩展功能
- 内部函数调用
	+ 站点信息 {{call .mons.site}}
	+ 导航菜单 {{call .mons.menu}}
- 添加自定义函数
	1. 在中间件或控制器,书写: context.Funcs["fname"] = func(m ...interface{}) (interface{},error),内部函数建议用小写名称,返回值最多两个,如果两个,第二个参数为error
	2. 调用: {{call .funcs.fname 参数...}}