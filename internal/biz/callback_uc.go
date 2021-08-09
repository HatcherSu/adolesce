package biz

import "adolesce/pkg/log"

type ExampleUsecase struct {
	logRepo  CallbackLogRepo
	infoRepo CallbackInfoRepo
	log      log.Logger
}

func NewExampleUsecase(logRepo CallbackLogRepo, infoRepo CallbackInfoRepo, log log.Logger) ExampleUsecase {
	return ExampleUsecase{logRepo, infoRepo, log}
}

func (uc *ExampleUsecase) CreateInfo(info *CallbackInfo) error {
	return uc.infoRepo.Create(info)
}

func (uc *ExampleUsecase) QueryInfoList(filter *CallbackInfoFilter) ([]*CallbackInfo, error) {
	return uc.infoRepo.QueryList(filter)
}

func (uc *ExampleUsecase) QueryInfoByCallbackId(callbackId string) (*CallbackInfo, error) {
	return uc.infoRepo.QueryByCallbackId(callbackId)
}

func (uc *ExampleUsecase) CreateLog(log *CallbackLog) error {
	return uc.logRepo.Create(log)
}

func (uc *ExampleUsecase) QueryLogList(filter *CallbackLogFilter) ([]*CallbackLog, error) {
	return uc.logRepo.QueryList(filter)
}

func (uc *ExampleUsecase) DeleteInfo(id int64) error {
	return uc.infoRepo.DeleteByID(id)
}
