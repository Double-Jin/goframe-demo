package router

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"goframe/internal/controller/admin"
)

func BindAdminController(group *ghttp.RouterGroup) {
	group.Group("/Admin", func(group *ghttp.RouterGroup) {
		//group.Middleware(middleware.AdminAuth)
		group.Bind(
			admin.Admin,
		)
	})
}
