package http

import (
	"cloud_callback/internal/pkg/http/middleware"
	"cloud_callback/internal/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/yudeguang/ratelimit"
	"go.uber.org/zap"
	"time"
)

// ServerOption setting Server fields
type ServerOption func(*Server)

// WithPort with port to server
func WithPort(port int) ServerOption {
	return func(s *Server) {
		s.port = port
	}
}

// WithIPAddr with ip addr to server
func WithIPAddr(ipAddr string) ServerOption {
	return func(s *Server) {
		s.ipAddr = ipAddr
	}
}

// WithTimeout with timeout per second
func WithTimeout(timeout int) ServerOption {
	return func(s *Server) {
		s.timeout = time.Duration(timeout) * time.Second
	}
}

// WithMiddleware with group middleware
func WithMiddleware(middleware gin.HandlerFunc) ServerOption {
	return func(s *Server) {
		s.Use(middleware)
	}
}

// WithMode setMode
func WithMode(mode string) ServerOption {
	return func(s *Server) {
		s.mode = mode
	}
}

// Add IP visit limit
func WithIpLimit(rule *ratelimit.Rule) ServerOption {
	return func(s *Server) {
		s.Use(middleware.IPLimitMiddleware(rule))
	}
}

func WithIpLimitOfRate(rates ...middleware.Rate) ServerOption {
	return func(s *Server) {
		r := ratelimit.NewRule()
		for _, rate := range rates {
			r.AddRule(rate.DefaultExpiration, rate.NumberOfAllowedAccesses, rate.EstimatedNumberOfOnlineUserNum...)
		}
		s.Use(middleware.IPLimitMiddleware(r))
	}
}

func WithIpLimitOfFile(fileName string, fileSaveSec time.Duration, rates ...middleware.Rate) ServerOption {
	return func(s *Server) {
		r := ratelimit.NewRule()
		for _, rate := range rates {
			r.AddRule(rate.DefaultExpiration, rate.NumberOfAllowedAccesses, rate.EstimatedNumberOfOnlineUserNum...)
		}
		if err := r.LoadingAndAutoSaveToDisc(fileName, fileSaveSec); err != nil {
			log.Error("WithIpLimitOfFile-->LoadingAndAutoSaveToDisc", zap.Error(err))
			return
		}
		s.Use(middleware.IPLimitMiddleware(r))
	}
}
