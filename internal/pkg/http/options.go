package http

import (
	"github.com/gin-gonic/gin"
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
