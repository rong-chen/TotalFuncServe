package friends

import (
	"ChatServe/src/utils"
	uuid "github.com/satori/go.uuid"
)

func AddFriends(addFriends *Friends) error {
	return utils.DB.Create(addFriends).Error
}

func FindAllFriends(userId uuid.UUID) (friendList []Friends) {
	utils.DB.Where("userId = ?", userId).Find(&friendList)
	return friendList
}

// 获取好友关系

func GetFriendsRelationship(userId uuid.UUID, friendId uuid.UUID) (friend Friends, err error) {
	err = utils.DB.Where("user_id = ? and friends_id = ?", userId, friendId).First(&friend).Error
	return friend, err
}

// 更新好友关系

func UpdateRelationship(userId uuid.UUID, friendId uuid.UUID, relationship string) error {
	var model Friends
	return utils.DB.Model(&model).Where("user_id = ? and friends_id = ?", userId, friendId).Update("relationship", relationship).Error
}
