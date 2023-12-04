package loginServe

import (
	userInfo "ChatServe/src/model/user"
	"ChatServe/src/utils"
)

func LoginInto(username string, password string) (user userInfo.UserInfo) {
	utils.DB.Where("phone = ? ", username).Find(&user)
	return user
}
