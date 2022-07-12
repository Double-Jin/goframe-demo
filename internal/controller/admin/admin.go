package admin

import (
	"context"
	"goframe/api/v1/admin"
	"goframe/internal/service"
)

var (
	Admin = cAdmin{}
)

type cAdmin struct{}

func (c *cAdmin) GetInfo(ctx context.Context, req *admin.AdminReq) (res *admin.AdminRes, err error) {
	res, err = service.Admin().GetInfo(ctx, req)
	return
}
