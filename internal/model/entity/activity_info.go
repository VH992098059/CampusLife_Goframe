// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ActivityInfo is the golang structure for table activity_info.
type ActivityInfo struct {
	Uuid                  int         `json:"uuid"                  orm:"uuid"                    description:"id"`
	ActivityPosters       string      `json:"activityPosters"       orm:"activity_posters"        description:"活动海报"`
	ActivityTitle         string      `json:"activityTitle"         orm:"activity_title"          description:"活动标题"`
	ActivityNumber        string      `json:"activityNumber"        orm:"activity_number"         description:"活动编号"`
	ActivityTypeId        int         `json:"activityTypeId"        orm:"activity_type_id"        description:"活动类型id"`
	Keywords              string      `json:"keywords"              orm:"keywords"                description:"关键词"`
	ReleaseTime           *gtime.Time `json:"releaseTime"           orm:"release_time"            description:"发布时间"`
	CheckStatus           int         `json:"checkStatus"           orm:"check_status"            description:"审核状态（0：草稿 1：待审核 2：报名中 3：审核失败 4：待举办 5：进行中 6：已结束）"`
	CheckTime             *gtime.Time `json:"checkTime"             orm:"check_time"              description:"审核时间"`
	CheckPersonContact    string      `json:"checkPersonContact"    orm:"check_person_contact"    description:"审核人联系方式"`
	CheckRemark           string      `json:"checkRemark"           orm:"check_remark"            description:"审核备注"`
	RegistrationStartTime *gtime.Time `json:"registrationStartTime" orm:"registration_start_time" description:"报名开始时间"`
	ActivityStartTime     *gtime.Time `json:"activityStartTime"     orm:"activity_start_time"     description:"活动开始时间"`
	Addr                  string      `json:"addr"                  orm:"addr"                    description:"地址"`
	PersonLimit           int         `json:"personLimit"           orm:"person_limit"            description:"人数限制"`
	RegistrationFee       float64     `json:"registrationFee"       orm:"registration_fee"        description:"报名费用"`
	WxCustomerCode        string      `json:"wxCustomerCode"        orm:"wx_customer_code"        description:"客服微信二维码"`
	PaymentMethod         string      `json:"paymentMethod"         orm:"payment_method"          description:"收款方式"`
	RegistrationEndTime   *gtime.Time `json:"registrationEndTime"   orm:"registration_end_time"   description:"报名结束时间"`
	ActivityEndTime       *gtime.Time `json:"activityEndTime"       orm:"activity_end_time"       description:"活动结束时间"`
	CheckNeed             int         `json:"checkNeed"             orm:"check_need"              description:"是否需要审核（1：需要 0：不需要）"`
	EnterActivistsId      int         `json:"enterActivistsId"      orm:"enter_activists_id"      description:"组织活动id"`
	CreateAt              *gtime.Time `json:"createAt"              orm:"create_at"               description:"创建时间"`
	UpdateAt              *gtime.Time `json:"updateAt"              orm:"update_at"               description:"修改时间"`
	DeleteAt              *gtime.Time `json:"deleteAt"              orm:"delete_at"               description:"软删除"`
}
