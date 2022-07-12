package router

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

func BindController(group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		// 绑定后台路由
		BindAdminController(group)
	})

}
