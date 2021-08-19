package api

import (
	"adolesce/internal/pkg/slerror"
	"adolesce/pkg/http"
	"adolesce/pkg/log"
	"github.com/gin-gonic/gin"
	nhttp "net/http"
)

type DeleteCallbackInfoReq struct {
	ID int64 `json:"ID"`
}

type CallbackLogListTable struct {
	Code  int64          `protobuf:"varint,1,opt,name=code,proto3" json:"code"`
	Msg   string         `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg"`
	Count int64          `protobuf:"varint,3,opt,name=count,proto3" json:"count"`
	Data  []*CallbackLog `protobuf:"bytes,4,rep,name=data,proto3" json:"data"`
}

// 回调日志
type CallbackLog struct {
	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID"`
	// 回调ID
	CallbackId string `protobuf:"bytes,2,opt,name=callback_id,json=callbackId,proto3" json:"callback_id"`
	// 请求IP地址
	IP string `protobuf:"bytes,3,opt,name=IP,proto3" json:"IP"`
	// 消息体
	MsgBody string `protobuf:"bytes,4,opt,name=msg_body,json=msgBody,proto3" json:"msg_body"`
}

type QueryCallbackLogListReq struct {
	Page       int64  `protobuf:"varint,1,opt,name=page,proto3" json:"page"`
	Limit      int64  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit"`
	CallbackId string `protobuf:"bytes,3,opt,name=callback_id,json=callbackId,proto3" json:"callback_id"`
}

type CallbackReq struct {
	CallbackId string `protobuf:"bytes,1,opt,name=callback_id,json=callbackId,proto3" json:"callback_id" uri:"callback_id"`
}

// 查询callback info 请求
type QueryCallbackInfoListReq struct {
	Page  int64 `protobuf:"varint,1,opt,name=page,proto3" json:"page"`
	Limit int64 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit"`
}

// 回调idList,layui格式
type CallbackInfoListTable struct {
	Code  int64           `protobuf:"varint,1,opt,name=code,proto3" json:"code"`
	Msg   string          `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg"`
	Count int64           `protobuf:"varint,3,opt,name=count,proto3" json:"count"`
	Data  []*CallbackInfo `protobuf:"bytes,4,rep,name=data,proto3" json:"data"`
}

// 回调
type CallbackInfo struct {
	ID          int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID"`
	CallbackId  string `protobuf:"bytes,2,opt,name=callback_id,json=callbackId,proto3" json:"callback_id"`
	AppId       string `protobuf:"bytes,3,opt,name=app_id,json=appId,proto3" json:"app_id"`
	VerifyToken string `protobuf:"bytes,4,opt,name=verify_token,json=verifyToken,proto3" json:"verify_token"`
	SecretKey   string `protobuf:"bytes,5,opt,name=secret_key,json=secretKey,proto3" json:"secret_key"`
}

// 创建回调ID的请求
type CreateCallbackIDReq struct {
	// 商铺ID
	AppId string `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id"`
	// 店铺ID
	ClientId string `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id"`
	// 密钥
	SecretKey string `protobuf:"bytes,3,opt,name=secret_key,json=secretKey,proto3" json:"secret_key"`
	// token
	VerifyToken string `protobuf:"bytes,4,opt,name=verify_token,json=verifyToken,proto3" json:"verify_token"`
}

// 回调ID响应
type CallbackIdResp struct {
	// 回调id
	CallbackId string `protobuf:"bytes,1,opt,name=callback_id,json=callbackId,proto3" json:"callback_id"`
}

// 回调请求体
type CreateCallbackIDRequest struct {
	// 	加密字符串
	Encrypt string `protobuf:"bytes,1,opt,name=encrypt,proto3" json:"encrypt"`
	// 	消息体签名
	MsgSignature string `protobuf:"bytes,2,opt,name=msg_signature,json=msgSignature,proto3" json:"msg_signature"`
	// 时间戳
	TimeStamp int64 `protobuf:"varint,3,opt,name=time_stamp,json=timeStamp,proto3" json:"time_stamp"`
	// 随机数nonce
	Nonce int64 `protobuf:"varint,4,opt,name=nonce,proto3" json:"nonce"`
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
