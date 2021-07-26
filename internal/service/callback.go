package service

import (
	api "cloud_callback/api/callback"
	"cloud_callback/internal/biz"
	"cloud_callback/internal/pkg/code"
	"cloud_callback/internal/pkg/hash"
	"cloud_callback/internal/pkg/log"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"io/ioutil"
)

// 实现接口
var _ api.CallbackHTTPServer = (*CallbackService)(nil)

func NewCallbackService(log log.Logger, uc biz.CallbackUsecase) api.CallbackHTTPServer {
	return &CallbackService{
		log: log,
		uc:  uc,
	}
}

type CallbackService struct {
	uc  biz.CallbackUsecase
	log log.Logger
}

func (s *CallbackService) DeleteCallbackInfo(context *gin.Context, req *api.DeleteCallbackInfoReq) error {
	if err := s.uc.DeleteInfo(req.ID); err != nil {
		s.log.Error("DeleteCallbackInfo-->DeleteInfo", zap.Int64("id", req.ID), zap.Error(err))
		return err
	}
	return nil
}

func (s *CallbackService) QueryCallbackLogList(context *gin.Context, req *api.QueryCallbackLogListReq) (*api.CallbackLogListTable, error) {
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

func (s *CallbackService) Callback(c *gin.Context, req *api.CallbackReq) error {
	var xmlData string
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		s.log.Error("Callback-->ReadAll", zap.Error(err))
		return err
	}
	xmlData = string(data)
	// 根据callbackId查询
	info, err := s.uc.QueryInfoByCallbackId(req.CallbackId)
	if err != nil {
		s.log.Error("Callback-->QueryInfoByCallbackId", zap.Error(err))
		return err
	}
	body, err := code.Decrypt(&code.OpenConf{
		AppId:       info.AppId,
		VerifyToken: info.VerifyToken,
		SecretKey:   info.SecretKey,
	}, xmlData)
	if err != nil {
		s.log.Error("Callback-->Decrypt", zap.String("xmlData", xmlData), zap.Error(err))
		return err
	}
	callbackLog := &biz.CallbackLog{
		CallbackId: req.CallbackId,
		IP:         c.ClientIP(),
		MsgBody:    body,
	}
	if err := s.uc.CreateLog(callbackLog); err != nil {
		s.log.Error("Callback-->CreateLog", zap.Any("callback_log", callbackLog), zap.Error(err))
		return err
	}
	return nil
}

func (s *CallbackService) QueryCallbackInfoList(_ *gin.Context, req *api.QueryCallbackInfoListReq) (*api.CallbackInfoListTable, error) {
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

func (s *CallbackService) CreateCallbackID(_ *gin.Context, req *api.CreateCallbackIDReq) (*api.CallbackIdResp, error) {
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
