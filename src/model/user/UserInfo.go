package userInfo

import (
	"ChatServe/src/utils"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

func (table *UserInfo) TableName() string {
	return "userinfo"
}
func CreateUser(info *UserInfo) *gorm.DB {
	return utils.DB.Create(info)
}

func PhoneByFindUser(phone string) (user UserInfo, err error) {
	err = utils.DB.Where("phone=?", phone).First(&user).Error
	return user, err
}

//	func FindUuid(uuid string) UserInfo {
//		var user UserInfo
//		utils.DB.Where("uuid = ?", uuid).Find(&user)
//		return user
//	}
func FindUserInfoByUUID(uuid uuid.UUID) (userinfo UserInfo) {
	utils.DB.Where("id = ?", uuid).Find(&userinfo)
	return userinfo
}
func UpdateUserInfo(updateUserInfo *UserInfo) error {
	return utils.DB.Save(updateUserInfo).Error
}
