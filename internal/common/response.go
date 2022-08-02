package common

import (
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"goframe/api/v1/admin"
	"time"
)

type DefaultResponse struct {
	Code    int         `json:"code"    dc:"Error code"`
	Message string      `json:"message" dc:"Error message"`
	Data    interface{} `json:"data"    dc:"Result data for certain request according API definition"`
}

// 统一响应
var Response = new(response)

type response struct{}


func (component *response) SuccessJson(r *ghttp.Request) {
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

	// TODO  清空响应
	r.Response.ClearBuffer()

	// TODO  写入响应
	r.Response.WriteJson(defaultResponse)

}


func (component *response) LoginFailJson(r *ghttp.Request, code int, message string) {

	Res := &admin.Response{
		Code:      code,
		Message:   message,
		Timestamp: time.Now().Unix(),
		ReqId:     Context.Get(r.Context()).ReqId,
	}

	defaultResponse := DefaultResponse{
		Code:    code,
		Message: message,
		Data:    "",
	}

	// TODO  清空响应
	r.Response.ClearBuffer()

	// TODO  写入响应
	r.Response.WriteJson(defaultResponse)

	// TODO  加入到上下文
	Context.SetResponse(r.Context(), Res)

}


