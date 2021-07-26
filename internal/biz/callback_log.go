package biz

import (
	"github.com/jinzhu/gorm"
)

// CallbackLog 回调数据库结构体
type CallbackLog struct {
	gorm.Model
	CallbackId string `json:"callback_id" gorm:"callback_id"`
	IP         string `json:"ip" gorm:"ip"`             // 请求IP地址
	MsgBody    string `json:"msg_body" gorm:"msg_body"` // 消息体
}

func (*CallbackLog) TableName() string {
	return "callback_log"
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
