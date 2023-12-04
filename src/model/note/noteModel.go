package noteModel

import (
	"ChatServe/src/model/gormMode"
	userInfo "ChatServe/src/model/user"
	uuid "github.com/satori/go.uuid"
)

// 日记表
type Note struct {
	Id          uuid.UUID `gorm:"primaryKey"`
	NoteTitle   string    `form:"title" json:"title"`
	NoteContent string    `form:"content" json:"content"`
	Status      bool      `form:"status" json:"status"`
	gormMode.Model
}

// 日记黑名单表
type BlackList struct {
	Id             uuid.UUID         `gorm:"primary_key"`
	BlackListUseId string            `gorm:"column:black_user;index;not null" form:"blackListUseId" json:"blackListUseId"`
	NoteId         string            `gorm:"column:note_id;index;not null" form:"noteId" json:"noteId"`
	UserInfo       userInfo.UserInfo `gorm:"foreignkey:BlackListUseId;association_foreignkey:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Note           Note              `gorm:"foreignkey:NoteId;association_foreignkey:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" `
	gormMode.Model
}

// TableName 日记表名
func (log *Note) TableName() string {
	return "note"
}

// TableName 日记黑名单表名
func (log *BlackList) TableName() string {
	return "blacklist"
}
