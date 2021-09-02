//go:generate go run github.com/google/wire/cmd/wire
//+build wireinject

package main

import (
	"adolesce/internal/conf"
	"adolesce/internal/pkg/app"
	"adolesce/internal/server"
	"adolesce/internal/service"
	"adolesce/pkg/log"
	"github.com/google/wire"
)

func initTimer(*conf.Configs, log.Logger) (*app.App, func(), error) {
	panic(wire.Build(server.ProviderSet, service.ProviderSet, newTimerServer))
}
