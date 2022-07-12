package router

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"goframe/internal/controller/admin"
	"goframe/internal/middleware"
)

func BindAdminController(group *ghttp.RouterGroup) {
	group.Group("/Admin", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.AdminAuth)
		group.Bind(
			admin.Admin,
		)
	})
}
