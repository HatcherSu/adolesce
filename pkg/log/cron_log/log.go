package cron_log

import "adolesce/pkg/log"

type cronLogger struct {
	logger log.Logger
}

func (l cronLogger) Info(msg string, keysAndValues ...interface{}) {
	l.logger.Infow(msg, keysAndValues)
}

func (l cronLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	l.logger.Errorw(msg, keysAndValues, "error", err)
}

func NewCronLogger(logger log.Logger) *cronLogger {
	return &cronLogger{logger}
}
