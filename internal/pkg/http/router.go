package http

// todo define router

import "github.com/gin-gonic/gin"

type Router struct {
	prefix      string
	router      *gin.RouterGroup
	middlewares []gin.HandlerFunc
}

func newRouter(prefix string, rg *gin.RouterGroup) *Router {
	return &Router{
		prefix:      prefix,
		router:      rg,
	}
}

func (r *Router) POST(relativePath string, h gin.HandlerFunc) {
	r.router.POST(relativePath, h)
}

func (r *Router) GET(relativePath string, h gin.HandlerFunc) {
	r.router.GET(relativePath, h)
}
