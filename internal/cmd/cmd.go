package cmd

import (
	"context"
	"goframe/internal/middleware"
	"goframe/internal/router"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()

			s.Group("/", func(group *ghttp.RouterGroup) {
				// 注册全局中间件
				group.Middleware(
					middleware.Middleware().Ctx, //必须第一个加载
					middleware.Middleware().CORS,
					ghttp.MiddlewareHandlerResponse,
				)

				router.BindController(group)
			})
			s.Run()
			return nil
		},
	}
)
