package http

import (
	"cloud_callback/internal/pkg/http/middleware"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/yudeguang/ratelimit"
	"net"
	"net/http"
	"strconv"
	"time"
)

const (
	DefaultServerPort = 8081
	DefaultTimeout    = 10
)

var (
	PortNotSetError = errors.New("http port not set")
)

type Server struct {
	*gin.Engine
	mode       string
	ipAddr     string        // ip地址
	port       int           // 端口
	timeout    time.Duration // 超时
	httpServer *http.Server
	rule       *ratelimit.Rule
}

func NewServer(opts ...ServerOption) *Server {
	s := &Server{
		Engine: gin.New(),
		port:   DefaultServerPort,
		mode:   gin.ReleaseMode,
	}
	for _, o := range opts {
		o(s)
	}
	// 设置middleware
	s.Use(gin.Recovery())
	s.Use(middleware.LoggerMiddleware())
	if s.timeout > 0 {
		s.Use(middleware.TimeoutMiddleware(s.timeout))
	}
	r := ratelimit.NewRule()
	r.AddRule(time.Second*5, 5)
	gin.SetMode(s.mode)
	return s
}

func DefaultServer() *Server {
	s := &Server{
		Engine:  gin.Default(),
		port:    DefaultServerPort,
		timeout: DefaultTimeout,
		mode:    gin.ReleaseMode,
	}
	s.Use(middleware.TimeoutMiddleware(s.timeout))
	return s
}

func (s *Server) Run() error {
	s.httpServer = &http.Server{
		Addr:    net.JoinHostPort(s.ipAddr, strconv.Itoa(s.port)),
		Handler: s,
	}
	// todo log
	if err := s.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		// todo log
		return err
	}

	// todo log
	return nil
}

func (s *Server) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		// todo log
	}

}
