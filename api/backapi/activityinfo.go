package backapi

import "github.com/gogf/gf/v2/frame/g"

type ActivityInfo struct {
	Uuid            string `json:"uuid"`             // id
	ActivityPosters string `json:"activity_posters"` // 活动海报
	ActivityTitle   string `json:"activity_title"`   // 活动标题
	ActivityTypeId  string `json:"activity_type_id"` // 活动类型id
	Keywords        string `json:"keywords"`         // 关键词
	ReleaseTime     string // 发布时间
	/*	CheckStatus           string // 审核状态（0：草稿 1：待审核 2：报名中 3：审核失败 4：待举办 5：进行中 6：已结束）
		CheckTime             string // 审核时间
		CheckPersonContact    string // 审核人联系方式
		CheckRemark           string // 审核备注*/
	RegistrationStartTime string // 报名开始时间
	ActivityStartTime     string // 活动开始时间
	Addr                  string // 地址
	PersonCurrent         string `json:"person_current"` //当前人数
	PersonLimit           string `json:"person_limit"`   // 人数限制
	/*	RegistrationFee       string // 报名费用
		WxCustomerCode        string // 客服微信二维码
		PaymentMethod         string // 收款方式*/
	RegistrationEndTime string // 报名结束时间
	ActivityEndTime     string // 活动结束时间
	CheckNeed           string // 是否需要审核（1：需要 0：不需要）
	EnterActivistsId    string // 组织活动id
	Popular             int64  `json:"popular"` //活动热门度
}

/*获取活动列表*/
type ActivityInfoListReq struct {
	g.Meta `path:"/activityList" method:"get"`
	Page   int `json:"page" description:"分页码" `
	Size   int `json:"size" description:"分页数量"`
}
type ActivityInfoListRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码" `
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}

/*获取活动内容*/
type ActivityInfoGetReq struct {
	g.Meta `path:"/GetActivity" method:"get"`
	Uuid   string `json:"uuid"`
}
type ActivityInfoGetRes struct {
	List interface{} `json:"list" description:"列表"`
}

/*搜索活动*/
type ActivitySearchReq struct {
	g.Meta        `path:"/activitySearch" method:"get"`
	ActivityTitle string `json:"activity_title"` // 活动标题
	Page          int    `json:"page" description:"分页码" `
	Size          int    `json:"size" description:"分页数量"`
}
type ActivitySearchRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码" `
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
