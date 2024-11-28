package backapi

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type ActivityInfo struct {
	//Uuid          string     `json:"uuid" description:"活动ID" v:"required#活动ID不能为空"`
	//UUID                  int       `json:"uuid" v:"required#UUID不能为空" description:"学号"`
	ActivityPosters string `json:"activity_posters" v:"required#活动海报不能为空" description:"活动海报"`
	ActivityTitle   string `json:"activity_title" v:"required#活动标题不能为空" description:"活动标题"`
	//ActivityNumber        string     `json:"activity_number" v:"required#活动编号不能为空" description:"活动编号"`
	ActivityTypeID        int        `json:"activity_type_id" v:"required#活动类型ID不能为空" description:"活动类型id"`
	Keywords              string     `json:"keywords" v:"required#关键词不能为空" description:"关键词"`
	ReleaseTime           gtime.Time `json:"release_time" v:"required#发布时间不能为空" description:"发布时间"`
	CheckStatus           int        `json:"check_status" v:"required#审核状态不能为空" description:"审核状态（0：草稿 1：待审核 2：报名中 3：审核失败 4：待举办 5：进行中 6：已结束）"`
	CheckTime             gtime.Time `json:"check_time" v:"required#审核时间不能为空" description:"审核时间"`
	CheckPersonContact    string     `json:"check_person_contact" v:"required#审核人联系方式不能为空" description:"审核人联系方式"`
	CheckRemark           string     `json:"check_remark" v:"required#审核备注不能为空" description:"审核备注"`
	RegistrationStartTime gtime.Time `json:"registration_start_time" v:"required#报名开始时间不能为空" description:"报名开始时间"`
	ActivityStartTime     gtime.Time `json:"activity_start_time" v:"required#活动开始时间不能为空" description:"活动开始时间"`
	Addr                  string     `json:"addr" v:"required#地址不能为空" description:"地址"`
	PersonLimit           int        `json:"person_limit" v:"required#人数限制不能为空" description:"人数限制"`
	RegistrationFee       float64    `json:"registration_fee" v:"required#报名费用不能为空" description:"报名费用"`
	WXCustomerCode        string     `json:"wx_customer_code" v:"required#客服微信二维码不能为空" description:"客服微信二维码"`
	PaymentMethod         string     `json:"payment_method" v:"required#收款方式不能为空" description:"收款方式"`
	RegistrationEndTime   gtime.Time `json:"registration_end_time" v:"required#报名结束时间不能为空" description:"报名结束时间"`
	ActivityEndTime       gtime.Time `json:"activity_end_time" v:"required#活动结束时间不能为空" description:"活动结束时间"`
	CheckNeed             int        `json:"check_need" v:"required#是否需要审核不能为空" description:"是否需要审核（1：需要 0：不需要）"`
	EnterActivistsID      int        `json:"enter_activists_id" v:"required#组织活动ID不能为空" description:"组织活动id"`
}
type ActivityInfoListReq struct {
	g.Meta `path:"/activityList" method:"get"`
	CommonPaginationRes
}
type ActivityInfoListRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码" `
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
type ActivityInfoUpdateReq struct {
	g.Meta `path:"/activityUpdate" method:"update"`
	ActivityInfo
}
type ActivityInfoUpdateRes struct{}
type ActivityInfoDeleteReq struct {
	g.Meta `path:"/activityDelete" method:"delete"`
	Uuid   string `json:"uuid" description:"活动ID" v:"required#活动ID不能为空"`
}
type ActivityInfoDeleteRes struct{}
type ActivityInfoAddReq struct {
	g.Meta `path:"/activityAdd" method:"post"`
	ActivityInfo
}
type ActivityInfoAddRes struct{}
