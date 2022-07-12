package admin

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type UploadFileReq struct {
	g.Meta `path:"/Admin/UploadFile" tags:"File" method:"post" summary:"文件上传"`
	File *ghttp.UploadFile `json:"file" type:"file" dc:"选择上传文件"`
}

type UploadFileRes struct {
	g.Meta `mime:"application/json"`
	Name string `json:"name" dc:"文件名"`
	Url string `json:"url" dc:"地址"`
}
