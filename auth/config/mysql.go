package config

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Db xxx
var Db *gorm.DB

// InitDb xxx
func InitDb() {
	var err error
	Db, err = gorm.Open("mysql", "root:@/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal("数据库启动失败", err.Error())
	}
}
