package login

import (
	loginServe "ChatServe/src/model/login"
	"ChatServe/src/utils"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var loginParams loginServe.Login
	err := c.Bind(&loginParams)
	if err != nil {
		c.JSON(utils.SuccessCode, utils.BackMessageResp(utils.FailedCode, "请填写完整"))
		panic(err)
	}
	user := loginServe.LoginInto(loginParams.UserName, loginParams.PassWord)
	if user.Phone == " " {
		c.JSON(utils.SuccessCode, utils.BackMessageResp(utils.FailedCode, "暂无该用户!"))
		return
	}
	if user.PassWord == loginParams.PassWord {
		token, errs := utils.GetToken(user.Id)
		if errs != nil {
			panic(errs)
		}
		c.JSON(utils.SuccessCode, utils.BackDataResp(utils.SuccessCode, "登陆成功", gin.H{
			"token": token,
		}))
		return
	} else {
		c.JSON(utils.SuccessCode, utils.BackMessageResp(utils.FailedCode, "账号和密码不正确！"))
	}

}
