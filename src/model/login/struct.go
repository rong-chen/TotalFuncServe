package loginServe

type Login struct {
	//手机号码
	UserName string `json:"username" form:"username" binding:"required"`
	//密码
	PassWord string `json:"password" form:"password" binding:"required"`
}
