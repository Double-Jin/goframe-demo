// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AdminDao is the data access object for table jin_admin.
type AdminDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of current DAO.
	columns AdminColumns // columns contains all the column names of Table for convenient usage.
}

// AdminColumns defines and stores column names for table jin_admin.
type AdminColumns struct {
	Id                 string //
	LoginName          string // 登录名
	AdminName          string // 员工名
	AdminImg           string // 员工图片
	Password           string // 密码
	Sex                string // 性别 0未知 1男 2女
	IdCard             string // 身份证号码
	Phone              string // 手机号
	FixedLineTelephone string // 固定电话
	AdminStatus        string // 状态 0停用 1正常
	CreateTime         string //
	UpdateTime         string //
	AdminSn            string //
	Ip                 string // 最后登陆ip
	Location           string // 最后登陆地址
}

//  adminColumns holds the columns for table jin_admin.
var adminColumns = AdminColumns{
	Id:                 "id",
	LoginName:          "login_name",
	AdminName:          "admin_name",
	AdminImg:           "admin_img",
	Password:           "password",
	Sex:                "sex",
	IdCard:             "id_card",
	Phone:              "phone",
	FixedLineTelephone: "fixed_line_telephone",
	AdminStatus:        "admin_status",
	CreateTime:         "create_time",
	UpdateTime:         "update_time",
	AdminSn:            "admin_sn",
	Ip:                 "ip",
	Location:           "location",
}

// NewAdminDao creates and returns a new DAO object for table data access.
func NewAdminDao() *AdminDao {
	return &AdminDao{
		group:   "default",
		table:   "jin_admin",
		columns: adminColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AdminDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AdminDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AdminDao) Columns() AdminColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AdminDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AdminDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AdminDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
