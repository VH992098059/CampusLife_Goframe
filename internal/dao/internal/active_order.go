// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ActiveOrderDao is the data access object for table active_order.
type ActiveOrderDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns ActiveOrderColumns // columns contains all the column names of Table for convenient usage.
}

// ActiveOrderColumns defines and stores column names for table active_order.
type ActiveOrderColumns struct {
	Id           string // id
	Status       string // 状态（1：待审核 2：待参加 3：已检票 4：已取消）
	ActivityId   string // 活动id
	CommonUserId string // 普通用户id
	CreateAt     string // 创建时间
	UpdateAt     string // 修改时间
	DeleteAt     string // 软删除
}

// activeOrderColumns holds the columns for table active_order.
var activeOrderColumns = ActiveOrderColumns{
	Id:           "id",
	Status:       "status",
	ActivityId:   "activity_id",
	CommonUserId: "common_user_id",
	CreateAt:     "create_at",
	UpdateAt:     "update_at",
	DeleteAt:     "delete_at",
}

// NewActiveOrderDao creates and returns a new DAO object for table data access.
func NewActiveOrderDao() *ActiveOrderDao {
	return &ActiveOrderDao{
		group:   "default",
		table:   "active_order",
		columns: activeOrderColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ActiveOrderDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ActiveOrderDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ActiveOrderDao) Columns() ActiveOrderColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ActiveOrderDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ActiveOrderDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ActiveOrderDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
