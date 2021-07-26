//go:generate go run github.com/google/wire/cmd/wire
//+build wireinject

package main

import (
	"cloud_callback/internal/biz"
	"cloud_callback/internal/conf"
	"cloud_callback/internal/data"
	"cloud_callback/internal/pkg/app"
	"cloud_callback/internal/pkg/log"
	"cloud_callback/internal/server"
	"cloud_callback/internal/service"
	"github.com/google/wire"
)

func initApp(*conf.Configs, log.Logger) (*app.App, func(), error) {
	panic(wire.Build(server.ProviderSet, service.ProviderSet, biz.ProviderSet, data.ProviderSet, newApp))
}
