package admin

import (
	"github.com/gogf/gf/v2/frame/g"
	"goframe/internal/model/entity"
)

type AdminReq struct {
	g.Meta `path:"/Admin/GetInfo" tags:"Admin" method:"get" summary:"You first admin api"`
}

type AdminRes struct {
	//g.Meta `mime:"application/json"`
	*entity.Admin
}
