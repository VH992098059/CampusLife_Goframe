// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserDormitoryDao is the data access object for table user_dormitory.
type UserDormitoryDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns UserDormitoryColumns // columns contains all the column names of Table for convenient usage.
}

// UserDormitoryColumns defines and stores column names for table user_dormitory.
type UserDormitoryColumns struct {
	UserId   string //
	BelongId string //
	School   string //
}

// userDormitoryColumns holds the columns for table user_dormitory.
var userDormitoryColumns = UserDormitoryColumns{
	UserId:   "user_id",
	BelongId: "belong_id",
	School:   "school",
}

// NewUserDormitoryDao creates and returns a new DAO object for table data access.
func NewUserDormitoryDao() *UserDormitoryDao {
	return &UserDormitoryDao{
		group:   "default",
		table:   "user_dormitory",
		columns: userDormitoryColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserDormitoryDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserDormitoryDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserDormitoryDao) Columns() UserDormitoryColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserDormitoryDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserDormitoryDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserDormitoryDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
