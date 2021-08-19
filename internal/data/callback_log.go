package data

import (
	"adolesce/internal/biz"
	"adolesce/pkg/log"
)

// 实现接口
var _ biz.CallbackLogRepo = (*callbackLogRepo)(nil)

func NewCallbackLogRepo(data *Data, log log.Logger) biz.CallbackLogRepo {
	return &callbackLogRepo{
		data,
		log,
	}
}

type callbackLogRepo struct {
	data *Data
	log  log.Logger
}

func (c *callbackLogRepo) Create(callbackLog *biz.CallbackLog) error {
	return c.data.db.Create(callbackLog).Error
}

func (c *callbackLogRepo) QueryList(filter *biz.CallbackLogFilter) ([]*biz.CallbackLog, error) {
	var infos []*biz.CallbackLog
	sqlDB := c.data.db.Model(biz.CallbackLog{})
	if filter.CallbackId != "" {
		sqlDB = sqlDB.Where("callback_id = ?", filter.CallbackId)
	}
	// count 查询必须在分页之前，不然会报错
	if err := sqlDB.Count(&filter.Count).Error; err != nil {
		return nil, err
	}
	if filter.Page > 0 && filter.PageSize > 0 {
		offset := (filter.Page - 1) * filter.PageSize
		sqlDB = sqlDB.Limit(filter.PageSize).Offset(offset)
	}
	if err := sqlDB.Count(&filter.Count).Find(&infos).Error; err != nil {
		return nil, err
	}
	return infos, nil
}
