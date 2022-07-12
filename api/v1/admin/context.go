package admin

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"goframe/internal/model/entity"
)

// Context 请求上下文结构
type Context struct {
	ReqId       string         // 上下文ID
	TakeUpTime  int64          // 请求耗时 ms
	Request     *ghttp.Request // 当前Request管理对象
	Admin       *entity.Admin      // 上下文用户信息
	ComResponse *Response      // 组件响应
	Data        g.Map          // 自定KV变量，业务模块根据需要设置，不固定
}

