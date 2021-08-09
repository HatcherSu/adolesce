package main

import (
	"adolesce/internal/pkg/app"
	"adolesce/pkg/log"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

func newTimerServer(logger log.Logger, c *cron.Cron) (*app.App, func(), error) {
	application := app.NewApp("App Timer",
		app.WithShort("App Timer Server"),
		app.WithLong("This is a timer for running func in time"),
		app.WithRunFunc(runTimer(logger, c)))
	return application, func() {
		c.Stop()
	}, nil
}

func runTimer(logger log.Logger, c *cron.Cron) app.RunFunc {
	return func(name string) error {
		logger.Info("timer server start", zap.String("name", name))
		c.Run()
		return nil
	}
}
