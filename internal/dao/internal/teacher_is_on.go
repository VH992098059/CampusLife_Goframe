// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TeacherIsOnDao is the data access object for table teacher_is_on.
type TeacherIsOnDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns TeacherIsOnColumns // columns contains all the column names of Table for convenient usage.
}

// TeacherIsOnColumns defines and stores column names for table teacher_is_on.
type TeacherIsOnColumns struct {
	TeacherId string // 职工编号
	UserId    string // 用户编号
	Name      string // 姓名
	IsOn      string // 是否在职
	SchoolId  string // 学校编号
	CreateAt  string // 创建时间
	UpdateAt  string // 修改时间
	DeleteAt  string // 软删除
}

// teacherIsOnColumns holds the columns for table teacher_is_on.
var teacherIsOnColumns = TeacherIsOnColumns{
	TeacherId: "teacher_id",
	UserId:    "user_id",
	Name:      "name",
	IsOn:      "is_on",
	SchoolId:  "school_id",
	CreateAt:  "create_at",
	UpdateAt:  "update_at",
	DeleteAt:  "delete_at",
}

// NewTeacherIsOnDao creates and returns a new DAO object for table data access.
func NewTeacherIsOnDao() *TeacherIsOnDao {
	return &TeacherIsOnDao{
		group:   "default",
		table:   "teacher_is_on",
		columns: teacherIsOnColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TeacherIsOnDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TeacherIsOnDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TeacherIsOnDao) Columns() TeacherIsOnColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TeacherIsOnDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TeacherIsOnDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TeacherIsOnDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
