package userInfo

import (
	"ChatServe/src/model/gormMode"
	uuid "github.com/satori/go.uuid"
)

type UserInfo struct {
	gormMode.Model
	Id          uuid.UUID `gorm:"primaryKey"`
	UserName    string    `json:"username" form:"username" binding:"required"`
	PassWord    string    `json:"-" form:"password" binding:"required"`
	Phone       string    `json:"phone" form:"phone" binding:"required"`
	Sex         int       `json:"sex"  form:"sex"`
	Birthday    int64     `json:"birthday"  form:"birthday"`
	Email       string    `json:"email"  form:"email"`
	State       bool      `json:"state" from:"state" gorm:"default:false;comment:'在线状态'"`
	Permissions string    `json:"permissions" default:"2"  form:"permissions"`
}
