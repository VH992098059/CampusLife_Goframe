package cmd

import (
	"context"
	"demo3/api/backapi"
	"demo3/internal/consts"
	"demo3/internal/controller"
	"demo3/internal/dao"
	"demo3/internal/model"
	"demo3/utility"
	"demo3/utility/aliyunCode"
	"demo3/utility/gtoken"
	"demo3/utility/middleware"
	"demo3/utility/response"
	"demo3/utility/validatePhone"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/text/gstr"
)

var UserTimeout = 24 * 60 * 60 * 1000 //毫秒单位，为1天时间
var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			loginFunc := Login       //登录操作
			loginAfter := LoginAfter //登录过程操作
			gfToken := gtoken.GfToken{
				Timeout:         UserTimeout,
				ServerName:      "demo",
				LoginPath:       "/login",
				LoginBeforeFunc: loginFunc,
				LoginAfterFunc:  loginAfter,
				LogoutPath:      "/logout",
				MultiLogin:      true,                  //多平台登录
				CacheMode:       gtoken.CacheModeRedis, //使用redis存储
			}
			//是否允许跨域操作
			s.Use(func(r *ghttp.Request) {
				r.Response.CORSDefault()
				r.Middleware.Next()
			})
			s.Group("/backend", func(group *ghttp.RouterGroup) {

				//获取验证码
				s.BindHandler("/sendCode", func(r *ghttp.Request) {
					phone := r.Get("phone").String()
					/*手机号正则表达*/
					err := validatePhone.ValidatePhone(phone)
					if err != true {
						r.Response.WriteJson(g.Map{
							"code": 500,
							"msg":  "手机号码不正确",
						})
						g.Log().Error(ctx, "手机号码不正确")
						return
					}
					/*阿里云验证码*/
					ValidMainErr := aliyunCode.ValidMain(phone, "SMS_476855598")
					if ValidMainErr != nil {
						r.Response.WriteJson(g.Map{
							"code": 500,
							"msg":  "验证码发送失败",
						})
						return
					}
					r.Response.WriteJson(g.Map{
						"code": 200,
						"msg":  "验证码发送成功",
					})
					return
				})
				/*中间件*/
				group.Middleware(
					middleware.MiddlewareHandlerResponse,
				)
				group.Bind(

					controller.RegisterUserInfo,
					controller.ActivityInfo,
				)
				//验证token令牌 JWT
				/*group.Bind(
					controller.LoginInfo,
					controller.LogoutInfo,

				)*/
				/*group.Middleware(service.Middleware().Auth)*/

				/*gtoken*/
				err := gfToken.Middleware(ctx, group)
				if err != nil {
					panic(err)
				}
				group.Bind(
					controller.Userinfo,
					controller.JoinActivity,
				)

			})
			s.Run()
			return nil
		},
	}
)

// Login 登录操作 （Gtoken）
func Login(r *ghttp.Request) (string, interface{}) {
	username := r.Get("username").String()
	password := r.Get("password").String()
	ctx := context.TODO()
	if username == "" || password == "" {
		r.Response.WriteJson(gtoken.Fail("账号或密码错误"))
		r.ExitAll()
	}
	UserInfo := model.UserInfo{}
	err := dao.UserInfo.Ctx(ctx).Where("username", username).Scan(&UserInfo)
	if err != nil {
		r.Response.WriteJson(gtoken.Fail("账号不存在"))
		r.ExitAll()
	}
	if utility.EncryptPassword(password, UserInfo.Usersalt) != UserInfo.Password {
		r.Response.WriteJson(gtoken.Fail("账号或密码错误"))
		r.ExitAll()
	}
	UserRedis := model.UserRedis{
		UserId:   UserInfo.UserId,
		UserName: UserInfo.UserName,
		NickName: UserInfo.NickName,
		Avatar:   UserInfo.Avatar,
	}

	/*生成token*/
	return consts.GTokenUserPrefix + UserInfo.UserId, UserRedis
}

// LoginAfter 登录过程操作
func LoginAfter(r *ghttp.Request, respData gtoken.Resp) {
	if !respData.Success() {
		respData.Code = 500
		r.Response.WriteJson(respData)
		r.ExitAll()
	} else {
		respData.Msg = "登录成功"
		userKey := respData.GetString("userKey")
		userId := gstr.StrEx(userKey, consts.GTokenUserPrefix)
		userInfo := model.UserRedis{}
		err := dao.UserInfo.Ctx(context.TODO()).Where("user_id", userId).Scan(&userInfo)
		if err != nil {
			return
		}
		response.SuccessWithData(r, &backapi.LoginRes{
			Uuid:     userId,
			Type:     "redis",
			Nickname: userInfo.NickName,
			Token:    respData.GetString("token"),
			ExpireIn: UserTimeout / 1000,
		})

	}
}
