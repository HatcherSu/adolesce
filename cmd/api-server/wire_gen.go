// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"adolesce/internal/biz"
	"adolesce/internal/conf"
	"adolesce/internal/data"
	"adolesce/internal/pkg/app"
	"adolesce/internal/server"
	"adolesce/internal/service"
	"adolesce/pkg/log"
)

// Injectors from wire.go:

func initApp(configs *conf.Configs, logger log.Logger) (*app.App, func(), error) {
	dataData, cleanup, err := data.NewData(configs, logger)
	if err != nil {
		return nil, nil, err
	}
	callbackLogRepo := data.NewCallbackLogRepo(dataData, logger)
	callbackInfoRepo := data.NewCallbackInfoRepo(dataData, logger)
	exampleUsecase := biz.NewExampleUsecase(callbackLogRepo, callbackInfoRepo, logger)
	callbackHTTPServer := service.NewCallbackService(logger, exampleUsecase)
	httpServer, cleanup2, err := server.NewHTTPServer(configs, callbackHTTPServer)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	appApp := newApp(logger, httpServer)
	return appApp, func() {
		cleanup2()
		cleanup()
	}, nil
}
