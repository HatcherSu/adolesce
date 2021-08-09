package server

import (
	api "adolesce/api/app-timer"
	"adolesce/pkg/log"
	"adolesce/pkg/log/cron_log"
	"github.com/robfig/cron/v3"
)

func NewTimerServer(logger log.Logger, srv api.ExampleTimerHandler) (*cron.Cron, func(), error) {
	c := cron.New(cron.WithLogger(cron_log.NewCronLogger(logger)))
	err := api.InitCron(c, logger, srv)
	return c, func() {
		c.Stop()
	}, err
}
