package api

import (
	"adolesce/internal/pkg/http"
	"github.com/gin-gonic/gin"
	nhttp "net/http"
)

func InitRouter(s *http.Server, srv CallbackHTTPServer) {
	s.Static("/web", "./api/callback/public")
	in := s.Group("/inner")
	{
		in.POST("/create", CallbackHandler_Create(srv))                   // 根据需要的信息创建回调id
		in.POST("/info_list", CallbackHandler_QueryCallbackInfoList(srv)) // 查询回调ID列表
		in.POST("/log_list", CallbackHandler_QueryCallbackLogList(srv))   // 查询回调日志列表
		in.POST("/info_delete", CallbackHandler_DeleteCallbackInfo(srv))  // 删除对应日志
		in.GET("/test", func(c *gin.Context) {
			c.JSON(nhttp.StatusOK, gin.H{
				"name": "test",
			})
		})
	}
}