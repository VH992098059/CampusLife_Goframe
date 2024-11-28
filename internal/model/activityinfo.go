package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type ActivityInfoModel struct {
	//UUID                  int       `json:"uuid"`                    // 学号
	ActivityPosters string `json:"activity_posters"` // 活动海报
	ActivityTitle   string `json:"activity_title"`   // 活动标题
	ActivityNumber  string `json:"activity_number"`  // 活动编号
	ActivityTypeID  int    `json:"activity_type_id"` // 活动类型id
	Keywords        string `json:"keywords"`         // 关键词
	//ReleaseTime           gtime.Time `json:"release_time"`            // 发布时间
	CheckStatus           int        `json:"check_status"`            // 审核状态（0：草稿 1：待审核 2：报名中 3：审核失败 4：待举办 5：进行中 6：已结束）
	CheckTime             gtime.Time `json:"check_time"`              // 审核时间
	CheckPersonContact    string     `json:"check_person_contact"`    // 审核人联系方式
	CheckRemark           string     `json:"check_remark"`            // 审核备注
	RegistrationStartTime gtime.Time `json:"registration_start_time"` // 报名开始时间
	ActivityStartTime     gtime.Time `json:"activity_start_time"`     // 活动开始时间
	Addr                  string     `json:"addr"`                    // 地址
	PersonLimit           int        `json:"person_limit"`            // 人数限制
	RegistrationFee       float64    `json:"registration_fee"`        // 报名费用
	WXCustomerCode        string     `json:"wx_customer_code"`        // 客服微信二维码
	PaymentMethod         string     `json:"payment_method"`          // 收款方式
	RegistrationEndTime   gtime.Time `json:"registration_end_time"`   // 报名结束时间
	ActivityEndTime       gtime.Time `json:"activity_end_time"`       // 活动结束时间
	CheckNeed             int        `json:"check_need"`              // 是否需要审核（1：需要 0：不需要）
	EnterActivistsID      int        `json:"enter_activists_id"`      // 组织活动id
}
type ActivityInfoModelListInput struct {
	Page  int `json:"page" description:"分页码"`
	Size  int `json:"size" description:"分页数量"`
	Total int `json:"total" description:"数据总数"`
}
type ActivityInfoModelListOutput struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
type ActivityInfoModelUpdateInput struct {
	ActivityPosters string `json:"activity_posters"` // 活动海报
	ActivityTitle   string `json:"activity_title"`   // 活动标题
	ActivityNumber  string `json:"activity_number"`  // 活动编号
	ActivityTypeID  int    `json:"activity_type_id"` // 活动类型id
	Keywords        string `json:"keywords"`         // 关键词
	//ReleaseTime           gtime.Time `json:"release_time"`            // 发布时间
	CheckStatus           int        `json:"check_status"`            // 审核状态（0：草稿 1：待审核 2：报名中 3：审核失败 4：待举办 5：进行中 6：已结束）
	CheckTime             gtime.Time `json:"check_time"`              // 审核时间
	CheckPersonContact    string     `json:"check_person_contact"`    // 审核人联系方式
	CheckRemark           string     `json:"check_remark"`            // 审核备注
	RegistrationStartTime gtime.Time `json:"registration_start_time"` // 报名开始时间
	ActivityStartTime     gtime.Time `json:"activity_start_time"`     // 活动开始时间
	Addr                  string     `json:"addr"`                    // 地址
	PersonLimit           int        `json:"person_limit"`            // 人数限制
	RegistrationFee       float64    `json:"registration_fee"`        // 报名费用
	WXCustomerCode        string     `json:"wx_customer_code"`        // 客服微信二维码
	PaymentMethod         string     `json:"payment_method"`          // 收款方式
	RegistrationEndTime   gtime.Time `json:"registration_end_time"`   // 报名结束时间
	ActivityEndTime       gtime.Time `json:"activity_end_time"`       // 活动结束时间
	CheckNeed             int        `json:"check_need"`              // 是否需要审核（1：需要 0：不需要）
	EnterActivistsID      int        `json:"enter_activists_id"`      // 组织活动id
}
type ActivityInfoModelUpdateOutput struct{}
type ActivityInfoModelDeleteInput struct {
	UUID int `json:"uuid"` // 学号
}
type ActivityInfoModelDeleteOutput struct{}
type ActivityInfoModelAddInput struct {
	ActivityPosters       string     `json:"activity_posters"`        // 活动海报
	ActivityTitle         string     `json:"activity_title"`          // 活动标题
	ActivityNumber        string     `json:"activity_number"`         // 活动编号
	ActivityTypeID        int        `json:"activity_type_id"`        // 活动类型id
	Keywords              string     `json:"keywords"`                // 关键词
	ReleaseTime           gtime.Time `json:"release_time"`            // 发布时间
	CheckStatus           int        `json:"check_status"`            // 审核状态（0：草稿 1：待审核 2：报名中 3：审核失败 4：待举办 5：进行中 6：已结束）
	CheckTime             gtime.Time `json:"check_time"`              // 审核时间
	CheckPersonContact    string     `json:"check_person_contact"`    // 审核人联系方式
	CheckRemark           string     `json:"check_remark"`            // 审核备注
	RegistrationStartTime gtime.Time `json:"registration_start_time"` // 报名开始时间
	ActivityStartTime     gtime.Time `json:"activity_start_time"`     // 活动开始时间
	Addr                  string     `json:"addr"`                    // 地址
	PersonLimit           int        `json:"person_limit"`            // 人数限制
	RegistrationFee       float64    `json:"registration_fee"`        // 报名费用
	WXCustomerCode        string     `json:"wx_customer_code"`        // 客服微信二维码
	PaymentMethod         string     `json:"payment_method"`          // 收款方式
	RegistrationEndTime   gtime.Time `json:"registration_end_time"`   // 报名结束时间
	ActivityEndTime       gtime.Time `json:"activity_end_time"`       // 活动结束时间
	CheckNeed             int        `json:"check_need"`              // 是否需要审核（1：需要 0：不需要）
	EnterActivistsID      int        `json:"enter_activists_id"`      // 组织活动id
}
type ActivityInfoModelAddOutput struct{}
