package common

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"goframe/api/v1/admin"
	"goframe/internal/consts"
	"time"
)

// 统一响应
var Response = new(response)

type response struct{}

func (component *response) JsonExit(r *ghttp.Request, code int, message string, data ...interface{}) {
	component.RJson(r, code, message, data...)
	r.Exit()
}


func (component *response) RJson(r *ghttp.Request, code int, message string, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}

	Res := &admin.Response{
		Code:      code,
		Message:   message,
		Timestamp: time.Now().Unix(),
		ReqId:     Context.Get(r.Context()).ReqId,
	}

	// TODO  如果不是正常的返回，则将data转为error
	if consts.CodeOK == code {
		Res.Data = responseData
	} else {
		Res.Error = responseData
	}

	// TODO  清空响应
	r.Response.ClearBuffer()

	// TODO  写入响应
	r.Response.WriteJson(Res)

	// TODO  加入到上下文
	Context.SetResponse(r.Context(), Res)
}


func (component *response) SusJson(isExit bool, r *ghttp.Request, message string, data ...interface{}) {
	if isExit {
		component.JsonExit(r, consts.CodeOK, message, data...)
	}
	component.RJson(r, consts.CodeOK, message, data...)
}

func (component *response) FailJson(isExit bool, r *ghttp.Request, message string, data ...interface{}) {
	if isExit {
		component.JsonExit(r, consts.CodeNil, message, data...)
	}
	component.RJson(r, consts.CodeNil, message, data...)
}


func (component *response) Redirect(r *ghttp.Request, location string, code ...int) {
	r.Response.RedirectTo(location, code...)
}

func (component *response) Download(r *ghttp.Request, location string, code ...int) {
	r.Response.ServeFileDownload("test.txt")
}
