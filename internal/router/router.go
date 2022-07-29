package router

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"goframe/internal/middleware"
)

func BindController(group *ghttp.RouterGroup) {
	group.Middleware(middleware.HandlerResponse)
	group.Group("/", func(group *ghttp.RouterGroup) {
		// 绑定后台路由
		BindAdminController(group)
	})

}
