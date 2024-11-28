// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// EmployeeIsOnDao is the data access object for table employee_is_on.
type EmployeeIsOnDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns EmployeeIsOnColumns // columns contains all the column names of Table for convenient usage.
}

// EmployeeIsOnColumns defines and stores column names for table employee_is_on.
type EmployeeIsOnColumns struct {
	EmployeeId string // 职工编号
	UserId     string // 用户编号
	Name       string // 姓名
	IsOn       string // 是否在职
	SchoolId   string // 学校编号
	CreateAt   string // 创建时间
	UpdateAt   string // 修改时间
	DeleteAt   string // 软删除
}

// employeeIsOnColumns holds the columns for table employee_is_on.
var employeeIsOnColumns = EmployeeIsOnColumns{
	EmployeeId: "employee_id",
	UserId:     "user_id",
	Name:       "name",
	IsOn:       "is_on",
	SchoolId:   "school_id",
	CreateAt:   "create_at",
	UpdateAt:   "update_at",
	DeleteAt:   "delete_at",
}

// NewEmployeeIsOnDao creates and returns a new DAO object for table data access.
func NewEmployeeIsOnDao() *EmployeeIsOnDao {
	return &EmployeeIsOnDao{
		group:   "default",
		table:   "employee_is_on",
		columns: employeeIsOnColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *EmployeeIsOnDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *EmployeeIsOnDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *EmployeeIsOnDao) Columns() EmployeeIsOnColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *EmployeeIsOnDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *EmployeeIsOnDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *EmployeeIsOnDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
