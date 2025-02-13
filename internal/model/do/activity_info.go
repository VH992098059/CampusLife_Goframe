// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ActivityInfo is the golang structure of table activity_info for DAO operations like Where/Data.
type ActivityInfo struct {
	g.Meta                `orm:"table:activity_info, do:true"`
	Uuid                  interface{} // id
	ActivityPosters       interface{} // 活动海报
	ActivityTitle         interface{} // 活动标题
	ActivityTypeId        interface{} // 活动类型id
	ActivityAbout         interface{} // 活动介绍
	Keywords              interface{} // 关键词
	ReleaseTime           *gtime.Time // 发布时间
	CheckStatus           interface{} // 审核状态（0：草稿 1：待审核 2：报名中 3：审核失败 4：待举办 5：进行中 6：已结束）
	CheckTime             *gtime.Time // 审核时间
	CheckPersonContact    interface{} // 审核人联系方式
	CheckRemark           interface{} // 审核备注
	RegistrationStartTime *gtime.Time // 报名开始时间
	ActivityStartTime     *gtime.Time // 活动开始时间
	Addr                  interface{} // 地址
	PersonCurrent         interface{} // 当前报名人数
	PersonLimit           interface{} // 人数限制
	RegistrationFee       interface{} // 报名费用
	WxCustomerCode        interface{} // 客服微信二维码
	PaymentMethod         interface{} // 收款方式
	RegistrationEndTime   *gtime.Time // 报名结束时间
	ActivityEndTime       *gtime.Time // 活动结束时间
	CheckNeed             interface{} // 是否需要审核（1：需要 0：不需要）
	EnterActivistsId      interface{} // 组织活动id
	Popular               interface{} // 活动热门度
	CreateAt              *gtime.Time // 创建时间
	UpdateAt              *gtime.Time // 修改时间
	DeleteAt              *gtime.Time // 软删除
}
