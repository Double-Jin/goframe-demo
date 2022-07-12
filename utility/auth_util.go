//
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package utils

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

// 权限认证类
var Auth = new(auth)

type auth struct{}


func (util *auth) IsExceptAuth(ctx context.Context, path string) bool {

	var pathList []string

	except, _ := g.Cfg().Get(ctx, "router.admin.exceptAuth")
	pathList = except.Strings()

	for i := 0; i < len(pathList); i++ {
		if Charset.IsExists(pathList[i], path) {
			return true
		}
	}

	return false
}

func (util *auth) IsExceptLogin(ctx context.Context, path string) bool {

	var pathList []string

	except, _ := g.Cfg().Get(ctx, "router.admin.exceptLogin")
	pathList = except.Strings()

	for i := 0; i < len(pathList); i++ {
		if Charset.IsExists(pathList[i], path) {
			return true
		}
	}

	return false
}
