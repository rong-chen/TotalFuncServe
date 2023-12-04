package friendsServe

import (
	"ChatServe/src/model/friends"
	"ChatServe/src/utils"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func AddFriends(c *gin.Context) {
	id, _ := c.Get("userid")
	uuidID := (id).(uuid.UUID)
	type Data struct {
		UserId       string `json:"userId" form:"userId"`
		FriendId     string `json:"friendId" form:"friendId"`
		Relationship string `json:"relationship" form:"relationship"`
	}
	var data Data
	err := c.Bind(&data)
	if err != nil {
		c.JSON(200, utils.BackMessageResp(201, "添加失败，参数错误"))
		return
	}
	var params friends.Friends
	FriendsId, err := uuid.FromString(data.FriendId)
	if err != nil {
		return
	}
	params.UserId = uuidID
	params.FriendsId = FriendsId
	params.Relationship = data.Relationship
	params.ID = uuid.NewV4()
	if err != nil {
		c.JSON(200, utils.BackMessageResp(201, "添加失败，Id错误"))
		return
	}
	f, err := friends.GetFriendsRelationship(params.UserId, params.FriendsId)
	if f.Relationship == "3" || f.Relationship == "2" || f.Relationship == "1" {
		c.JSON(200, utils.BackMessageResp(201, "已有该记录，请勿重复添加"))
		return
	}
	errs := friends.AddFriends(&params)
	if errs != nil {
		c.JSON(200, utils.BackMessageResp(201, "申请好友失败"))
		return
	}
	c.JSON(200, utils.BackMessageResp(200, "添加成功"))
}

func UpdateFrinds(c *gin.Context) {
	var params friends.Friends
	err := c.Bind(&params)
	if err != nil {
		return
	}
	err = friends.UpdateRelationship(params.UserId, params.FriendsId, params.Relationship)
	if err != nil {
		c.JSON(200, utils.BackMessageResp(201, "参数错误"))
		return
	}
	c.JSON(200, utils.BackMessageResp(200, "更新成功"))
}
