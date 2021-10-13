package server

import (
	api "adolesce/api/callback"
	"adolesce/internal/conf"
	"adolesce/internal/pkg/http"
)

func NewHTTPServer(config *conf.Configs, cbServer api.CallbackHTTPServer) (*http.Server, func(), error) {
	var opts = []http.ServerOption{
		http.WithPort(config.HttpServerConfig.HttpPort),
		http.WithTimeout(config.HttpServerConfig.DialTimeoutSecond),
		http.WithIPAddr(config.HttpServerConfig.HttpIPAddr),
		http.WithMode(config.HttpServerConfig.HttpMode),
	}
	srv := http.NewServer(opts...)
	api.InitRouter(srv, cbServer)
	return srv, func() {
		srv.Close()
	}, nil
}
