// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ConfigDao is the data access object for table jin_config.
type ConfigDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns ConfigColumns // columns contains all the column names of Table for convenient usage.
}

// ConfigColumns defines and stores column names for table jin_config.
type ConfigColumns struct {
	Id         string //
	Name       string // 配置名
	Value      string // 配置内容
	Type       string // 配置类型 0后台 1会员小程序
	Remark     string //
	CreateTime string //
	UpdateTime string //
}

//  configColumns holds the columns for table jin_config.
var configColumns = ConfigColumns{
	Id:         "id",
	Name:       "name",
	Value:      "value",
	Type:       "type",
	Remark:     "remark",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewConfigDao creates and returns a new DAO object for table data access.
func NewConfigDao() *ConfigDao {
	return &ConfigDao{
		group:   "default",
		table:   "jin_config",
		columns: configColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ConfigDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ConfigDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ConfigDao) Columns() ConfigColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ConfigDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ConfigDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ConfigDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}