package backapi

import "github.com/gogf/gf/v2/frame/g"

type RegisterUserInfo struct {
	UserId   string `json:"user_id" `
	UserName string `json:"username" v:"required#用户名不能为空"`
	NickName string `json:"nickname" v:"required#昵称不能为空"`
	Password string `json:"password" v:"required#密码不能为空"`
	Phone    string `json:"phone" v:"required|phone#用户名不能为空|电话格式不正确"`
	Type     string `json:"type" v:"required#用户类型不能为空"`
	Usersalt string `json:"user_salt"`
}
type RegisterUserReq struct {
	g.Meta `path:"/register" method:"post"`
	RegisterUserInfo
}
type RegisterUserRes struct {
	UserId string `json:"user_id" `
}
