// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// JinAdmin is the golang structure of table jin_admin for DAO operations like Where/Data.
type JinAdmin struct {
	g.Meta             `orm:"table:jin_admin, do:true"`
	Id                 interface{} //
	LoginName          interface{} // 登录名
	AdminName          interface{} // 员工名
	AdminImg           interface{} // 员工图片
	Password           interface{} // 密码
	Sex                interface{} // 性别 0未知 1男 2女
	IdCard             interface{} // 身份证号码
	Phone              interface{} // 手机号
	FixedLineTelephone interface{} // 固定电话
	AdminStatus        interface{} // 状态 0停用 1正常
	CreateTime         interface{} //
	UpdateTime         interface{} //
	AdminSn            interface{} //
	Ip                 interface{} // 最后登陆ip
	Location           interface{} // 最后登陆地址
}
