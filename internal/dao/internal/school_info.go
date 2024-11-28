// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SchoolInfoDao is the data access object for table school_info.
type SchoolInfoDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns SchoolInfoColumns // columns contains all the column names of Table for convenient usage.
}

// SchoolInfoColumns defines and stores column names for table school_info.
type SchoolInfoColumns struct {
	SchoolId   string // 学校编号
	SchoolName string // 学校名称
	Province   string // 学校省份
	City       string // 学校城市
	County     string // 学号市区
	Address    string // 学校详细地址
}

// schoolInfoColumns holds the columns for table school_info.
var schoolInfoColumns = SchoolInfoColumns{
	SchoolId:   "school_id",
	SchoolName: "school_name",
	Province:   "province",
	City:       "city",
	County:     "county",
	Address:    "address",
}

// NewSchoolInfoDao creates and returns a new DAO object for table data access.
func NewSchoolInfoDao() *SchoolInfoDao {
	return &SchoolInfoDao{
		group:   "db2",
		table:   "school_info",
		columns: schoolInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SchoolInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SchoolInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SchoolInfoDao) Columns() SchoolInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SchoolInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SchoolInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SchoolInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
