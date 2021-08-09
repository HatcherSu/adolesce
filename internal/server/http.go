package server

import (
	api "adolesce/api/callback"
	"adolesce/internal/conf"
	"adolesce/internal/pkg/http"
)

func NewHTTPServer(config *conf.Configs, cbServer api.CallbackHTTPServer) (*http.Server, func(), error) {
	var opts = []http.ServerOption{
		http.WithPort(config.HttpConf.HttpPort),
		http.WithTimeout(config.HttpConf.DialTimeoutSecond),
		http.WithIPAddr(config.HttpConf.HttpIPAddr),
		http.WithMode(config.HttpConf.HttpMode),
	}
	srv := http.NewServer(opts...)
	api.InitRouter(srv, cbServer)
	return srv, func() {
		srv.Close()
	}, nil
}
