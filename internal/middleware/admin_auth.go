//
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package middleware

import (
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe/api/v1/admin"
	"goframe/internal/common"
	"goframe/internal/consts"
	"goframe/internal/model/entity"
	"goframe/utility"
)

func AdminAuth(r *ghttp.Request) {

	var (
		ctx           = r.Context()
		adminInfo          = new(entity.Admin)
		authorization = common.Jwt.GetAuthorization(r)
	)

	// TODO  替换掉模块前缀
	routerPrefix, _ := g.Cfg().Get(ctx, "router.admin.prefix", "/Admin")
	path := gstr.Replace(r.URL.Path, routerPrefix.String(), "", 1)

	/// TODO  不需要验证登录的路由地址
	if utils.Auth.IsExceptLogin(ctx, path) {
		r.Middleware.Next()
	}
	if authorization == "" {
		common.Response.LoginFailJson(r,consts.TokenExpired.Code(),consts.TokenExpired.Desc())
		return
	}

	// TODO  获取jwtToken
	jwtToken := consts.RedisJwtToken + gmd5.MustEncryptString(authorization)
	jwtSign, _ := g.Cfg().Get(ctx, "jwt.sign", "jin")

	data, ParseErr := common.Jwt.ParseToken(authorization, jwtSign.Bytes())
	if ParseErr != nil {
		common.Response.LoginFailJson(r,consts.TokenExpired.Code(),consts.TokenExpired.Desc())
		return

	}

	parseErr := gconv.Struct(data, &adminInfo)
	if parseErr != nil {
		common.Response.LoginFailJson(r,consts.TokenFail.Code(),consts.TokenFail.Desc())
		return
	}

	// TODO  判断token跟redis的缓存的token是否一样
	cache := common.Cache.New()
	isContains, containsErr := cache.Contains(ctx, jwtToken)
	if containsErr != nil {
		common.Response.LoginFailJson(r,consts.TokenFail.Code(),consts.TokenFail.Desc())
		return
	}
	if !isContains {
		common.Response.LoginFailJson(r,consts.TokenFail.Code(),consts.TokenFail.Desc())
		return
	}

	// TODO  保存到上下文
	customCtx := &admin.Context{}
	if adminInfo != nil {
		customCtx.Admin = &entity.Admin{
			Id:         adminInfo.Id,
			LoginName:   adminInfo.LoginName,
			AdminName:   adminInfo.AdminName,
			AdminImg:     adminInfo.AdminImg,
			Phone:      adminInfo.Phone,
		}
	}

	common.Context.SetAdmin(ctx, customCtx.Admin)

	r.Middleware.Next()
}
