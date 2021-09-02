package service

import (
	api "adolesce/api/app-timer"
	"adolesce/pkg/log"
)

// 实现接口
var _ api.ExampleTimerHandler = (*exampleTimerService)(nil)

func NewExampleTimerService(log log.Logger) api.ExampleTimerHandler {
	return &exampleTimerService{
		log: log,
	}
}

type exampleTimerService struct {
	log log.Logger
}

func (t *exampleTimerService) DeleteExample() {
	log.Info("delete example")
}
