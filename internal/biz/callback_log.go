package biz

//go:generate mockgen -destination  mock/mock_callback_log.go -package biz -source callback_log.go

import (
	"time"
)

// CallbackLog 回调数据库结构体
type CallbackLog struct {
	ID         uint `gorm:"primary_key"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	CallbackId string `json:"callback_id" gorm:"callback_id"`
	IP         string `json:"ip" gorm:"ip"`             // 请求IP地址
	MsgBody    string `json:"msg_body" gorm:"msg_body"` // 消息体
}

func (*CallbackLog) TableName() string {
	return "t_tool_callback_log"
}

type CallbackLogFilter struct {
	CallbackId            string
	Page, PageSize, Count int64
}

type CallbackLogRepo interface {
	// Create 创键回调日志
	Create(*CallbackLog) error
	// QueryList 分页查询回调函数
	QueryList(*CallbackLogFilter) ([]*CallbackLog, error)
}
