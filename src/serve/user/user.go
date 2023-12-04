package user

import (
	"ChatServe/src/model/friends"
	userInfo "ChatServe/src/model/user"
	"ChatServe/src/utils"
	"errors"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// 注册用户
func CreateUser(c *gin.Context) {
	var userinfo userInfo.UserInfo
	c.Bind(&userinfo)
	_, err := userInfo.PhoneByFindUser(userinfo.Phone)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		userinfo.Id = uuid.NewV4()
		userInfo.CreateUser(&userinfo)
		c.JSON(utils.SuccessCode, utils.BackSuccessResp())
	} else {
		c.JSON(utils.SuccessCode, utils.BackMessageResp(utils.PhoneInUse, "phone already in use"))
	}
}

// 获取用户信息
func GetUserInfo(c *gin.Context) {
	id, _ := c.Get("userid")
	strID := (id).(uuid.UUID)
	userinfo := userInfo.FindUserInfoByUUID(strID)
	if userinfo.Id == id {
		c.JSON(200, utils.BackDataResp(200, "", userinfo))
	} else {
		c.JSON(200, utils.BackMessageResp(utils.FailedCode, "请稍后再试"))
	}
}

//根据手机号码获取用户信息

func GetUserInfoByPhone(c *gin.Context) {
	id, _ := c.Get("userid")
	strID := (id).(uuid.UUID)
	type Data struct {
		Phone string `json:"phone" form:"phone"`
	}
	var data Data
	err := c.Bind(&data)
	if err != nil {
		c.JSON(200, utils.BackMessageResp(0, "获取失败"))
		return
	}
	Users, err := userInfo.PhoneByFindUser(data.Phone)
	if err != nil {
		//
		c.JSON(200, utils.BackMessageResp(0, "获取失败"))
		return
	}
	friendTableInfo, _ := friends.GetFriendsRelationship(strID, Users.Id)

	type result struct {
		Relationship string
		Data         userInfo.UserInfo
	}
	var r result
	r.Relationship = friendTableInfo.Relationship
	r.Data = Users

	c.JSON(200, gin.H{
		"code": 0,
		"msg":  r,
	})
}

// 修改用户信息
func SetUserInfo(c *gin.Context) {
	var userinfo userInfo.UserInfo
	c.Bind(&userinfo)
}
