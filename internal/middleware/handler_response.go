package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"goframe/internal/common"
	"goframe/internal/consts"
)

type DefaultResponse struct {
	Code    int         `json:"code"    dc:"Error code"`
	Message string      `json:"message" dc:"Error message"`
	Data    interface{} `json:"data"    dc:"Result data for certain request according API definition"`
}

func HandlerResponse(r *ghttp.Request) {
	r.Middleware.Next()

	if r.Response.BufferLength() > 0 && r.Response.Status == consts.CodeOK.Code() {
		return
	}

	common.Response.SuccessJson(r)
}
