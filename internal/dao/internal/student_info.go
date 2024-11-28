// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// StudentInfoDao is the data access object for table student_info.
type StudentInfoDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns StudentInfoColumns // columns contains all the column names of Table for convenient usage.
}

// StudentInfoColumns defines and stores column names for table student_info.
type StudentInfoColumns struct {
	StudentId  string // 学号
	Name       string // 学生姓名
	Age        string // 年龄
	Class      string // 学生班级
	Department string // 所在系
	Phone      string // 学生电话
	CreateAt   string // 创建时间
	UpdateAt   string // 修改时间
	DeleteAt   string // 软删除
}

// studentInfoColumns holds the columns for table student_info.
var studentInfoColumns = StudentInfoColumns{
	StudentId:  "student_id",
	Name:       "name",
	Age:        "age",
	Class:      "class",
	Department: "department",
	Phone:      "phone",
	CreateAt:   "create_at",
	UpdateAt:   "update_at",
	DeleteAt:   "delete_at",
}

// NewStudentInfoDao creates and returns a new DAO object for table data access.
func NewStudentInfoDao() *StudentInfoDao {
	return &StudentInfoDao{
		group:   "default",
		table:   "student_info",
		columns: studentInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *StudentInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *StudentInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *StudentInfoDao) Columns() StudentInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *StudentInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *StudentInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *StudentInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
