package service

import (
	api "adolesce/api/app-timer"
	"adolesce/internal/biz"
	"adolesce/pkg/log"
)

// 实现接口
var _ api.ExampleTimerHandler = (*exampleTimerService)(nil)

func NewExampleTimerService(log log.Logger, uc biz.CallbackUsecase) api.ExampleTimerHandler {
	return &exampleTimerService{
		log: log,
	}
}

type exampleTimerService struct {
	log log.Logger
}

func (t *exampleTimerService) DeleteExample() {
}
