// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ActivityViewContextDao is the data access object for table activity_view_context.
type ActivityViewContextDao struct {
	table   string                     // table is the underlying table name of the DAO.
	group   string                     // group is the database configuration group name of current DAO.
	columns ActivityViewContextColumns // columns contains all the column names of Table for convenient usage.
}

// ActivityViewContextColumns defines and stores column names for table activity_view_context.
type ActivityViewContextColumns struct {
	Uuid                  string // id
	ActivityTitle         string // 活动标题
	ActivityAbout         string // 活动介绍
	Popular               string // 活动热门度
	RegistrationStartTime string // 报名开始时间
	RegistrationEndTime   string // 报名结束时间
	ActivityStartTime     string // 活动开始时间
	ActivityEndTime       string // 活动结束时间
	PersonLimit           string // 人数限制
	PersonCurrent         string // 当前报名人数
}

// activityViewContextColumns holds the columns for table activity_view_context.
var activityViewContextColumns = ActivityViewContextColumns{
	Uuid:                  "uuid",
	ActivityTitle:         "activity_title",
	ActivityAbout:         "activity_about",
	Popular:               "popular",
	RegistrationStartTime: "registration_start_time",
	RegistrationEndTime:   "registration_end_time",
	ActivityStartTime:     "activity_start_time",
	ActivityEndTime:       "activity_end_time",
	PersonLimit:           "person_limit",
	PersonCurrent:         "person_current",
}

// NewActivityViewContextDao creates and returns a new DAO object for table data access.
func NewActivityViewContextDao() *ActivityViewContextDao {
	return &ActivityViewContextDao{
		group:   "default",
		table:   "activity_view_context",
		columns: activityViewContextColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ActivityViewContextDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ActivityViewContextDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ActivityViewContextDao) Columns() ActivityViewContextColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ActivityViewContextDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ActivityViewContextDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ActivityViewContextDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
