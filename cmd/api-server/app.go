package main

import (
	"adolesce/internal/pkg/app"
	"adolesce/internal/pkg/http"
	"adolesce/pkg/log"
	"go.uber.org/zap"
)

func newApp(logger log.Logger, hs *http.Server) *app.App {
	application := app.NewApp("cloud_callback",
		app.WithShort("Cloud Callback Server"),
		app.WithLong("The Callback Server is for the message consumer"),
		app.WithRunFunc(runServer(logger, hs)))
	return application
}

func runServer(logger log.Logger, hs *http.Server) app.RunFunc {
	return func(name string) error {
		logger.Info("server start", zap.String("name", name))

		if err := hs.Run(); err != nil {
			logger.Error("server run error", zap.Error(err))
		}
		return nil
	}
}
