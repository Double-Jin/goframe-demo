package admin

import (
	"context"
	"goframe/api/v1/admin"
	"goframe/internal/service"
)

var (
	Login = cLogin{}
)

type cLogin struct{}

func (c *cAdmin) Sign(ctx context.Context, req *admin.LoginReq) (res *admin.LoginRes, err error) {
	res, err = service.Login().Sign(ctx, req)
	return
}

func (c *cAdmin) Logout(ctx context.Context, req *admin.LoginLogoutReq) (res *admin.LoginLogoutRes, err error) {
	res, err = service.Login().Logout(ctx, req)
	return
}
