// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserStateDao is the data access object for table user_state.
type UserStateDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns UserStateColumns // columns contains all the column names of Table for convenient usage.
}

// UserStateColumns defines and stores column names for table user_state.
type UserStateColumns struct {
	UserId   string // 用户ID
	Nickname string // 用户昵称
	Username string // 用户账号
	Type     string // 用户类型（活动方，普通用户）
	Avater   string // 用户头像
	Phone    string // 用户电话
	Wechat   string // 用户微信
	Birthday string // 用户生日
	Email    string // 邮箱
	Sex      string // 用户性别
}

// userStateColumns holds the columns for table user_state.
var userStateColumns = UserStateColumns{
	UserId:   "user_id",
	Nickname: "nickname",
	Username: "username",
	Type:     "type",
	Avater:   "avater",
	Phone:    "phone",
	Wechat:   "wechat",
	Birthday: "birthday",
	Email:    "email",
	Sex:      "sex",
}

// NewUserStateDao creates and returns a new DAO object for table data access.
func NewUserStateDao() *UserStateDao {
	return &UserStateDao{
		group:   "default",
		table:   "user_state",
		columns: userStateColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserStateDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserStateDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserStateDao) Columns() UserStateColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserStateDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserStateDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserStateDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
