package server

import (
	api "adolesce/api/app-timer"
	"adolesce/pkg/log"
	"github.com/robfig/cron/v3"
)

func NewTimerServer(logger log.Logger, srv api.ExampleTimerHandler) (*cron.Cron, func(), error) {
	c := cron.New(cron.WithLogger(cron.VerbosePrintfLogger(log.StdInfoLogger())))
	err := api.InitCron(c, logger, srv)
	return c, func() {
		c.Stop()
	}, err
}
