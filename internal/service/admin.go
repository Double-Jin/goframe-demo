package service

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	apiAdmin "goframe/api/v1/admin"
	"goframe/internal/consts"
	"goframe/internal/dao"
)

type sAdmin interface {
	GetInfo(ctx context.Context, req *apiAdmin.AdminReq) (res *apiAdmin.AdminRes, err error)
}

type adminImpl struct {
}

var adminService = adminImpl{}

func Admin() sAdmin {
	return &adminService
}

func (s *adminImpl) GetInfo(ctx context.Context, req *apiAdmin.AdminReq) (res *apiAdmin.AdminRes, err error) {

	err = dao.Admin.Ctx(ctx).Where("id=?", 1).Scan(&res)

	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return nil, err

	}

	return
}
