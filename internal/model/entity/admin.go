// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Admin is the golang structure for table admin.
type Admin struct {
	Id                 uint64 `json:"id"                 ` //
	LoginName          string `json:"loginName"          ` // 登录名
	AdminName          string `json:"adminName"          ` // 员工名
	AdminImg           string `json:"adminImg"           ` // 员工图片
	Password           string `json:"password"           ` // 密码
	Sex                int    `json:"sex"                ` // 性别 0未知 1男 2女
	IdCard             string `json:"idCard"             ` // 身份证号码
	Phone              string `json:"phone"              ` // 手机号
	FixedLineTelephone string `json:"fixedLineTelephone" ` // 固定电话
	AdminStatus        int    `json:"adminStatus"        ` // 状态 0停用 1正常
	CreateTime         int    `json:"createTime"         ` //
	UpdateTime         int    `json:"updateTime"         ` //
	AdminSn            string `json:"adminSn"            ` //
	Ip                 string `json:"ip"                 ` // 最后登陆ip
	Location           string `json:"location"           ` // 最后登陆地址
}
