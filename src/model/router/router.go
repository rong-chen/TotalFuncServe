package routers

import (
	"ChatServe/src/utils"
	"gorm.io/gorm"
)

type Router struct {
	gorm.Model
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	Path        string `json:"path"`
	Level       int    `json:"level"`
	Parent      string `json:"parent"`
	Status      int    `json:"status"`
	Permissions int    `json:"permissions"`
}

func (r *Router) TableName() string {
	return "router"
}
func GetRouters() (r []Router) {
	utils.DB.Model(&Router{}).Find(&r)
	return
}
