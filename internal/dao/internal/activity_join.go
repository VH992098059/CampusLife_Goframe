// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ActivityJoinDao is the data access object for table activity_join.
type ActivityJoinDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns ActivityJoinColumns // columns contains all the column names of Table for convenient usage.
}

// ActivityJoinColumns defines and stores column names for table activity_join.
type ActivityJoinColumns struct {
	JoinUuid     string // 参加活动id
	ActivityUuid string // 活动id
	UserId       string // 用户id
	CreateAt     string // 创建时间
	UpdateAt     string // 修改时间
	DeleteAt     string // 删除时间
}

// activityJoinColumns holds the columns for table activity_join.
var activityJoinColumns = ActivityJoinColumns{
	JoinUuid:     "join_uuid",
	ActivityUuid: "activity_uuid",
	UserId:       "user_id",
	CreateAt:     "create_at",
	UpdateAt:     "update_at",
	DeleteAt:     "delete_at",
}

// NewActivityJoinDao creates and returns a new DAO object for table data access.
func NewActivityJoinDao() *ActivityJoinDao {
	return &ActivityJoinDao{
		group:   "default",
		table:   "activity_join",
		columns: activityJoinColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ActivityJoinDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ActivityJoinDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ActivityJoinDao) Columns() ActivityJoinColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ActivityJoinDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ActivityJoinDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ActivityJoinDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
