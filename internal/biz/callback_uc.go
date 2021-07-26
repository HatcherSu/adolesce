package biz

import "cloud_callback/internal/pkg/log"

type CallbackUsecase struct {
	logRepo  CallbackLogRepo
	infoRepo CallbackInfoRepo
	log      log.Logger
}

func NewCallbackLogUsecase(logRepo CallbackLogRepo, infoRepo CallbackInfoRepo, log log.Logger) CallbackUsecase {
	return CallbackUsecase{logRepo, infoRepo, log}
}

func (uc *CallbackUsecase) CreateInfo(info *CallbackInfo) error {
	return uc.infoRepo.Create(info)
}

func (uc *CallbackUsecase) QueryInfoList(filter *CallbackInfoFilter) ([]*CallbackInfo, error) {
	return uc.infoRepo.QueryList(filter)
}

func (uc *CallbackUsecase) QueryInfoByCallbackId(callbackId string) (*CallbackInfo, error) {
	return uc.infoRepo.QueryByCallbackId(callbackId)
}

func (uc *CallbackUsecase) CreateLog(log *CallbackLog) error {
	return uc.logRepo.Create(log)
}

func (uc *CallbackUsecase) QueryLogList(filter *CallbackLogFilter) ([]*CallbackLog, error) {
	return uc.logRepo.QueryList(filter)
}

func (uc *CallbackUsecase) DeleteInfo(id int64) error {
	return uc.infoRepo.DeleteByID(id)
}
