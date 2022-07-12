package admin

import (
	"github.com/gogf/gf/v2/frame/g"
	"goframe/internal/model/entity"
)

type LoginReq struct {
	g.Meta   `path:"/Login/Sign" method:"post" tags:"登录" summary:"提交登录"`
	Phone string `json:"phone" v:"required#手机号不能为空" dc:"手机号"`
	Password string `json:"password" v:"required#密码不能为空" dc:"密码"`
}

type LoginRes struct {
	AdminInfo entity.Admin `json:"admin_info"`
	Token string `json:"token" dc:"登录token"`
}

//  注销登录
type LoginLogoutReq struct {
	g.Meta `path:"/Login/Logout" method:"post" tags:"登录" summary:"注销登录"`
}

type LoginLogoutRes struct{}