package chatContent

import (
	"ChatServe/src/model/gormMode"
	uuid "github.com/satori/go.uuid"
)

type ChatContent struct {
	ID       uuid.UUID `json:"id" form:"id" gorm:"id"`
	UserId   uuid.UUID `json:"userId" form:"userId" gorm:"userId"`
	FriendId uuid.UUID `json:"friendId" form:"friendId" gorm:"friendId"`
	Content  string    `json:"content" form:"content" gorm:"content"`
	gormMode.Model
}

func (c *ChatContent) TableName() string {
	return "chatContent"
}
