package utils

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	logger2 "gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func InitMySQL() {
	username := viper.GetString("Develop.username")
	password := viper.GetString("Develop.password")
	localhost := viper.GetString("Develop.localhost")
	port := viper.GetString("Develop.port")
	sqlname := viper.GetString("Develop.sqlname")
	dns := username + ":" + password + "@tcp" + "(" + localhost + ":" + port + ")" + "/" + sqlname + "?charset=utf8mb4&parseTime=True&loc=Local"
	newLogger := logger2.New(
		log.New(os.Stdout, "\n\n", log.LstdFlags),
		logger2.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger2.Info,
			Colorful:      false,
		})
	DB, _ = gorm.Open(mysql.Open(dns), &gorm.Config{
		Logger: newLogger,
	})
}
