package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
)

var (
	DB *gorm.DB
)

func InitMySQL() (err error) {
	DB, err = gorm.Open("mysql", "root:1234@(127.0.0.1:3306)/school?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		logrus.Error("init mysql failed")
		panic(err)
	}
	return
}
func Close() {
	DB.Close()
	RDB.Close()
}
