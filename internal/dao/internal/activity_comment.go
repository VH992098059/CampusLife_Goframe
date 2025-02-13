// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ActivityCommentDao is the data access object for table activity_comment.
type ActivityCommentDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns ActivityCommentColumns // columns contains all the column names of Table for convenient usage.
}

// ActivityCommentColumns defines and stores column names for table activity_comment.
type ActivityCommentColumns struct {
	ActivityUuid string // 活动ID
	UserUuid     string // 用户ID
	Comment      string // 用户评论
	CreateAt     string // 发表时间
	DeleteAt     string // 删除时间
}

// activityCommentColumns holds the columns for table activity_comment.
var activityCommentColumns = ActivityCommentColumns{
	ActivityUuid: "activity_uuid",
	UserUuid:     "user_uuid",
	Comment:      "comment",
	CreateAt:     "create_at",
	DeleteAt:     "delete_at",
}

// NewActivityCommentDao creates and returns a new DAO object for table data access.
func NewActivityCommentDao() *ActivityCommentDao {
	return &ActivityCommentDao{
		group:   "default",
		table:   "activity_comment",
		columns: activityCommentColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ActivityCommentDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ActivityCommentDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ActivityCommentDao) Columns() ActivityCommentColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ActivityCommentDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ActivityCommentDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ActivityCommentDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
