// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ActivityInfoDao is the data access object for table activity_info.
type ActivityInfoDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns ActivityInfoColumns // columns contains all the column names of Table for convenient usage.
}

// ActivityInfoColumns defines and stores column names for table activity_info.
type ActivityInfoColumns struct {
	Uuid                  string // id
	ActivityPosters       string // 活动海报
	ActivityTitle         string // 活动标题
	ActivityTypeId        string // 活动类型id
	ActivityAbout         string // 活动介绍
	Keywords              string // 关键词
	ReleaseTime           string // 发布时间
	CheckStatus           string // 审核状态（0：草稿 1：待审核 2：报名中 3：审核失败 4：待举办 5：进行中 6：已结束）
	CheckTime             string // 审核时间
	CheckPersonContact    string // 审核人联系方式
	CheckRemark           string // 审核备注
	RegistrationStartTime string // 报名开始时间
	ActivityStartTime     string // 活动开始时间
	Addr                  string // 地址
	PersonCurrent         string // 当前报名人数
	PersonLimit           string // 人数限制
	RegistrationFee       string // 报名费用
	WxCustomerCode        string // 客服微信二维码
	PaymentMethod         string // 收款方式
	RegistrationEndTime   string // 报名结束时间
	ActivityEndTime       string // 活动结束时间
	CheckNeed             string // 是否需要审核（1：需要 0：不需要）
	EnterActivistsId      string // 组织活动id
	Popular               string // 活动热门度
	CreateAt              string // 创建时间
	UpdateAt              string // 修改时间
	DeleteAt              string // 软删除
}

// activityInfoColumns holds the columns for table activity_info.
var activityInfoColumns = ActivityInfoColumns{
	Uuid:                  "uuid",
	ActivityPosters:       "activity_posters",
	ActivityTitle:         "activity_title",
	ActivityTypeId:        "activity_type_id",
	ActivityAbout:         "activity_about",
	Keywords:              "keywords",
	ReleaseTime:           "release_time",
	CheckStatus:           "check_status",
	CheckTime:             "check_time",
	CheckPersonContact:    "check_person_contact",
	CheckRemark:           "check_remark",
	RegistrationStartTime: "registration_start_time",
	ActivityStartTime:     "activity_start_time",
	Addr:                  "addr",
	PersonCurrent:         "person_current",
	PersonLimit:           "person_limit",
	RegistrationFee:       "registration_fee",
	WxCustomerCode:        "wx_customer_code",
	PaymentMethod:         "payment_method",
	RegistrationEndTime:   "registration_end_time",
	ActivityEndTime:       "activity_end_time",
	CheckNeed:             "check_need",
	EnterActivistsId:      "enter_activists_id",
	Popular:               "popular",
	CreateAt:              "create_at",
	UpdateAt:              "update_at",
	DeleteAt:              "delete_at",
}

// NewActivityInfoDao creates and returns a new DAO object for table data access.
func NewActivityInfoDao() *ActivityInfoDao {
	return &ActivityInfoDao{
		group:   "default",
		table:   "activity_info",
		columns: activityInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ActivityInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ActivityInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ActivityInfoDao) Columns() ActivityInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ActivityInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ActivityInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ActivityInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
