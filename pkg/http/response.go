package http

import (
	"cloud_callback/internal/pkg/slerror"
	"github.com/gin-gonic/gin"
	"net/http"
)

type httpResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func WriteSuccessResp(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, httpResponse{
		Code: slerror.SuccessCode,
		Data: data,
	})
}


func WriteFailResp(c *gin.Context, code int, err error) {
	c.JSON(http.StatusInternalServerError, httpResponse{
		Code: code,
		Msg: err.Error(),
	})
}
