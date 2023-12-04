package friends

import uuid "github.com/satori/go.uuid"

// 1 好友 2 黑名单 3 请求添加好友

type Friends struct {
	ID           uuid.UUID `form:"Id" gorm:"primaryKey"`
	UserId       uuid.UUID `form:"userId" json:"userId"`
	FriendsId    uuid.UUID `form:"friendsId" json:"friendsId"`
	Relationship string    `form:"relationship" json:"relationship"`
}

func (friends *Friends) TableName() string {
	return "friendTable"
}
