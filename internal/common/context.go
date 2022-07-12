package common

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"goframe/api/v1/admin"
	"goframe/internal/consts"
	"goframe/internal/model/entity"
)

// 上下文
var Context = new(commonContext)

type commonContext struct{}

func (component *commonContext) Init(r *ghttp.Request, customCtx *admin.Context) {
	r.SetCtxVar(consts.ContextKey, customCtx)
}

func (component *commonContext) Get(ctx context.Context) *admin.Context {
	value := ctx.Value(consts.ContextKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(*admin.Context); ok {
		return localCtx
	}

	return nil
}

func (component *commonContext) SetAdmin(ctx context.Context, admin *entity.Admin) {
	component.Get(ctx).Admin = admin
}

func (component *commonContext) SetResponse(ctx context.Context, response *admin.Response) {
	component.Get(ctx).ComResponse = response
}

func (component *commonContext) SetTakeUpTime(ctx context.Context, takeUpTime int64) {
	component.Get(ctx).TakeUpTime = takeUpTime
}

func (component *commonContext) GetUserId(ctx context.Context) uint64 {
	admin := component.Get(ctx).Admin

	if admin == nil {
		return 0
	}

	return admin.Id
}
