package server

import (
	api "adolesce/api/callback"
	"adolesce/internal/conf"
	"adolesce/internal/pkg/http"
)

func NewHTTPServer(config *conf.Configs, cbServer api.CallbackHTTPServer) (*http.Server, func(), error) {
	var opts = []http.ServerOption{
		http.WithPort(config.Http.HttpPort),
		http.WithTimeout(config.Http.DialTimeoutSecond),
		http.WithIPAddr(config.Http.HttpIPAddr),
		http.WithMode(config.Http.HttpMode),
	}
	srv := http.NewServer(opts...)
	api.InitRouter(srv, cbServer)
	return srv, func() {
		srv.Close()
	}, nil
}
