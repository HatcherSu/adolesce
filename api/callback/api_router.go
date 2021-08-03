package api

import (
	"cloud_callback/internal/pkg/http"
	"cloud_callback/internal/pkg/log"
	"cloud_callback/internal/pkg/slerror"
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

type CallbackHTTPServer interface {
	CreateCallbackID(*gin.Context, *CreateCallbackIDReq) (*CallbackIdResp, error)
	QueryCallbackInfoList(*gin.Context, *QueryCallbackInfoListReq) (*CallbackInfoListTable, error)
	QueryCallbackLogList(*gin.Context, *QueryCallbackLogListReq) (*CallbackLogListTable, error)
	DeleteCallbackInfo(*gin.Context, *DeleteCallbackInfoReq) error
}

// todo 生成Handler
func CallbackHandler_Create(srv CallbackHTTPServer) func(*gin.Context) {
	return func(c *gin.Context) {
		var in CreateCallbackIDReq
		if err := c.ShouldBind(&in); err != nil {
			http.WriteFailResp(c, slerror.InvalidParamErrCode, err)
			return
		}
		resp, err := srv.CreateCallbackID(c, &in)
		if err != nil {
			http.WriteFailResp(c, slerror.InnerServerErrCode, err)
			return
		}
		http.WriteSuccessResp(c, resp)
	}
}

func CallbackHandler_QueryCallbackInfoList(srv CallbackHTTPServer) func(*gin.Context) {
	return func(c *gin.Context) {
		var in QueryCallbackInfoListReq
		if err := c.ShouldBind(&in); err != nil {
			http.WriteFailResp(c, slerror.InvalidParamErrCode, err)
			return
		}
		resp, err := srv.QueryCallbackInfoList(c, &in)
		if err != nil {
			http.WriteFailResp(c, slerror.InnerServerErrCode, err)
			return
		}
		log.Info("")
		c.JSON(nhttp.StatusOK, resp)
	}
}

func CallbackHandler_QueryCallbackLogList(srv CallbackHTTPServer) func(*gin.Context) {
	return func(c *gin.Context) {
		var in QueryCallbackLogListReq
		if err := c.ShouldBind(&in); err != nil {
			http.WriteFailResp(c, slerror.InvalidParamErrCode, err)
			return
		}
		resp, err := srv.QueryCallbackLogList(c, &in)
		if err != nil {
			http.WriteFailResp(c, slerror.InnerServerErrCode, err)
			return
		}
		c.JSON(nhttp.StatusOK, resp)
	}
}

func CallbackHandler_DeleteCallbackInfo(srv CallbackHTTPServer) func(*gin.Context) {
	return func(c *gin.Context) {
		var in DeleteCallbackInfoReq
		if err := c.ShouldBind(&in); err != nil {
			http.WriteFailResp(c, slerror.InvalidParamErrCode, err)
			return
		}
		err := srv.DeleteCallbackInfo(c, &in)
		if err != nil {
			http.WriteFailResp(c, slerror.InnerServerErrCode, err)
			return
		}
		http.WriteSuccessResp(c, nil)
	}
}
