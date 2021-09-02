package api

import (
	"adolesce/pkg/log"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

type cronCmd struct {
	Cron    string
	Name    string
	Desc    string
	RunFunc func()
}

func InitCron(c *cron.Cron, logger log.Logger, srv ExampleTimerHandler) error {
	for _, cmd := range []cronCmd{
		{
			Cron:    "01 * * * *",
			Name:    "DeleteExample",
			Desc:    "delete example in 00:00",
			RunFunc: srv.DeleteExample,
		},
	} {
		if _, err := c.AddFunc(cmd.Cron, cmd.RunFunc); err != nil {
			logger.Error("InitCron-->AddFunc", zap.String("name", cmd.Name), zap.Error(err))
			return err
		}
	}
	return nil
}

type ExampleTimerHandler interface {
	DeleteExample()
}
