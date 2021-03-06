package service

import (
	api "adolesce/api/callback"
	"adolesce/internal/biz"
	"adolesce/internal/pkg/hash"
	"adolesce/pkg/log"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
)

// 实现接口
var _ api.CallbackHTTPServer = (*callbackService)(nil)

func NewCallbackService(log log.Logger, uc biz.ExampleUsecase) api.CallbackHTTPServer {
	return &callbackService{
		log: log,
		uc:  uc,
	}
}

type callbackService struct {
	uc  biz.ExampleUsecase
	log log.Logger
}

func (s *callbackService) DeleteCallbackInfo(context *gin.Context, req *api.DeleteCallbackInfoReq) error {
	if err := s.uc.DeleteInfo(req.ID); err != nil {
		s.log.Error("DeleteCallbackInfo-->DeleteInfo", zap.Int64("id", req.ID), zap.Error(err))
		return err
	}
	return nil
}

func (s *callbackService) QueryCallbackLogList(context *gin.Context, req *api.QueryCallbackLogListReq) (*api.CallbackLogListTable, error) {
	filter := &biz.CallbackLogFilter{
		CallbackId: req.CallbackId,
		Page:       req.Page,
		PageSize:   req.Limit,
	}
	logs, err := s.uc.QueryLogList(filter)
	if err != nil {
		s.log.Error("QueryCallbackLogList-->QueryLogList", zap.Error(err))
		return nil, err
	}
	var data []*api.CallbackLog
	if err := copier.Copy(&data, logs); err != nil {
		s.log.Error("QueryCallbackLogList-->Copy", zap.Error(err))
		return nil, err
	}
	return &api.CallbackLogListTable{
		Code:  0,
		Count: filter.Count,
		Data:  data,
	}, nil
}

func (s *callbackService) QueryCallbackInfoList(_ *gin.Context, req *api.QueryCallbackInfoListReq) (*api.CallbackInfoListTable, error) {
	filter := &biz.CallbackInfoFilter{
		Page:     req.Page,
		PageSize: req.Limit,
	}
	infos, err := s.uc.QueryInfoList(filter)
	if err != nil {
		s.log.Error("QueryCallbackInfoList-->QueryInfoList", zap.Error(err))
		return nil, err
	}
	var data []*api.CallbackInfo
	if err := copier.Copy(&data, infos); err != nil {
		s.log.Error("QueryCallbackInfoList-->Copy", zap.Error(err))
		return nil, err
	}
	return &api.CallbackInfoListTable{
		Code:  0,
		Count: filter.Count,
		Data:  data,
	}, nil
}

func (s *callbackService) CreateCallbackID(_ *gin.Context, req *api.CreateCallbackIDReq) (*api.CallbackIdResp, error) {
	// 生成ID
	callbackId := hash.MD5Hash(fmt.Sprintf("%s-%s-%s", req.AppId, req.SecretKey, req.VerifyToken))
	if err := s.uc.CreateInfo(&biz.CallbackInfo{
		CallbackId:  callbackId,
		AppId:       req.AppId,
		VerifyToken: req.VerifyToken,
		SecretKey:   req.SecretKey,
	}); err != nil {
		s.log.Error("CreateCallbackID-->CreateInfo", zap.Error(err))
		return nil, err
	}
	return &api.CallbackIdResp{
		CallbackId: callbackId,
	}, nil
}
