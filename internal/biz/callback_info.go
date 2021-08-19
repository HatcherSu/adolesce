package biz

import "github.com/jinzhu/gorm"

// CallbackInfo 回调数据库结构体
type CallbackInfo struct {
	gorm.Model
	CallbackId  string `json:"callback_id" gorm:"callback_id"`
	AppId       string `json:"app_id" gorm:"app_id"`
	VerifyToken string `json:"verify_token" gorm:"verify_token"`
	SecretKey   string `json:"secret_key" gorm:"secret_key"`
}

func (CallbackInfo) TableName() string {
	return "callback_info"
}

type CallbackInfoFilter struct {
	Page, PageSize, Count int64
}

type CallbackInfoRepo interface {
	// Create 创建新的回到信息
	Create(info *CallbackInfo) error
	// QueryList 查询callbackinfo列表
	QueryList(*CallbackInfoFilter) ([]*CallbackInfo, error)
	// QueryByCallbackId 根据callbackId查询回调信息
	QueryByCallbackId(callbackId string) (*CallbackInfo, error)
	// DeleteByID 根据ID删除
	DeleteByID(id int64) error
}
