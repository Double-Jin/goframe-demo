package admin

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	apiAdmin "goframe/api/v1/admin"
	"goframe/internal/consts"
	"goframe/internal/dao"
	"goframe/internal/service"
)

type sAdmin struct {}

func init()  {
	service.RegisterAdmin(New())
}

func New() *sAdmin {
	return &sAdmin{}
}


func (s *sAdmin) GetInfo(ctx context.Context, req *apiAdmin.AdminReq) (res *apiAdmin.AdminRes, err error) {

	err = dao.Admin.Ctx(ctx).Where("id=?", 1).Scan(&res)

	if err != nil {
		err = gerror.NewCode(gcode.New(consts.AccountPassWordFailed.Code(), consts.AccountPassWordFailed.Desc(), nil))
		//fmt.Println(err)
		return nil, err
	}

	return
}
