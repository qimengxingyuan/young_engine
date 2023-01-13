package handler

import (
	"github.com/cloudwego/hertz/pkg/app"
)

const (
	SuccessMsg = "success"

	SuccessCode      = 0
	ServiceErrCode   = 10001
	ParamErrCode     = 10002
	CompileErrCode   = 20001
	RuleNotExistCode = 20002
	RuleExecErrCode  = 20003
)

type BaseResp struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func BindResp(c *app.RequestContext, code int, msg string, data interface{}) {
	c.JSON(200, BaseResp{
		Code:    code,
		Message: msg,
		Data:    data,
	})
}
