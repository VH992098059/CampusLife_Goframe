// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// StudentIsOnDao is the data access object for table student_is_on.
type StudentIsOnDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns StudentIsOnColumns // columns contains all the column names of Table for convenient usage.
}

// StudentIsOnColumns defines and stores column names for table student_is_on.
type StudentIsOnColumns struct {
	StudentId string // 学号
	UserId    string // 用户编号
	Name      string // 姓名
	IsOn      string // 是否在校
	SchoolId  string // 学校编号
	CreateAt  string // 创建时间
	UpdateAt  string // 修改时间
	DeleteAt  string // 软删除
}

// studentIsOnColumns holds the columns for table student_is_on.
var studentIsOnColumns = StudentIsOnColumns{
	StudentId: "student_id",
	UserId:    "user_id",
	Name:      "name",
	IsOn:      "is_on",
	SchoolId:  "school_id",
	CreateAt:  "create_at",
	UpdateAt:  "update_at",
	DeleteAt:  "delete_at",
}

// NewStudentIsOnDao creates and returns a new DAO object for table data access.
func NewStudentIsOnDao() *StudentIsOnDao {
	return &StudentIsOnDao{
		group:   "default",
		table:   "student_is_on",
		columns: studentIsOnColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *StudentIsOnDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *StudentIsOnDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *StudentIsOnDao) Columns() StudentIsOnColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *StudentIsOnDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *StudentIsOnDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *StudentIsOnDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
