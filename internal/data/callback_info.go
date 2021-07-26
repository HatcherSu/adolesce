package data

import (
	"cloud_callback/internal/biz"
	"cloud_callback/internal/pkg/log"
)

// 实现接口
var _ biz.CallbackInfoRepo = (*callbackInfoRepo)(nil)

func NewCallbackInfoRepo(data *Data, log log.Logger) biz.CallbackInfoRepo {
	return &callbackInfoRepo{
		data,
		log,
	}
}

type callbackInfoRepo struct {
	data *Data
	log  log.Logger
}

func (c *callbackInfoRepo) DeleteByID(id int64) error {
	return c.data.GetDB().Delete(biz.CallbackInfo{}, id).Error
}

func (c *callbackInfoRepo) QueryByCallbackId(callbackId string) (*biz.CallbackInfo, error) {
	var info biz.CallbackInfo
	if err := c.data.GetDB().Where("callback_id = ?", callbackId).First(&info).Error; err != nil {
		return nil, err
	}
	return &info, nil
}

func (c *callbackInfoRepo) Create(info *biz.CallbackInfo) error {
	return c.data.GetDB().Create(info).Error
}

// QueryList 分页查询info
func (c callbackInfoRepo) QueryList(filter *biz.CallbackInfoFilter) ([]*biz.CallbackInfo, error) {
	var infos []*biz.CallbackInfo
	sqlDB := c.data.GetDB().Model(biz.CallbackInfo{})
	// count 查询必须在分页之前，不然会报错
	if err := sqlDB.Count(&filter.Count).Error; err != nil {
		return nil, err
	}
	if filter.Page > 0 && filter.PageSize > 0 {
		offset := (filter.Page - 1) * filter.PageSize
		sqlDB = sqlDB.Limit(filter.PageSize).Offset(offset)
	}
	if err := sqlDB.Find(&infos).Error; err != nil {
		return nil, err
	}
	return infos, nil
}
