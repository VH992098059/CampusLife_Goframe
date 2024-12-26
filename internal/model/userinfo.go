package model

type UserInfo struct {
	UserId   string `json:"user_id"`
	UserName string `json:"username"`
	NickName string `json:"nickname"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Wechat   string `json:"wechat"`
	Usersalt string `json:"user_salt"`
}
type UserRedis struct {
	UserId   string `json:"user_id"`
	UserName string `json:"username"`
	NickName string `json:"nickname"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Wechat   string `json:"wechat"`
}

type UserInfoModelInput struct {
	Page  int `json:"page" description:"分页码"`
	Size  int `json:"size" description:"分页数量"`
	Total int `json:"total" description:"数据总数"`
}
type UserInfoModelOutput struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}

type UserInfoModelCreateInput struct {
	UserId   string `json:"user_id"`
	UserName string `json:"username"`
	NickName string `json:"nickname"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Wechat   string `json:"wechat"`
	Usersalt string `json:"user_salt"`
}
type UserInfoModelCreateOutput struct {
	Userid uint `json:"user_id"`
}

type UserInfoModelUpdateInput struct {
	UserId   string `json:"user_id"`
	UserName string `json:"username"`
	NickName string `json:"nickname"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Wechat   string `json:"wechat"`

	Usersalt string `json:"user_salt"`
}
type UserInfoModelUpdateOutput struct {
	//Userid uint `json:"userid"`
}
