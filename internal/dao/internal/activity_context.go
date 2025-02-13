// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ActivityContextDao is the data access object for table activity_context.
type ActivityContextDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns ActivityContextColumns // columns contains all the column names of Table for convenient usage.
}

// ActivityContextColumns defines and stores column names for table activity_context.
type ActivityContextColumns struct {
	ActivityUuid string // 活动ID
	Pic          string // 图片内容
}

// activityContextColumns holds the columns for table activity_context.
var activityContextColumns = ActivityContextColumns{
	ActivityUuid: "activity_uuid",
	Pic:          "pic",
}

// NewActivityContextDao creates and returns a new DAO object for table data access.
func NewActivityContextDao() *ActivityContextDao {
	return &ActivityContextDao{
		group:   "default",
		table:   "activity_context",
		columns: activityContextColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ActivityContextDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ActivityContextDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ActivityContextDao) Columns() ActivityContextColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ActivityContextDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ActivityContextDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ActivityContextDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
