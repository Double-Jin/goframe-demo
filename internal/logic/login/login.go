package login

import (
	"context"
	apiAdmin "goframe/api/v1/admin"
	"goframe/internal/common"
	"goframe/internal/consts"
	"goframe/internal/dao"
	"goframe/internal/model/entity"
	"goframe/internal/service"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
)

type sLogin struct {}

func init()  {
	service.RegisterLogin(New())
}

func New() *sLogin {
	return &sLogin{}
}


func (s *sLogin) Sign(ctx context.Context, req *apiAdmin.LoginReq) (res *apiAdmin.LoginRes, err error) {

	var adminInfo *entity.Admin

	err = dao.Admin.Ctx(ctx).Where("phone=?", req.Phone).Scan(&adminInfo)

	if err != nil {
		err = gerror.Wrap(err, consts.SqlError.Desc())
		return nil, err
	}

	if adminInfo == nil {
		err = gerror.New(consts.SqlError.Desc())
		return
	}

	if adminInfo.Password != gmd5.MustEncryptString(req.Password) {
		err = gerror.New("用户密码不正确")
		return
	}

	token, err := common.Jwt.GenerateAdminLoginToken(ctx, adminInfo, false)

	if err != nil {
		err = gerror.New(err.Error())
		return
	}

	res = &apiAdmin.LoginRes{
		AdminInfo: *adminInfo,
		Token:     gconv.String(token),
	}

	return
}

func (s *sLogin) Logout(ctx context.Context, req *apiAdmin.LoginLogoutReq) (res *apiAdmin.LoginLogoutRes, err error) {

	var authorization = common.Jwt.GetAuthorization(common.Context.Get(ctx).Request)

	// TODO  获取jwtToken
	jwtToken := consts.RedisJwtToken + gmd5.MustEncryptString(authorization)
	if len(jwtToken) == 0 {
		err = gerror.New("当前用户未登录！")
		return res, err
	}

	// TODO  删除登录token
	cache := common.Cache.New()
	_, err = cache.Remove(ctx, jwtToken)
	if err != nil {
		return res, err
	}

	return
}
