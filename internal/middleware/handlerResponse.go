package middleware

import (
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
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
	var (
		err  = r.GetError()
		res  = r.GetHandlerResponse()
		code = gerror.Code(err)
	)
	defaultResponse := DefaultResponse{
		Code:    code.Code(),
		Message: code.Message(),
		Data:    res,
	}

	r.Response.WriteJson(defaultResponse)
}
