package main

import (
	app "ChatServe/src"
	"ChatServe/src/model/friends"
	noteModel "ChatServe/src/model/note"
	routers "ChatServe/src/model/router"
	userInfo "ChatServe/src/model/user"
	utils "ChatServe/src/utils"
	"github.com/spf13/viper"
)

func main() {
	Viper()
	utils.InitMySQL()
	//initTable()
	r := app.APP()
	r.Run(":6661")
}

// 初始化表
func initTable() {
	utils.DB.AutoMigrate(&userInfo.UserInfo{})
	utils.DB.AutoMigrate(&noteModel.Note{})
	utils.DB.AutoMigrate(&noteModel.BlackList{})
	utils.DB.AutoMigrate(&routers.Router{})
	utils.DB.AutoMigrate(&friends.Friends{})
}
func Viper() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
