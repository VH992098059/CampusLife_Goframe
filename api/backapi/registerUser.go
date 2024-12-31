package backapi

import "github.com/gogf/gf/v2/frame/g"

type RegisterUserInfo struct {
	UserId   string `json:"user_id" dc:"用户ID"`
	UserName string `json:"username" v:"required#用户名不能为空" dc:"用户名"`
	NickName string `json:"nickname" v:"required#昵称不能为空" dc:"用户昵称"`
	Password string `json:"password" v:"required#密码不能为空" dc:"用户密码"`
	Phone    string `json:"phone" v:"required|phone#手机号不能为空|手机格式不正确" dc:"用户手机号"`
	Type     string `json:"type" v:"required#用户类型不能为空" dc:"用户类型"`
	Usersalt string `json:"user_salt" dc:"加密盐"`
	Code     string `json:"verification" dc:"验证码"`
}
type RegisterUserReq struct {
	g.Meta `path:"/register" method:"post"`
	RegisterUserInfo
}
type RegisterUserRes struct {
	UserId string `json:"user_id" `
}
