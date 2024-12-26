package backapi

import (
	"github.com/gogf/gf/v2/frame/g"
)

type UserInfo struct {
	UserId   string `json:"user_id" v:"required|length:14,14#用户编号不能为空|用户ID必须在{max}位"`
	UserName string `json:"username" v:"required#用户名不能为空"`
	NickName string `json:"nickname" v:"required#昵称不能为空"`
	Password string `json:"password" v:"required#密码不能为空"`
	Phone    string `json:"phone" v:"required|phone#用户名不能为空|电话格式不正确"`
	Wechat   string `json:"wechat" `
	Usersalt string `json:"user_salt"`
}
type UserLoginInfo struct {
	Id          int    `json:"id"`
	IdentityKey string `json:"identity_key"`
	Payload     string `json:"payload"`
}

// UserInfoListReq 查询用户
type UserInfoListReq struct {
	g.Meta `path:"/userInfoList" method:"get"`
	CommonPaginationReq
}
type UserInfoListRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码" `
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}

type UserInfoCreateReq struct {
	g.Meta `path:"/userInfoCreateList" method:"post"`
	UserInfo
}
type UserInfoCreateRes struct {
	Userid uint `json:"user_id"`
}

type UserInfoUpdateReq struct {
	g.Meta `path:"/userInfoUpdateList" method:"put"`
	UserInfo
}
type UserInfoUpdateRes struct {
	//Userid uint `json:"userid"`
}
type UserInfoDeleteReq struct {
	g.Meta `path:"/userInfoDeleteList/{Userid}" method:"delete"`
	Userid string `json:"user_id"`
}
type UserInfoDeleteRes struct{}

type UserGetInfoReq struct {
	g.Meta  `path:"/info" method:"get"`
	UserKey string `json:"user_key"`
}

type UserGetInfoRes struct {
	//JWT验证
	/*Id          int    `json:"id"`
	IdentityKey string `json:"identity_key"`
	Payload     string `json:"payload"`*/
	UserId   string      `json:"id"`
	Username interface{} `json:"username"`
}
