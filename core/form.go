package core

import ()

// 表单字段
type Field struct {
	Label string // 标签文本
	Value string // 字段值
	Error string // 错误消息
}

type Form struct {
	Fields map[string]Field
}
